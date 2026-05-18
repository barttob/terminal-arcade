package snake

import (
	"time"

	"github.com/barttob/terminal-arcade/internal/engine"
	tea "github.com/charmbracelet/bubbletea"
)

const ID = "snake"
const DefaultWidth = 30
const DefaultHeight = 15

const DefaultTickRate = 120 * time.Millisecond

func (m *Model) Name() string {
	return "Snake"
}

func (m *Model) Description() string {
	return "A classic snake game where you control a snake to eat food and grow longer."
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
