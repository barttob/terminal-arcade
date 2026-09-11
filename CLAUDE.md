# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
go run ./cmd/arcade      # run the arcade locally (alt-screen TUI)
go build ./...           # compile everything
go vet ./...             # vet
gofmt -l .               # list unformatted files
go test ./...            # no tests exist yet
go test ./internal/engine -run TestName   # single test, once tests exist
```

There is no Makefile (the design doc proposes one). `go.mod` currently marks every dependency `// indirect`; `go mod tidy` will correct that once the direct imports settle.

## Architecture

Bubble Tea (Elm architecture) app in three layers: launcher (`internal/app`, `internal/menu`), shared engine (`internal/engine`), and individual games (`internal/games/<game>`). `cmd/arcade/main.go` only wraps `app.New()` in a `tea.Program` with `WithAltScreen`.

**Root model** — `internal/app` holds an `AppState` enum ([state.go](internal/app/state.go)) and switches on it in both `Update` and `View`. Every screen is rendered through the `center()` helper in [model.go](internal/app/model.go#L34), which is why child views should not do their own full-terminal centering. `appModel` is a value type; `menu.Model` is too, so `menu.Update` returns `tea.Model` and must be type-asserted back on assignment.

**Game interface** — `engine.Game` ([game.go](internal/engine/game.go)) is deliberately *not* `tea.Model`: its `Update` returns `(engine.Game, tea.Cmd)`. Games implement it on `*Model` (pointer receivers), so they mutate in place and return themselves. `Reset()` returns a fresh `engine.Game` rather than mutating.

**Game registration** — `gameRegistry` in [internal/menu/model.go](internal/menu/model.go#L23) is the single source of truth for what appears in the menu. Adding a game means: create `internal/games/<name>/` following the four-file convention below, export `New() engine.Game`, and append a `gameEntry{Name, Factory}`. The menu appends "Settings" and "Exit" after the registry entries, and its `enter` handler indexes off `len(m.choices)` to tell them apart.

**Per-game file convention** (mirrored by `snake` and `minesweeper`):
- `model.go` — the `Model` struct and `NewModel(width, height)` constructor
- `update.go` — input handling and per-tick `step()` logic
- `view.go` — rendering only
- `<game>.go` — the `engine.Game` interface methods, package consts (`ID`, `DefaultWidth`, `DefaultHeight`, `DefaultTickRate`), and `New() engine.Game`

**Shared engine** — `Point`/`Grid`, `Direction` with `DirectionFromKey` (arrows + WASD) and `Opposite()`, `IsCollision`/`ContainsPoint`, `Score`, and `Tick(d)` which wraps `tea.Tick` into a `TickMsg`. Note `Tick` multiplies its argument by `time.Millisecond`, so callers passing an already-scaled `time.Duration` (as both games currently do with `DefaultTickRate`) get a far longer interval than intended.

**Styling** — all lipgloss styles live in [internal/styles/styles.go](internal/styles/styles.go). Do not define colors inside game packages.

**Message flow** — menu emits `menu.GameSelectedMsg{Game}`; `app.Update` catches it, stores `currentGame`, sets `StatePlaying`, and returns `currentGame.Init()`.

## Current state

Early scaffolding. Wiring gaps to be aware of before touching game code:
- `app.Update` never dispatches to `m.currentGame.Update`, so a started game receives neither keys nor ticks.
- `snake.Update` unconditionally type-asserts `msg.(tea.KeyMsg)` and will panic on any other message once dispatch is hooked up.
- `snake.step()` and `snake.View()` are stubs; `minesweeper` is a skeleton with no board.
- `menu` cursor bound is `len(m.choices)+1` (off by one past the last selectable item).

## Design doc

[terminal_arcade_collection_go_bubbletea.md](terminal_arcade_collection_go_bubbletea.md) is the full project spec — folder layout, per-file responsibilities, controls table, rendering characters, and planned games. Consult it for intended design before inventing new structure. Planned but unimplemented: SSH mode via Wish (`cmd/arcade-server`, `internal/server`), multiplayer lobby/rooms (`internal/lobby`, the `StateLobby`/`StateRoom` states, and `appModel.remote`), and games beyond Snake. The wish/ssh/keygen dependencies are already in `go.mod` but unused.

Consistent controls across games: WASD/arrows to move, `space` action, `p` pause, `r` restart, `esc` back to menu, `q` quit.
