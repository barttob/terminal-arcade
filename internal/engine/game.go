package engine

import tea "github.com/charmbracelet/bubbletea"

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
