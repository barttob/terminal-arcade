# Terminal Arcade Collection in Go with Bubble Tea

## 1. Project Overview

This project is a collection of classic arcade-style terminal games built in **Go** using the **Bubble Tea** TUI framework.

The goal is to create one terminal application that works like an arcade launcher. The user starts the app, chooses a game from a menu, plays it, and can return to the main menu without restarting the program.

Example games:

- Snake
- Pong
- Breakout
- Tetris
- Space Invaders
- Flappy Bird
- Tron Light Cycles
- ASCII Racing
- Whack-a-Mole
- Minesweeper
- 2048

The application should be modular so that new games can be added easily.

---

## 2. Main Goals

### Functional goals

- Run multiple games from one terminal app.
- Provide a clean main menu for selecting games.
- Allow pausing, restarting, quitting, and returning to the menu.
- Keep a consistent control system across games where possible.
- Support score tracking.
- Support simple game-over screens.
- Make each game independent from the others.

### Technical goals

- Use Go as the main language.
- Use Bubble Tea for terminal UI architecture.
- Use Lip Gloss for styling.
- Use a shared game interface for all games.
- Keep rendering, input handling, and game state separated.
- Make the project easy to expand.

---

## 3. Recommended Tech Stack

| Tool                    | Purpose                                |
| ----------------------- | -------------------------------------- |
| Go                      | Main programming language              |
| Bubble Tea              | Terminal UI framework                  |
| Lip Gloss               | Styling terminal views                 |
| Bubbles                 | Optional reusable components           |
| Wish                    | SSH server framework for terminal apps |
| Bubble Tea SSH renderer | Run Bubble Tea apps over SSH           |
| Go modules              | Dependency management                  |

Recommended dependencies:

```bash
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
go get github.com/charmbracelet/bubbles
go get github.com/charmbracelet/wish
go get github.com/charmbracelet/wish/bubbletea
go get github.com/charmbracelet/wish/logging
```

Wish allows the arcade to run as an SSH-accessible terminal application. This means players can connect remotely using a normal SSH client and play supported games without installing the app locally.

Example usage after the server is running:

```bash
ssh arcade.example.com
```

Or locally during development:

```bash
ssh localhost -p 23234
```

\---|---| | Go | Main programming language | | Bubble Tea | Terminal UI framework | | Lip Gloss | Styling terminal views | | Bubbles | Optional reusable components | | Go modules | Dependency management |

Recommended dependencies:

```bash
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
go get github.com/charmbracelet/bubbles
```

---

## 4. High-Level Application Structure

The app should have three main layers:

1. **Launcher layer**\
   Handles the main menu, game selection, settings, and navigation.

2. **Game engine layer**\
   Contains shared types, interfaces, messages, timing utilities, score handling, and rendering helpers.

3. **Individual game layer**\
   Each game has its own model, update logic, view rendering, and rules.

Conceptually:

```text
App
├── Local Terminal Mode
├── SSH Server Mode
├── Main Menu
├── Game Router
├── Multiplayer Session Manager
├── Shared Engine
└── Games
    ├── Snake
    ├── Pong
    ├── Tetris
    ├── Tron Light Cycles
    └── Co-op Games
```

The application should support two ways of running:

1. **Local mode** — runs directly in the user's terminal.
2. **SSH mode** — starts a Wish SSH server so remote players can connect and play through SSH.

---

## 5. Suggested Folder Structure

```text
terminal-arcade/
├── cmd/
│   ├── arcade/
│   │   └── main.go
│   └── arcade-server/
│       └── main.go
│
├── internal/
│   ├── app/
│   │   ├── model.go
│   │   ├── update.go
│   │   ├── view.go
│   │   └── state.go
│   │
│   ├── server/
│   │   ├── ssh.go
│   │   ├── session.go
│   │   ├── lobby.go
│   │   └── auth.go
│   │
│   ├── engine/
│   │   ├── game.go
│   │   ├── multiplayer.go
│   │   ├── tick.go
│   │   ├── input.go
│   │   ├── grid.go
│   │   ├── score.go
│   │   ├── collision.go
│   │   └── render.go
│   │
│   ├── menu/
│   │   ├── model.go
│   │   ├── update.go
│   │   └── view.go
│   │
│   ├── lobby/
│   │   ├── model.go
│   │   ├── update.go
│   │   └── view.go
│   │
│   ├── styles/
│   │   └── styles.go
│   │
│   └── games/
│       ├── snake/
│       │   ├── model.go
│       │   ├── update.go
│       │   ├── view.go
│       │   └── snake.go
│       │
│       ├── pong/
│       │   ├── model.go
│       │   ├── update.go
│       │   ├── view.go
│       │   └── pong.go
│       │
│       ├── breakout/
│       │   ├── model.go
│       │   ├── update.go
│       │   ├── view.go
│       │   └── breakout.go
│       │
│       └── tetris/
│           ├── model.go
│           ├── update.go
│           ├── view.go
│           └── tetris.go
│
├── assets/
│   └── ascii/
│       └── logo.txt
│
├── docs/
│   ├── controls.md
│   ├── adding-new-games.md
│   └── roadmap.md
│
├── go.mod
├── go.sum
├── README.md
└── Makefile
```

---

## 6. File and Directory Responsibilities

This section describes what each file and directory in the suggested structure should contain.

### Root directory

```text
terminal-arcade/
```

The root directory contains the whole project. It should include the Go module files, documentation, build helpers, assets, and all application source code.

### `cmd/`

```text
cmd/
├── arcade/
│   └── main.go
└── arcade-server/
    └── main.go
```

The `cmd` directory contains executable entry points. Each subdirectory should build into a separate binary.

#### `cmd/arcade/main.go`

This is the entry point for the local terminal version of the arcade.

It should:

- Create the root Bubble Tea app model with `app.New()`.
- Start Bubble Tea with `tea.NewProgram`.
- Enable full-screen terminal mode with `tea.WithAltScreen()`.
- Handle startup errors.
- Exit cleanly if the app fails.

This file should stay small. It should not contain game logic, menu logic, rendering logic, or configuration parsing beyond what is needed to start the app.

#### `cmd/arcade-server/main.go`

This is the entry point for the SSH multiplayer server.

It should:

- Load server configuration.
- Start the Wish SSH server.
- Register Bubble Tea middleware for SSH sessions.
- Create one remote app model per connected SSH session using `app.NewRemote(...)`.
- Configure host, port, host key, and logging middleware.
- Handle graceful shutdown when the process receives an interrupt signal.

This file should not directly manage rooms or game state. That responsibility belongs to the `internal/server` and `internal/lobby` packages.

---

### `internal/`

```text
internal/
```

The `internal` directory contains application code that should not be imported by external projects. This is where most of the arcade logic lives.

---

### `internal/app/`

```text
internal/app/
├── model.go
├── update.go
├── view.go
└── state.go
```

The `app` package contains the root Bubble Tea model. It controls the current screen, routes messages, and connects the menu, games, lobby, and global controls.

#### `internal/app/model.go`

Contains the main root model struct and constructors.

It should define things like:

```go
type Model struct {
    state       AppState
    menu        menu.Model
    lobby       lobby.Model
    currentGame engine.Game
    playerID    engine.PlayerID
    width       int
    height      int
    remote      bool
}
```

It should contain constructors such as:

```go
func New() Model
func NewRemote(options RemoteOptions) Model
```

Use `New()` for local terminal play and `NewRemote()` for SSH sessions.

#### `internal/app/update.go`

Contains the root Bubble Tea `Update` function.

It should:

- Handle global keyboard shortcuts.
- Track terminal resize messages.
- Route messages to the menu when in the main menu.
- Route messages to the active game when playing.
- Route messages to the lobby when in SSH lobby mode.
- Switch between states such as menu, playing, paused, game over, settings, and help.

This file is the main message router of the application.

#### `internal/app/view.go`

Contains the root Bubble Tea `View` function.

It should:

- Render the correct screen depending on the current app state.
- Show the main menu, game view, pause screen, game-over screen, lobby screen, settings screen, or help screen.
- Handle small terminal warnings.
- Apply shared layout styles.

The root view should not manually draw every game. It should call `currentGame.View()` or `currentGame.ViewForPlayer(...)`.

#### `internal/app/state.go`

Contains the app state enum and state-related helpers.

It should define:

```go
type AppState int

const (
    StateMainMenu AppState = iota
    StatePlaying
    StatePaused
    StateGameOver
    StateSettings
    StateHelp
    StateLobby
    StateRoom
)
```

It can also contain helper methods such as:

```go
func (s AppState) String() string
```

---

### `internal/server/`

```text
internal/server/
├── ssh.go
├── session.go
├── lobby.go
└── auth.go
```

The `server` package contains infrastructure for running the arcade through SSH using Wish.

#### `internal/server/ssh.go`

Contains the Wish server setup.

It should:

- Create and configure the SSH server.
- Set the server address and port.
- Load or generate the SSH host key.
- Attach Bubble Tea middleware.
- Attach logging middleware.
- Create a remote app model for each SSH session.
- Provide a function such as `Start(config Config) error`.

This file connects Wish to your Bubble Tea application.

#### `internal/server/session.go`

Tracks connected SSH users.

It should define:

```go
type SessionID string

type RemoteSession struct {
    ID       SessionID
    PlayerID engine.PlayerID
    Username string
    RoomID   lobby.RoomID
}
```

It should also contain a session manager:

```go
type SessionManager struct {
    sessions map[SessionID]RemoteSession
}
```

Responsibilities:

- Add new sessions.
- Remove disconnected sessions.
- Find sessions by ID.
- List active sessions.
- Associate a session with a player ID and room ID.

#### `internal/server/lobby.go`

Connects SSH sessions to the multiplayer lobby and room manager.

It should:

- Hold a shared lobby or room manager instance used by all SSH sessions.
- Provide methods for joining rooms, leaving rooms, and listing rooms.
- Route remote player input into the correct room.
- Clean up player state when an SSH session disconnects.

This file should not contain visual lobby rendering. Visual lobby rendering belongs to `internal/lobby/view.go`.

#### `internal/server/auth.go`

Contains optional SSH authentication logic.

It can support:

- No authentication for local development.
- Public key authentication.
- Password authentication.
- Loading allowed public keys from an `authorized_keys` file.
- Rejecting unknown users.

For the first version, this file can be minimal or even return permissive development authentication.

---

### `internal/engine/`

```text
internal/engine/
├── game.go
├── multiplayer.go
├── tick.go
├── input.go
├── grid.go
├── score.go
├── collision.go
└── render.go
```

The `engine` package contains shared types and helpers used by multiple games. It should not know about specific games like Snake or Pong.

#### `internal/engine/game.go`

Defines the base single-player game interface.

It should contain:

```go
type Game interface {
    Name() string
    Description() string
    Init() tea.Cmd
    Update(msg tea.Msg) (Game, tea.Cmd)
    View() string
    Reset() Game
    IsGameOver() bool
    Score() int
}
```

It can also contain shared metadata types used by game registration.

#### `internal/engine/multiplayer.go`

Defines multiplayer-specific types and interfaces.

It should contain:

- `PlayerID`
- `Player`
- `PlayerInput`
- `PlayerInputMsg`
- `MultiplayerGame`

This file is the contract that remote co-op and PvP games must implement.

#### `internal/engine/tick.go`

Contains reusable timing messages and tick helpers.

It should define:

```go
type TickMsg time.Time

func Tick(d time.Duration) tea.Cmd
```

Games use this to update themselves repeatedly without writing their own timing code every time.

#### `internal/engine/input.go`

Contains shared input types and helpers.

It should define:

- Direction enum: `Up`, `Down`, `Left`, `Right`.
- Functions for converting key strings into directions.
- Helpers for WASD and arrow-key controls.

Example helpers:

```go
func DirectionFromKey(key string) (Direction, bool)
func IsActionKey(key string) bool
func IsQuitKey(key string) bool
```

#### `internal/engine/grid.go`

Contains grid and point utilities.

It should define:

- `Point`
- `Grid`
- `NewGrid`
- `InBounds`
- `Set`
- `Get`
- `Clear`
- `Render`

This is useful for Snake, Tetris, Minesweeper, Pac-Man-style games, Bomberman-style games, and roguelike arena games.

#### `internal/engine/score.go`

Contains score-related types and helpers.

It should define:

```go
type Score struct {
    Current int
    High    int
}
```

Later, this file can also include score formatting, score comparison, and high-score update helpers.

Persistent score storage may eventually be moved into a separate package, such as `internal/storage`.

#### `internal/engine/collision.go`

Contains reusable collision helpers.

It should include:

- Point equality.
- Checking whether a list contains a point.
- Rectangle overlap detection.
- Boundary collision helpers.

Useful for Snake, Pong, Breakout, Tron, and many other games.

#### `internal/engine/render.go`

Contains generic rendering helpers.

It should include functions such as:

```go
func RenderBox(title string, body string) string
func RenderCentered(width int, text string) string
func RenderGameHeader(name string, score int) string
```

This file should only contain general rendering helpers. Game-specific rendering belongs inside each game's `view.go`.

---

### `internal/menu/`

```text
internal/menu/
├── model.go
├── update.go
└── view.go
```

The `menu` package contains the local main menu used to choose games.

#### `internal/menu/model.go`

Contains the menu model struct and constructor.

It should define:

```go
type Model struct {
    choices []games.GameInfo
    cursor  int
}
```

It can also define selection result messages, for example:

```go
type GameSelectedMsg struct {
    GameID string
}
```

#### `internal/menu/update.go`

Handles menu input.

It should:

- Move the cursor up and down.
- Select a game when Enter is pressed.
- Return a message to the root app when a game is selected.
- Optionally handle help/settings shortcuts.

The menu should not start games directly. It should notify the root app which game was selected.

#### `internal/menu/view.go`

Renders the main menu.

It should:

- Show the arcade title.
- Show available games.
- Highlight the selected game.
- Show short controls at the bottom.
- Use shared styles from `internal/styles`.

---

### `internal/lobby/`

```text
internal/lobby/
├── model.go
├── update.go
└── view.go
```

The `lobby` package contains the UI and state for SSH multiplayer room selection.

#### `internal/lobby/model.go`

Contains the lobby Bubble Tea model.

It should store:

- Available rooms.
- Cursor position.
- Current player ID.
- Selected game mode.
- Whether the player is ready.

It may also define room-related messages such as:

```go
type JoinRoomMsg struct {
    RoomID RoomID
}

type CreateRoomMsg struct {
    GameID string
}
```

#### `internal/lobby/update.go`

Handles lobby input.

It should:

- Move through available rooms.
- Create a room.
- Join a room.
- Leave a room.
- Toggle ready state.
- Refresh room list.
- Notify the root app when the player enters a room or game.

#### `internal/lobby/view.go`

Renders the SSH lobby.

It should:

- Show room list.
- Show player counts.
- Show room states: waiting, playing, finished.
- Show controls for creating and joining rooms.
- Highlight the selected room.

---

### `internal/styles/`

```text
internal/styles/
└── styles.go
```

The `styles` package contains reusable Lip Gloss styles.

#### `internal/styles/styles.go`

Contains shared visual styles used by menus, lobbies, game screens, headers, warnings, and footers.

It should define styles such as:

- `Title`
- `Subtitle`
- `Border`
- `Selected`
- `Muted`
- `Error`
- `Success`
- `Header`
- `Footer`

Keeping styles in one place makes it easier to change the visual identity of the app later.

---

### `internal/games/`

```text
internal/games/
├── snake/
├── pong/
├── breakout/
└── tetris/
```

The `games` directory contains one package per game. Each game should be independent and should expose a small public API, usually just `New()` and optional metadata.

A typical game package should use this structure:

```text
internal/games/example/
├── model.go
├── update.go
├── view.go
└── example.go
```

#### `model.go` in each game package

Contains the game's state struct and internal data types.

For example, Snake's `model.go` should contain:

- Snake body positions.
- Current direction.
- Next direction.
- Food position.
- Score.
- Width and height.
- Game-over flag.

For Tetris, `model.go` would contain the board, active piece, next piece, score, lines, and level.

#### `update.go` in each game package

Contains the game's input handling and tick update logic.

It should:

- Handle game-specific key presses.
- React to `engine.TickMsg`.
- Move the game simulation forward.
- Detect collisions.
- Update score.
- Detect game-over state.

Game logic should be split into small helper methods where possible, such as `step()`, `moveBall()`, `spawnFood()`, or `clearLines()`.

#### `view.go` in each game package

Contains the game's rendering code.

It should:

- Render the playfield.
- Render the score.
- Render the game-over message if needed.
- Render short controls.
- Use shared styles where useful.

This file should avoid changing game state. It should only read state and return a string.

#### `<game>.go` in each game package

Contains the public constructor and game metadata.

For Snake, this is:

```text
snake.go
```

It should contain:

```go
func New() engine.Game
```

It may also contain constants such as:

```go
const ID = "snake"
const Name = "Snake"
const Description = "Eat food, grow longer, and avoid crashing."
```

For multiplayer games, it can also contain:

```go
func NewMultiplayer() engine.MultiplayerGame
```

---

### `internal/games/snake/`

Snake is the recommended first game.

```text
internal/games/snake/
├── model.go
├── update.go
├── view.go
└── snake.go
```

- `model.go` stores snake position, direction, food, score, dimensions, and game-over state.
- `update.go` handles direction changes, ticking, movement, food collection, growth, and collision detection.
- `view.go` renders the snake board, score, food, and game-over screen.
- `snake.go` exposes `New()`, metadata constants, default board size, and tick rate.

---

### `internal/games/pong/`

Pong is the recommended first multiplayer game.

```text
internal/games/pong/
├── model.go
├── update.go
├── view.go
└── pong.go
```

- `model.go` stores paddle positions, ball position, ball velocity, scores, player IDs, dimensions, and game-over state.
- `update.go` handles paddle movement, ball movement, wall collisions, paddle collisions, scoring, CPU movement if single-player, and player-specific remote input.
- `view.go` renders the Pong arena, paddles, ball, scores, and controls.
- `pong.go` exposes `New()` for local play and `NewMultiplayer()` for SSH or local two-player play.

---

### `internal/games/breakout/`

Breakout is a good follow-up after Pong because it reuses paddle and ball logic.

```text
internal/games/breakout/
├── model.go
├── update.go
├── view.go
└── breakout.go
```

- `model.go` stores paddle position, ball position, velocity, bricks, lives, score, dimensions, and game-over state.
- `update.go` handles paddle movement, ball movement, wall collision, brick collision, scoring, losing lives, and level completion.
- `view.go` renders the arena, bricks, paddle, ball, lives, and score.
- `breakout.go` exposes `New()`, metadata, default dimensions, and tick rate.

---

### `internal/games/tetris/`

Tetris should be added later because it has more complex rules.

```text
internal/games/tetris/
├── model.go
├── update.go
├── view.go
└── tetris.go
```

- `model.go` stores the board, active piece, next piece, held piece if used, score, lines, level, and game-over state.
- `update.go` handles falling, soft drop, hard drop, horizontal movement, rotation, collision with fixed blocks, locking pieces, line clearing, and level speed.
- `view.go` renders the board, active piece, ghost piece if used, next piece, score, lines, and controls.
- `tetris.go` exposes `New()`, metadata, board dimensions, piece definitions, and default tick rate.

---

### `assets/`

```text
assets/
└── ascii/
    └── logo.txt
```

The `assets` directory contains non-code resources.

#### `assets/ascii/logo.txt`

Contains ASCII art for the arcade logo.

It can be loaded and displayed on the main menu or SSH lobby screen. If you want the binary to be self-contained, later you can embed it using Go's `embed` package.

---

### `docs/`

```text
docs/
├── controls.md
├── adding-new-games.md
└── roadmap.md
```

The `docs` directory contains extra project documentation.

#### `docs/controls.md`

Documents global controls and game-specific controls.

It should include:

- Global app controls.
- Menu controls.
- Lobby controls.
- Snake controls.
- Pong controls.
- Tetris controls.
- Multiplayer controls.

#### `docs/adding-new-games.md`

Explains how to add a new game package.

It should cover:

- Creating a new directory under `internal/games`.
- Implementing `engine.Game`.
- Implementing `engine.MultiplayerGame` if needed.
- Adding metadata.
- Registering the game in the registry.
- Adding tests.

#### `docs/roadmap.md`

Contains the longer project roadmap.

It can include:

- Planned games.
- Planned multiplayer features.
- Persistence features.
- Deployment improvements.
- UI polish ideas.

---

### `go.mod`

Defines the Go module path and direct dependencies.

It should include dependencies such as Bubble Tea, Lip Gloss, Bubbles, and Wish.

---

### `go.sum`

Contains cryptographic checksums for Go module dependencies.

This file is generated automatically by Go and should be committed to version control.

---

### `README.md`

The main project documentation shown on GitHub or another repository host.

It should include:

- Project description.
- Screenshots or terminal previews.
- Installation instructions.
- Local running instructions.
- SSH server running instructions.
- Controls.
- List of available games.
- Development roadmap.

---

### `Makefile`

Contains useful development commands.

Example targets:

```makefile
run:
	go run ./cmd/arcade

server:
	go run ./cmd/arcade-server

build:
	go build -o bin/arcade ./cmd/arcade
	go build -o bin/arcade-server ./cmd/arcade-server

test:
	go test ./...

fmt:
	go fmt ./...
```

This makes common project commands easier to remember and run.

---

## 7. Application States

The main application should behave like a state machine.

Possible app states:

```go
type AppState int

const (
    StateMainMenu AppState = iota
    StatePlaying
    StatePaused
    StateGameOver
    StateSettings
    StateHelp
)
```

The root Bubble Tea model controls which screen is currently active.

Example flow:

```text
Start App
   ↓
Main Menu
   ↓
Select Game
   ↓
Playing
   ↓
Pause / Game Over / Return to Menu
```

---

## 7. Core Game Interface

Every single-player game should implement a common interface.

```go
type Game interface {
    Name() string
    Description() string
    Init() tea.Cmd
    Update(msg tea.Msg) (Game, tea.Cmd)
    View() string
    Reset() Game
    IsGameOver() bool
    Score() int
}
```

This allows the main application to treat all single-player games the same way.

Example usage:

```go
var currentGame engine.Game

currentGame = snake.New()
currentGame, cmd = currentGame.Update(msg)
view := currentGame.View()
```

For multiplayer or co-op games, use a separate interface that includes player information and session-level input.

```go
type PlayerID string

type Player struct {
    ID       PlayerID
    Name     string
    IsRemote bool
}

type PlayerInput struct {
    PlayerID PlayerID
    Key      string
}

type MultiplayerGame interface {
    Name() string
    Description() string
    MinPlayers() int
    MaxPlayers() int
    AddPlayer(player Player) error
    RemovePlayer(playerID PlayerID) error
    Init() tea.Cmd
    Update(msg tea.Msg) (MultiplayerGame, tea.Cmd)
    HandlePlayerInput(input PlayerInput) tea.Cmd
    ViewForPlayer(playerID PlayerID) string
    Reset() MultiplayerGame
    IsGameOver() bool
    Scores() map[PlayerID]int
}
```

A multiplayer game may show the same view to all players, or it may render a different view per player using `ViewForPlayer`.

---

## 8. Bubble Tea Architecture

Bubble Tea applications are based on three main methods:

```go
Init() tea.Cmd
Update(msg tea.Msg) (tea.Model, tea.Cmd)
View() string
```

For this project, there will be:

- One root app model.
- One menu model.
- One model per game.

The root model decides whether input should go to the menu, the active game, the pause screen, or another screen.

---

## 9. Root App Model Example

```go
type Model struct {
    state       AppState
    menu        menu.Model
    currentGame engine.Game
    width       int
    height      int
}
```

Responsibilities of the root model:

- Store current app state.
- Store terminal size.
- Route messages to the correct screen.
- Start selected games.
- Handle global keyboard shortcuts.

Suggested global shortcuts:

| Key   | Action               |
| ----- | -------------------- |
| `q`   | Quit app             |
| `esc` | Back / pause         |
| `p`   | Pause game           |
| `r`   | Restart current game |
| `m`   | Return to main menu  |
| `?`   | Show help            |

---

## 10. Game Loop and Ticks

Most arcade games need repeated updates over time. In Bubble Tea, this can be done with custom tick messages.

Example tick message:

```go
type TickMsg time.Time

func Tick(d time.Duration) tea.Cmd {
    return tea.Tick(d, func(t time.Time) tea.Msg {
        return TickMsg(t)
    })
}
```

Each game can define its own speed:

```go
const SnakeTickRate = 120 * time.Millisecond
const PongTickRate = 30 * time.Millisecond
const TetrisTickRate = 500 * time.Millisecond
```

The game update function receives the tick and moves the game forward.

```go
case engine.TickMsg:
    m.updateWorld()
    return m, engine.Tick(SnakeTickRate)
```

---

## 11. Shared Engine Components

The `internal/engine` package should contain reusable utilities.

### Grid

A shared grid type can help with Snake, Tetris, Minesweeper, Pac-Man, and similar games.

```go
type Point struct {
    X int
    Y int
}

type Grid struct {
    Width  int
    Height int
    Cells  [][]rune
}
```

Useful methods:

```go
func NewGrid(width, height int) Grid
func (g Grid) InBounds(p Point) bool
func (g *Grid) Set(p Point, r rune)
func (g Grid) Get(p Point) rune
func (g *Grid) Clear()
func (g Grid) Render() string
```

### Collision

Shared collision helpers:

```go
func PointEquals(a, b Point) bool
func ContainsPoint(points []Point, target Point) bool
func RectsOverlap(a, b Rect) bool
```

### Score

Reusable score structure:

```go
type Score struct {
    Current int
    High    int
}
```

### Input

Common input mapping:

```go
type Direction int

const (
    Up Direction = iota
    Down
    Left
    Right
)
```

---

## 12. Game Registration System

Instead of hardcoding all games directly into the menu, use a registry.

```go
type GameFactory func() engine.Game
type MultiplayerGameFactory func() engine.MultiplayerGame

type GameMode int

const (
    ModeSinglePlayer GameMode = iota
    ModeLocalMultiplayer
    ModeRemoteMultiplayer
    ModeCoop
)

type GameInfo struct {
    ID          string
    Name        string
    Description string
    Modes       []GameMode
    Factory     GameFactory
    MultiFactory MultiplayerGameFactory
}

var Registry = []GameInfo{
    {
        ID:          "snake",
        Name:        "Snake",
        Description: "Eat food, grow longer, and avoid crashing.",
        Modes:       []GameMode{ModeSinglePlayer},
        Factory:     snake.New,
    },
    {
        ID:          "pong",
        Name:        "Pong",
        Description: "Classic paddle and ball arcade game.",
        Modes:       []GameMode{ModeSinglePlayer, ModeLocalMultiplayer, ModeRemoteMultiplayer},
        Factory:     pong.New,
        MultiFactory: pong.NewMultiplayer,
    },
    {
        ID:          "tron",
        Name:        "Tron Light Cycles",
        Description: "Multiplayer trail survival game.",
        Modes:       []GameMode{ModeLocalMultiplayer, ModeRemoteMultiplayer},
        MultiFactory: tron.NewMultiplayer,
    },
}
```

The menu can read from this registry and display all available games automatically.

Games that support remote SSH play should declare `ModeRemoteMultiplayer` or `ModeCoop`.

---

## 13. Main Menu Design

The main menu should be simple and fast to use.

Example layout:

```text
╔════════════════════════════════╗
║        TERMINAL ARCADE         ║
╚════════════════════════════════╝

> Snake
  Pong
  Breakout
  Tetris
  Space Invaders

Use ↑/↓ to move, Enter to select, q to quit.
```

Menu model fields:

```go
type Model struct {
    choices []games.GameInfo
    cursor  int
}
```

Menu controls:

| Key          | Action           |
| ------------ | ---------------- |
| `up` / `k`   | Move cursor up   |
| `down` / `j` | Move cursor down |
| `enter`      | Select game      |
| `q`          | Quit             |

---

## 14. Rendering Strategy

Terminal rendering should be consistent across games.

Recommended style:

- Use a fixed-size playfield where possible.
- Use Unicode box drawing characters for borders.
- Use simple ASCII or Unicode characters for sprites.
- Use Lip Gloss for colors and layout.

Example characters:

| Object      | Character  |
| ----------- | ---------- |
| Wall        | `█`        |
| Empty cell  | ` `        |
| Snake body  | `●`        |
| Snake food  | `◆`        |
| Paddle      | `┃` or `━` |
| Ball        | `●`        |
| Brick       | `▣`        |
| Player ship | `▲`        |
| Enemy       | `♟` or `W` |

Recommended render helper:

```go
func RenderBox(title string, body string) string
```

Example game screen:

```text
Snake                         Score: 120
┌──────────────────────────────┐
│                              │
│        ● ● ● ◆               │
│            ●                 │
│                              │
└──────────────────────────────┘

Controls: arrows/WASD move · p pause · r restart · esc menu
```

---

## 15. Styling with Lip Gloss

Create a central style package.

```go
package styles

import "github.com/charmbracelet/lipgloss"

var Title = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("205")).
    Padding(1, 2)

var Border = lipgloss.NewStyle().
    Border(lipgloss.RoundedBorder()).
    Padding(1, 2)

var Selected = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("82"))

var Muted = lipgloss.NewStyle().
    Foreground(lipgloss.Color("240"))
```

Keep styles reusable and avoid defining colors directly inside each game.

---

## 16. Controls

Use consistent controls when possible.

| Key           | Action                |
| ------------- | --------------------- |
| `w` / `up`    | Move up               |
| `s` / `down`  | Move down             |
| `a` / `left`  | Move left             |
| `d` / `right` | Move right            |
| `space`       | Action / shoot / jump |
| `p`           | Pause                 |
| `r`           | Restart               |
| `esc`         | Back to menu or pause |
| `q`           | Quit                  |

Some games may need special controls, but the basic scheme should stay familiar.

---

## 17. Example Game: Snake

### Snake state

```go
type Model struct {
    width     int
    height    int
    snake     []engine.Point
    direction engine.Direction
    nextDir   engine.Direction
    food      engine.Point
    score     int
    gameOver  bool
}
```

### Snake rules

- Snake moves one cell per tick.
- Eating food increases score and grows the snake.
- Crashing into a wall ends the game.
- Crashing into itself ends the game.
- Player cannot instantly reverse direction.

### Snake files

```text
internal/games/snake/
├── model.go      # Snake model struct and constructor
├── update.go     # Input handling and game tick logic
├── view.go       # Rendering
└── snake.go      # Public New() function and metadata
```

---

## 18. Example Game: Pong

### Pong state

```go
type Model struct {
    width       int
    height      int
    leftPaddle  int
    rightPaddle int
    ballX       float64
    ballY       float64
    ballVX      float64
    ballVY      float64
    leftScore   int
    rightScore  int
    gameOver    bool
}
```

### Pong rules

- Ball moves every tick.
- Ball bounces off top and bottom walls.
- Ball bounces off paddles.
- Missing the ball gives the opponent a point.
- First player or CPU to a target score wins.

Possible modes:

- Player vs CPU
- Player vs Player

---

## 19. Example Game: Tetris

Tetris is more complex than Snake or Pong, so it should be added later.

### Tetris state

```go
type Model struct {
    board       [][]int
    activePiece Piece
    nextPiece   Piece
    score       int
    lines       int
    level       int
    gameOver    bool
}
```

### Tetris rules

- Piece falls on each tick.
- Player can move piece left/right.
- Player can rotate piece.
- Completed rows are cleared.
- Falling speed increases with level.
- Game ends when pieces reach the top.

---

## 20. Score and High Scores

At first, scores can exist only during one app session.

Later, high scores can be saved to a local JSON file.

Suggested file:

```text
~/.terminal-arcade/scores.json
```

Example structure:

```json
{
  "snake": {
    "high_score": 420
  },
  "pong": {
    "high_score": 7
  },
  "tetris": {
    "high_score": 12000
  }
}
```

Score service:

```go
type Store interface {
    Load(gameID string) (int, error)
    Save(gameID string, score int) error
}
```

---

## 21. Settings

Possible settings:

- Theme
- Sound on/off if terminal bell is used
- Difficulty
- Game speed
- Player name
- High scores reset
- SSH server port
- SSH host key path
- Maximum remote players
- Public lobby enabled/disabled

Settings can be stored in:

```text
~/.terminal-arcade/config.json
```

Example:

```json
{
  "theme": "default",
  "difficulty": "normal",
  "show_help": true,
  "ssh": {
    "enabled": true,
    "host": "0.0.0.0",
    "port": 23234,
    "host_key_path": "./.ssh/arcade_ed25519",
    "max_players": 16,
    "public_lobby": true
  }
}
```

---

## 22. SSH Remote Play with Wish

The project should support remote players through an SSH server built with the **Wish** library.

Wish is useful because it allows terminal applications to be served over SSH. Players only need an SSH client. They do not need to install the game locally.

### SSH server responsibilities

The SSH server should:

- Accept remote terminal connections.
- Start a Bubble Tea app for each connected user.
- Assign a unique player ID to every SSH session.
- Allow players to enter a lobby.
- Allow players to create or join multiplayer rooms.
- Route player input to the correct multiplayer game session.
- Disconnect players cleanly when they leave.

### Suggested server structure

```text
internal/server/
├── ssh.go       # Starts Wish SSH server
├── session.go   # Tracks connected users
├── lobby.go     # Room and matchmaking logic
└── auth.go      # Optional public key/password auth
```

### Example SSH server entry point

```go
package main

import (
    "context"
    "errors"
    "fmt"
    "net"
    "os"
    "os/signal"
    "syscall"
    "time"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/wish"
    wishtea "github.com/charmbracelet/wish/bubbletea"
    "github.com/charmbracelet/wish/logging"

    "github.com/your-username/terminal-arcade/internal/app"
)

func main() {
    host := "0.0.0.0"
    port := "23234"

    s, err := wish.NewServer(
        wish.WithAddress(net.JoinHostPort(host, port)),
        wish.WithHostKeyPath(".ssh/arcade_ed25519"),
        wish.WithMiddleware(
            wishtea.Middleware(func(sess wish.Session) (tea.Model, []tea.ProgramOption) {
                username := sess.User()

                model := app.NewRemote(app.RemoteOptions{
                    Username: username,
                    SessionID: sess.RemoteAddr().String(),
                })

                return model, []tea.ProgramOption{
                    tea.WithAltScreen(),
                }
            }),
            logging.Middleware(),
        ),
    )
    if err != nil {
        fmt.Println("Could not start SSH server:", err)
        os.Exit(1)
    }

    done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGTERM)

    go func() {
        if err := s.ListenAndServe(); err != nil && !errors.Is(err, wish.ErrServerClosed) {
            fmt.Println("SSH server error:", err)
            os.Exit(1)
        }
    }()

    fmt.Printf("Arcade SSH server running on %s:%s
", host, port)

    <-done

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := s.Shutdown(ctx); err != nil {
        fmt.Println("SSH server shutdown error:", err)
    }
}
```

This should live in:

```text
cmd/arcade-server/main.go
```

Local play and SSH server play should have separate entry points:

```text
cmd/arcade/main.go          # Local terminal app
cmd/arcade-server/main.go   # Remote SSH server
```

---

## 23. Remote Player Sessions

Each SSH connection should become a player session.

```go
type SessionID string

type RemoteSession struct {
    ID       SessionID
    PlayerID engine.PlayerID
    Username string
    RoomID   string
}
```

The session manager tracks active users.

```go
type SessionManager struct {
    sessions map[SessionID]RemoteSession
}
```

Useful methods:

```go
func (m *SessionManager) Add(session RemoteSession)
func (m *SessionManager) Remove(id SessionID)
func (m *SessionManager) Get(id SessionID) (RemoteSession, bool)
func (m *SessionManager) List() []RemoteSession
```

Each remote player should be able to:

- Set a display name.
- Enter the public lobby.
- Create a room.
- Join an existing room.
- Leave a room.
- Start a game when enough players are ready.

---

## 24. Multiplayer Lobby and Rooms

Remote co-op games need a lobby system.

### Lobby responsibilities

- Show available rooms.
- Show player counts.
- Show which game each room is running.
- Allow room creation.
- Allow joining and leaving rooms.
- Allow ready checks.

Example lobby screen:

```text
╔════════════════════════════════╗
║         TERMINAL ARCADE        ║
║            SSH LOBBY           ║
╚════════════════════════════════╝

Rooms:

> Room 1  | Pong              | 1/2 players | Waiting
  Room 2  | Tron Light Cycles | 3/4 players | Waiting
  Room 3  | Co-op Shooter     | 2/4 players | In Game

n: new room · enter: join · r: refresh · q: quit
```

### Room state

```go
type RoomID string

type RoomState int

const (
    RoomWaiting RoomState = iota
    RoomPlaying
    RoomFinished
)

type Room struct {
    ID        RoomID
    GameID    string
    State     RoomState
    Players   []engine.Player
    Game      engine.MultiplayerGame
    CreatedAt time.Time
}
```

### Room manager

```go
type RoomManager struct {
    rooms map[RoomID]*Room
}
```

Useful methods:

```go
func (m *RoomManager) CreateRoom(gameID string, owner engine.Player) (*Room, error)
func (m *RoomManager) JoinRoom(roomID RoomID, player engine.Player) error
func (m *RoomManager) LeaveRoom(roomID RoomID, playerID engine.PlayerID) error
func (m *RoomManager) StartRoom(roomID RoomID) error
func (m *RoomManager) ListRooms() []*Room
```

---

## 25. Multiplayer Input Routing

In a local Bubble Tea game, keyboard input belongs to one user. In SSH multiplayer, each remote terminal connection has its own input stream.

The app must convert each keypress into a player-specific input message.

```go
type PlayerInputMsg struct {
    PlayerID engine.PlayerID
    Key      string
}
```

When a remote player presses a key:

```go
case tea.KeyMsg:
    return m, func() tea.Msg {
        return engine.PlayerInputMsg{
            PlayerID: m.playerID,
            Key:      msg.String(),
        }
    }
```

The room or multiplayer game receives this message and updates only that player's state.

Example:

```go
func (g Model) HandlePlayerInput(input engine.PlayerInput) tea.Cmd {
    switch input.PlayerID {
    case g.leftPlayerID:
        g.handleLeftPlayer(input.Key)
    case g.rightPlayerID:
        g.handleRightPlayer(input.Key)
    }

    return nil
}
```

---

## 26. Shared Multiplayer Game State

For co-op games, multiple players need to interact with the same game world.

Important rule:

```text
The room owns the shared multiplayer game state.
Individual SSH sessions should not each create their own copy of the game.
```

Recommended model:

```text
SSH Session App
    ↓ sends input
Room Manager
    ↓ routes input
Shared Multiplayer Game Instance
    ↓ produces updated views
All connected players render latest state
```

Because multiple SSH sessions can send input at the same time, shared room state should be protected.

```go
type Room struct {
    ID      RoomID
    Game    engine.MultiplayerGame
    Players []engine.Player
    mu      sync.Mutex
}
```

Use the mutex when joining, leaving, updating, or reading the shared game state.

```go
room.mu.Lock()
defer room.mu.Unlock()

room.Game.HandlePlayerInput(input)
```

For a cleaner design, each room can also run its own event loop through a channel.

```go
type RoomEvent interface{}

type RoomInputEvent struct {
    PlayerID engine.PlayerID
    Key      string
}

type RoomTickEvent struct{}

type Room struct {
    Events chan RoomEvent
}
```

This avoids many direct locks because all state changes happen in one goroutine.

---

## 27. Rendering Remote Multiplayer Views

There are two common rendering models.

### Same view for all players

Works well for:

- Pong
- Tron Light Cycles
- Bomberman
- Co-op shooter
- Shared arena games

```go
func (g Model) ViewForPlayer(playerID engine.PlayerID) string {
    return g.RenderSharedArena()
}
```

### Different view per player

Works well for:

- Hidden-information games
- Split-screen dungeon games
- Games with private inventory
- Games with role-specific UI

```go
func (g Model) ViewForPlayer(playerID engine.PlayerID) string {
    player := g.players[playerID]
    return g.RenderPlayerPerspective(player)
}
```

For the first version of the project, prefer same-view games. They are much easier to implement and debug.

---

## 28. Recommended Co-op and Remote Multiplayer Games

Good games for SSH multiplayer:

| Game                 | Mode        | Player Count | Difficulty |
| -------------------- | ----------- | ------------ | ---------- |
| Pong                 | PvP         | 2            | Easy       |
| Tron Light Cycles    | PvP         | 2-4          | Medium     |
| Bomberman-style game | PvP / Co-op | 2-4          | Medium     |
| Co-op Space Shooter  | Co-op       | 2-4          | Medium     |
| Tank Battle          | PvP / Teams | 2-4          | Medium     |
| Maze Escape          | Co-op       | 2-4          | Medium     |
| ASCII Racing         | PvP         | 2-4          | Medium     |
| Survival Arena       | Co-op       | 2-4          | Medium     |
| Tetris Battle        | PvP         | 2            | Hard       |
| Pac-Man vs Ghosts    | Asymmetric  | 2-5          | Hard       |

Recommended first remote multiplayer game: **Pong**.

Recommended first true co-op game: **Co-op Space Shooter** or **Survival Arena**.

---

## 29. Authentication and Access

For development, the SSH server can allow simple access on localhost.

For public hosting, consider:

- Public key authentication.
- Password authentication.
- Maximum connection limit.
- Per-IP rate limiting.
- Room size limits.
- Idle timeout.
- Spectator mode instead of unlimited players.

Possible config:

```json
{
  "ssh": {
    "auth_mode": "public_key",
    "allowed_keys_path": "./authorized_keys",
    "idle_timeout_seconds": 300
  }
}
```

A public arcade server should not allow unlimited anonymous users without limits.

---

## 30. Deployment Options

The SSH arcade server can be deployed on:

- VPS
- Raspberry Pi
- Home server with port forwarding
- Fly.io
- Railway
- Render
- Docker on any Linux server

Example command:

```bash
go run ./cmd/arcade-server
```

Example connection:

```bash
ssh your-domain.com -p 23234
```

### Docker idea

```dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o arcade-server ./cmd/arcade-server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/arcade-server .
EXPOSE 23234
CMD ["./arcade-server"]
```

---

## 31. Error Handling

The app should avoid crashing when possible.

Recommended rules:

- If terminal size is too small, show a warning screen.
- If score file cannot be read, continue with empty scores.
- If score file cannot be written, show a non-blocking warning.
- If a game panics during development, recover at the app level if possible.
- If an SSH player disconnects, remove them from their room.
- If a multiplayer room becomes empty, delete it.
- If a host disconnects, transfer ownership or close the room.
- If a player disconnects during a co-op game, pause briefly or let remaining players continue.

Small terminal warning:

```text
Terminal too small.
Please resize to at least 80x24.
```

---

## 22. Error Handling

The app should avoid crashing when possible.

Recommended rules:

- If terminal size is too small, show a warning screen.
- If score file cannot be read, continue with empty scores.
- If score file cannot be written, show a non-blocking warning.
- If a game panics during development, recover at the app level if possible.

Small terminal warning:

```text
Terminal too small.
Please resize to at least 80x24.
```

---

## 32. Testing Strategy

Game logic should be testable without rendering.

Test pure functions such as:

- Collision detection
- Movement
- Score updates
- Grid boundaries
- Piece rotation in Tetris
- Ball collision in Pong
- Menu selection logic

Example test layout:

```text
internal/games/snake/snake_test.go
internal/games/pong/pong_test.go
internal/engine/collision_test.go
internal/engine/grid_test.go
```

Example test:

```go
func TestSnakeDiesWhenHittingWall(t *testing.T) {
    game := New()
    game.snake[0] = engine.Point{X: 0, Y: 0}
    game.direction = engine.Left

    game.step()

    if !game.gameOver {
        t.Fatal("expected game over after hitting wall")
    }
}
```

---

## 33. Development Roadmap

### Phase 1: Project foundation

- Create Go module.
- Add Bubble Tea and Lip Gloss.
- Build root app model.
- Build main menu.
- Add app states.
- Add shared styles.

### Phase 1.5: SSH foundation

- Add Wish dependency.
- Create `cmd/arcade-server` entry point.
- Start basic SSH server.
- Launch Bubble Tea app per SSH session.
- Add remote player identity.
- Add lobby screen placeholder.

### Phase 2: Shared engine

- Add tick system.
- Add grid utilities.
- Add point and direction types.
- Add collision helpers.
- Add shared score type.

### Phase 3: First playable game

Recommended first game: **Snake**.

Tasks:

- Create Snake model.
- Handle movement input.
- Implement food spawning.
- Implement collision.
- Render playfield.
- Add score and game-over screen.

### Phase 4: Add more simple games

Recommended order:

1. Snake
2. Pong
3. Breakout
4. Flappy Bird
5. Whack-a-Mole
6. ASCII Racing

### Phase 4.5: First remote multiplayer mode

Recommended first game: **Pong over SSH**.

Tasks:

- Add room manager.
- Add player sessions.
- Add two-player remote Pong.
- Route SSH key input by player ID.
- Render shared game state to both players.
- Handle disconnects.

### Phase 5: Add more complex games

Recommended order:

1. Tetris
2. Space Invaders
3. Tron Light Cycles
4. Minesweeper
5. 2048
6. Pac-Man-style maze game

### Phase 6: Persistence and polish

- Save high scores.
- Add settings.
- Add help screen.
- Add themes.
- Add better game-over screens.
- Add animations and transitions.

### Phase 7: Multiplayer polish

- Add public lobby.
- Add private rooms.
- Add spectator mode.
- Add player names.
- Add ready system.
- Add reconnect handling.
- Add server config file.
- Add simple moderation/admin commands.

---

## 34. Suggested First Milestone

The first milestone should be:

```text
A working terminal app with:

- Main menu
- Snake game
- Pause
- Restart
- Return to menu
- Quit
```

Once this is complete, adding new games becomes much easier because the architecture is already prepared.

---

## 35. Example Initial Commands

```bash
mkdir terminal-arcade
cd terminal-arcade

go mod init github.com/your-username/terminal-arcade

go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
go get github.com/charmbracelet/bubbles
go get github.com/charmbracelet/wish
go get github.com/charmbracelet/wish/bubbletea
go get github.com/charmbracelet/wish/logging

mkdir -p cmd/arcade
mkdir -p cmd/arcade-server
mkdir -p internal/{app,engine,menu,lobby,server,styles,games}
mkdir -p internal/games/{snake,pong,breakout,tetris,tron}
mkdir -p docs
mkdir -p .ssh
```

Generate a local SSH host key for development:

```bash
ssh-keygen -t ed25519 -f .ssh/arcade_ed25519 -N ""
```

Run local app:

```bash
go run ./cmd/arcade
```

Run SSH server:

```bash
go run ./cmd/arcade-server
```

Connect locally:

```bash
ssh localhost -p 23234
```

---

## 36. Example `main.go`

```go
package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/your-username/terminal-arcade/internal/app"
)

func main() {
    p := tea.NewProgram(
        app.New(),
        tea.WithAltScreen(),
    )

    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}
```

---

## 37. Design Principles

### Keep games independent

Each game should own its own rules and state. The root app should not know how Snake, Pong, or Tetris work internally.

### Keep shared logic generic

Only move code into `engine` if at least two games can use it.

### Prefer simple rendering

Terminal games should be readable and responsive. Avoid overly complex visuals at the beginning.

### Separate game logic from UI

A game step should be testable without Bubble Tea or Lip Gloss.

For example:

```go
func (m *Model) Step() {
    // update game state only
}
```

Then Bubble Tea simply calls it:

```go
case engine.TickMsg:
    m.Step()
    return m, engine.Tick(SnakeTickRate)
```

### Build one game fully before adding many games

It is better to finish Snake completely than to have five unfinished prototypes.

### Treat multiplayer as a separate architecture layer

Do not force every game to be multiplayer. Single-player and multiplayer games can share utilities, but multiplayer needs extra concepts like players, sessions, rooms, synchronization, and disconnect handling.

### Keep shared multiplayer state in rooms

Remote SSH sessions should send input to a shared room. They should not each run independent copies of the multiplayer game.

### Start remote play with simple same-screen games

Pong, Tron, and co-op arena games are easier than games where every player needs a private view.

---

## 38. Possible Future Features

- Local multiplayer
- Online leaderboard
- Replay system
- Custom themes
- Configurable controls
- Save states
- Achievements
- Game statistics
- Terminal sound effects using bell character
- Mouse support for selected games
- AI opponents
- Procedural maps
- Challenge modes
- SSH remote multiplayer
- Public arcade server
- Private invite-only rooms
- Spectator mode
- Co-op campaigns
- Server-side leaderboards

---

## 39. Summary

This project should be structured as a modular Bubble Tea application with a root app model, a menu system, shared engine utilities, independent game packages, and an optional Wish-powered SSH server.

The most important architectural idea is that every game should implement a common interface. Single-player games can use a simple `Game` interface, while co-op and remote games can use a `MultiplayerGame` interface.

For remote play, each SSH connection becomes a player session. Players enter a lobby, join rooms, and send input to a shared multiplayer game instance owned by the room.

Start with local Snake to prove the basic architecture. Then add local Pong. After that, use Pong as the first remote SSH multiplayer game because it is simple, visual, and naturally supports two players.

