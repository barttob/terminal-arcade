package menu

import (
	"github.com/barttob/terminal-arcade/internal/engine"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)+1 {
				m.cursor++
			}
		case "enter":
			n := len(m.choices)
			switch {
			case m.cursor < n:
				id := m.choices[m.cursor].ID
				return m, func() tea.Msg { return engine.StartGameMsg{GameID: id} }
			case m.cursor == n:
				return m, func() tea.Msg { return engine.OpenSettingsMsg{} }
			case m.cursor == n+1:
				return m, tea.Quit
			}

		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}
