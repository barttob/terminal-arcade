package snake

import (
	"time"

	"github.com/barttob/terminal-arcade/internal/engine"
	tea "github.com/charmbracelet/bubbletea"
)

const ID = "snake"

// Board size limits; the view adds 5 rows and 3 columns around a board drawn
// two characters per cell, so the maximum fits an 80×24 terminal.
const (
	MinSize   = 8
	MaxWidth  = 38
	MaxHeight = 19
)

// FoodScore is the number of points each piece of food is worth.
const FoodScore = 10

type Speed struct {
	Name     string
	TickRate time.Duration
}

var Speeds = []Speed{
	{Name: "Slow", TickRate: 180 * time.Millisecond},
	{Name: "Normal", TickRate: 120 * time.Millisecond},
	{Name: "Fast", TickRate: 80 * time.Millisecond},
	{Name: "Insane", TickRate: 50 * time.Millisecond},
}

type Config struct {
	Width  int
	Height int
	// Speed indexes Speeds.
	Speed int
	// Walls ends the game at the board edge; without them the snake wraps.
	Walls bool
}

var DefaultConfig = Config{Width: 30, Height: 15, Speed: 1, Walls: true}

// Clamped returns c with the board size and speed forced into valid ranges.
func (c Config) Clamped() Config {
	c.Width = max(MinSize, min(c.Width, MaxWidth))
	c.Height = max(MinSize, min(c.Height, MaxHeight))
	c.Speed = max(0, min(c.Speed, len(Speeds)-1))
	return c
}

func (m *Model) Name() string {
	return "Snake"
}

func (m *Model) Description() string {
	return "A classic snake game where you control a snake to eat food and grow longer."
}

func (m *Model) Init() tea.Cmd {
	return engine.Tick(m.tickID, m.tickRate)
}

func (m *Model) Reset() engine.Game {
	return NewModel(m.width, m.height, m.tickRate, m.walls)
}

func (m *Model) IsGameOver() bool {
	return m.gameOver
}

func (m *Model) Score() int {
	return m.score
}

func New(config Config) engine.Game {
	config = config.Clamped()
	return NewModel(config.Width, config.Height, Speeds[config.Speed].TickRate, config.Walls)
}
