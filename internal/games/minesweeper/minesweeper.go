package minesweeper

import (
	"time"

	"github.com/barttob/terminal-arcade/internal/engine"
	tea "github.com/charmbracelet/bubbletea"
)

const ID = "minesweeper"
const DefaultWidth = 30
const DefaultHeight = 15

const DefaultTickRate = 120 * time.Millisecond

func (m *Model) Name() string {
	return "Minesweeper"
}

func (m *Model) Description() string {
	return "A classic minesweeper game where you uncover tiles to find mines and avoid them."
}

func (m *Model) Init() tea.Cmd {
	return engine.Tick(DefaultTickRate)
}

func (m *Model) Reset() engine.Game {
	return NewModel(m.width, m.height)
}

func (m *Model) IsGameOver() bool {
	return m.gameOver
}

func (m *Model) Score() int {
	return m.score
}

func New() engine.Game {
	return NewModel(DefaultWidth, DefaultHeight)
}
