package menu

import (
	"github.com/barttob/terminal-arcade/internal/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	// var b strings.Builder

	options := []string{}
	for _, entry := range gameRegistry {
		options = append(options, "Play "+entry.Name)
	}
	options = append(options, "Settings", "Exit")

	var body string

	for i, option := range options {
		prefix := "  "
		text := option

		if m.cursor == i {
			prefix = "> "
			text = styles.Selected.Render(option)
		}

		body += prefix + text + "\n"
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		styles.Title.Width(m.terminalWidth/2).Render("TERMINAL ARCADE"),
		styles.Menu.Width(m.terminalWidth/2).Render(body),
		styles.Help.Render("↑/↓ move • enter select • q quit"),
	)
}
