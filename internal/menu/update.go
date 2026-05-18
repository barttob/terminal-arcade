package menu

import (
	"github.com/barttob/terminal-arcade/internal/engine"
	tea "github.com/charmbracelet/bubbletea"
)

type GameSelectedMsg struct {
	Game engine.Game
}

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
				game := m.choices[m.cursor].Factory()
				return m, func() tea.Msg { return GameSelectedMsg{Game: game} }
			case m.cursor == n:
				return m, tea.Quit
			case m.cursor == n+1:
				return m, tea.Quit
			}

		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}
