package menu

import (
	"github.com/barttob/terminal-arcade/internal/games"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	choices []games.Entry
	cursor  int

	terminalWidth  int
	terminalHeight int
}

func MenuModel() Model {
	return Model{
		choices: games.Registry,
		cursor:  0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
