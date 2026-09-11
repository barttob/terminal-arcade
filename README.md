# Terminal Arcade

A collection of classic arcade games that run in your terminal, written in Go with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss).

Start the app, pick a game from the menu, play, and return to the menu without restarting. Each game is a self-contained module behind a shared game interface, so new games can be added easily. The long-term goal is to also serve the arcade over SSH, so friends can connect and play together from their own terminals.

> [!WARNING]
> **Work in progress.** This project is in early development. Expect missing features, rough edges, and breaking changes between commits.

## Status

| Game        | State                                                                    |
| ----------- | ------------------------------------------------------------------------ |
| Minesweeper | ✅ Playable: Beginner/Intermediate/Expert presets and custom boards       |
| Snake       | 🚧 Listed in the menu but not playable yet                               |

## Getting started

### Requirements

- [Go](https://go.dev/dl/) 1.25 or newer
- A terminal with Unicode and 256-color support (e.g. Windows Terminal, iTerm2, most Linux terminals), at least 80×24

### Run from source

```bash
git clone https://github.com/barttob/terminal-arcade.git
cd terminal-arcade
go run ./cmd/arcade
```

### Build a binary

```bash
go build -o arcade ./cmd/arcade    # on Windows: go build -o arcade.exe ./cmd/arcade
./arcade
```

## Controls

| Where       | Keys                                                                                      |
| ----------- | ----------------------------------------------------------------------------------------- |
| Everywhere  | `q` / `ctrl+c` quit                                                                       |
| Main menu   | `↑`/`↓` (or `k`/`j`) move · `enter` select                                                |
| Settings    | `tab` / `shift+tab` switch game · `↑`/`↓` move · `←`/`→` change · `enter` play · `esc` back |
| Minesweeper | arrows/WASD move · `space` reveal · `f` flag · `o` settings · `r` restart · `esc` menu    |

## Roadmap

A rough plan. The full design and phase breakdown is in [terminal_arcade_collection_go_bubbletea.md](terminal_arcade_collection_go_bubbletea.md).

- [x] **Foundation**: app shell, main menu, shared styles, shared engine (grid, input, collision, ticks)
- [x] **Minesweeper**: first playable game
- [x] **Settings**: per-game settings tabs
- [ ] **Snake**: finish movement, collision and rendering
- [ ] **Core screens**: pause, game-over and help screens
- [ ] **More games**: Pong, Breakout, Tetris, Space Invaders, 2048, and others
- [ ] **Persistence**: save settings and high scores between runs
- [ ] **SSH mode**: host the arcade with [Wish](https://github.com/charmbracelet/wish) so others can play via `ssh`
- [ ] **Multiplayer**: lobby, rooms, and two-player games over SSH, starting with Pong
