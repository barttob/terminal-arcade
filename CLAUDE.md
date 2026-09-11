# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
go run ./cmd/arcade      # run the arcade locally (alt-screen TUI)
go build ./...           # compile everything
go vet ./...             # vet
gofmt -l .               # list unformatted files (see CRLF note below)
go test ./...
go test ./internal/games/minesweeper -run TestFlagBlocksReveal   # single test
```

The repo is checked out with `core.autocrlf=true`, so `gofmt -l` reports nearly every file purely because of CRLF endings. To find real formatting problems, strip `\r` first (e.g. `tr -d '\r' < file | gofmt -d`).

There is no Makefile (the design doc proposes one). `go.mod` currently marks every dependency `// indirect`; `go mod tidy` will correct that once the direct imports settle.

## Architecture

Bubble Tea (Elm architecture) app in three layers: launcher (`internal/app`, `internal/menu`), shared engine (`internal/engine`), and individual games (`internal/games/<game>`). `cmd/arcade/main.go` only wraps `app.New()` in a `tea.Program` with `WithAltScreen`.

**Root model** — `internal/app` holds an `AppState` enum ([state.go](internal/app/state.go)) and switches on it in both `Update` and `View`. Every screen is rendered through the `center()` helper in [model.go](internal/app/model.go#L34), which is why child views should not do their own full-terminal centering. `appModel` is a value type; so are `menu.Model` and `settings.Model`, whose `Update` returns `tea.Model` and must be type-asserted back on assignment.

**Game interface** — `engine.Game` ([game.go](internal/engine/game.go)) is deliberately *not* `tea.Model`: its `Update` returns `(engine.Game, tea.Cmd)`. Games implement it on `*Model` (pointer receivers), so they mutate in place and return themselves. `Reset()` returns a fresh `engine.Game` rather than mutating.

**Game registration** — `games.Registry` in [internal/games/registry.go](internal/games/registry.go) is the single source of truth for which games exist: each `Entry{ID, Name, Factory, Settings}` automatically gets a main-menu item and its own tab (with a Play button) on the settings screen. Adding a game means: create `internal/games/<name>/` following the four-file convention below, export an `ID` const and a `New(...)` constructor, and append an entry. `Factory` is `func(settings.Settings) engine.Game` (games without settings ignore the argument), and `Settings` is the game's `[]settings.Row` (nil shows "No settings yet"). The registry lives in its own package because `settings` can't import it: `games` → `settings` → `minesweeper` would cycle. The menu appends "Settings" and "Exit" after the registry entries, and its `enter` handler indexes off `len(m.choices)` to tell them apart.

**Per-game file convention** (mirrored by `snake` and `minesweeper`):
- `model.go` — the `Model` struct and `NewModel(width, height)` constructor
- `update.go` — input handling and per-tick `step()` logic
- `view.go` — rendering only
- `<game>.go` — the `engine.Game` interface methods, package consts (`ID`, `DefaultWidth`, `DefaultHeight`, `DefaultTickRate`), and `New() engine.Game`

Turn-based games (minesweeper) skip the tick: `Init()` returns `nil`, there's no `step()`, and `r` restart is handled by returning `m.Reset()` from `Update`.

**Shared engine** — `Point`/`Grid`, `Direction` with `DirectionFromKey` (arrows + WASD) and `Opposite()`, `IsCollision`/`ContainsPoint`, `Score`, and `Tick(id, d)` which wraps `tea.Tick` into a `TickMsg{ID, Time}`. A tick-based game takes an ID from `engine.NewTickID()` at construction and ignores `TickMsg`s with any other ID; otherwise a tick still in flight from a restarted or abandoned game would start a second tick chain and double the speed. Snake shows the pattern: `Init` arms the first tick, each own tick re-arms, and pausing or game over lets the chain lapse. Unpausing takes a fresh ID before re-arming, so a tick still in flight from before the pause can't start a second chain.

**Styling** — all lipgloss styles live in [internal/styles/styles.go](internal/styles/styles.go). Do not define colors inside game packages.

**Settings** — [settings/settings.go](internal/settings/settings.go) holds the `Settings` value (one field per game, e.g. `Minesweeper minesweeper.Config`) and each game's `[]Row` (`Label`, `Value`, `Adjust(delta)`); adding a setting means adding a row there. The screen ([settings/model.go](internal/settings/model.go)) shows one tab per registry entry, built by `games.SettingsSections()`; its cursor runs over the tab's rows plus a trailing Play button, and `enter` anywhere starts that tab's game. Validation belongs to the game (`minesweeper.Config.Clamped()`), not the screen. The app owns a `settings.Model` for the whole session, rather than a global, so each future SSH session gets its own. Settings are in-memory only; the design doc proposes `~/.terminal-arcade/config.json` for persistence.

**Message flow** — navigation goes through two messages in [engine/game.go](internal/engine/game.go), which the menu, the settings screen and games all emit:
- `engine.StartGameMsg{GameID}` → `app.startGame` looks the ID up in the registry, builds it with `Factory(m.settings.Values())`, and returns `Init()`.
- `engine.OpenSettingsMsg{GameID}` → switches to `StateSettings` with `settings.Focus(GameID)` (an empty ID keeps the last tab). Both games send this on `o`.

While playing, `app.Update` forwards *every* message (keys, `TickMsg`, `WindowSizeMsg`) to `currentGame.Update`, so games must type-switch rather than assert. `q`/`ctrl+c` quit globally before dispatch. `esc` is handled in `app.Update`: from a game it drops the game and returns to the menu. From settings it resumes `currentGame` if settings were opened from a game, otherwise it returns to the menu. Messages aren't delivered to a game while settings are open, and the app sends nothing on resume, so a tick-based game should pause itself before emitting `OpenSettingsMsg` (as snake does) and re-arm when the player unpauses.

## Current state

- Minesweeper is playable, with Beginner/Intermediate/Expert/custom boards configurable from Settings. Mines are placed on the first reveal, so the opening is always safe.
- Snake is playable, with speed (Slow/Normal/Fast/Insane), board size up to 38×19 (the most that fits 80×24), and walls on/off (off wraps around) configurable from Settings. Key presses are queued (up to 3) and applied one per tick, so fast turns aren't dropped and can't reverse the snake into itself.

## Design doc

[terminal_arcade_collection_go_bubbletea.md](terminal_arcade_collection_go_bubbletea.md) is the full project spec — folder layout, per-file responsibilities, controls table, rendering characters, and planned games. Consult it for intended design before inventing new structure. Planned but unimplemented: SSH mode via Wish (`cmd/arcade-server`, `internal/server`), multiplayer lobby/rooms (`internal/lobby`, the `StateLobby`/`StateRoom` states, and `appModel.remote`), and games beyond Snake. The wish/ssh/keygen dependencies are already in `go.mod` but unused.

Consistent controls across games: WASD/arrows to move, `space` action, `p` pause, `r` restart, `esc` back to menu, `q` quit.
