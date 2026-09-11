package app

import (
	"github.com/barttob/terminal-arcade/internal/menu"
	tea "github.com/charmbracelet/bubbletea"
)

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
	case menu.GameSelectedMsg:
		m.currentGame = msg.Game
		m.state = StatePlaying
		return m, m.currentGame.Init()
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	if m.state == StateMainMenu {
		updatedMenu, cmd := m.menu.Update(msg)
		m.menu = updatedMenu.(menu.Model)
		return m, cmd
	}

	if m.state == StatePlaying {
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
			m.currentGame = nil
			m.state = StateMainMenu
			return m, nil
		}

		updatedGame, cmd := m.currentGame.Update(msg)
		m.currentGame = updatedGame
		return m, cmd
	}

	return m, nil
}
