package minesweeper

import (
	"github.com/barttob/terminal-arcade/internal/engine"
	tea "github.com/charmbracelet/bubbletea"
)

const ID = "minesweeper"
const DefaultWidth = 16
const DefaultHeight = 16
const DefaultMines = 40

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

func New() engine.Game {
	return NewModel(DefaultWidth, DefaultHeight, DefaultMines)
}
