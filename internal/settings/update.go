package settings

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/barttob/terminal-arcade/internal/engine"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	rows := m.current().Rows
	n := len(m.sections)

	switch key.String() {
	case "tab":
		m.section = (m.section + 1) % n
		m.cursor = 0
	case "shift+tab":
		m.section = (m.section - 1 + n) % n
		m.cursor = 0
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(rows) {
			m.cursor++
		}
	case "left", "h":
		if m.cursor < len(rows) {
			rows[m.cursor].Adjust(&m.values, -1)
		}
	case "right", "l":
		if m.cursor < len(rows) {
			rows[m.cursor].Adjust(&m.values, 1)
		}
	case "enter", " ":
		id := m.current().GameID
		return m, func() tea.Msg { return engine.StartGameMsg{GameID: id} }
	}

	return m, nil
}
