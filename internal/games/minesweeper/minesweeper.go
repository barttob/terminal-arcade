package minesweeper

import (
	"github.com/barttob/terminal-arcade/internal/engine"
	tea "github.com/charmbracelet/bubbletea"
)

const ID = "minesweeper"

// Board size limits; MaxHeight keeps the view within a 24-row terminal.
const (
	MinSize   = 5
	MaxWidth  = 30
	MaxHeight = 16
)

type Config struct {
	Width  int
	Height int
	Mines  int
}

// MaxMines leaves room for the 3×3 safe opening around the first reveal.
func (c Config) MaxMines() int {
	return c.Width*c.Height - 9
}

// Clamped returns c with the board size and mine count forced into valid ranges.
func (c Config) Clamped() Config {
	c.Width = max(MinSize, min(c.Width, MaxWidth))
	c.Height = max(MinSize, min(c.Height, MaxHeight))
	c.Mines = max(1, min(c.Mines, c.MaxMines()))
	return c
}

type Preset struct {
	Name   string
	Config Config
}

var Presets = []Preset{
	{Name: "Beginner", Config: Config{Width: 9, Height: 9, Mines: 10}},
	{Name: "Intermediate", Config: Config{Width: 16, Height: 16, Mines: 40}},
	{Name: "Expert", Config: Config{Width: 30, Height: 16, Mines: 99}},
}

var DefaultConfig = Presets[1].Config

func (m *Model) Name() string {
	return "Minesweeper"
}

func (m *Model) Description() string {
	return "A classic minesweeper game where you uncover tiles to find mines and avoid them."
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Reset() engine.Game {
	return NewModel(m.width, m.height, m.mines)
}

func (m *Model) IsGameOver() bool {
	return m.gameOver
}

// Score is the number of safe cells uncovered.
func (m *Model) Score() int {
	return m.revealed
}

func New(config Config) engine.Game {
	config = config.Clamped()
	return NewModel(config.Width, config.Height, config.Mines)
}
