package settings

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/barttob/terminal-arcade/internal/styles"
)

func (m Model) View() string {
	// Size the box to the widest tab so switching tabs doesn't resize it.
	width := 0
	for i := range m.sections {
		width = max(width, lipgloss.Width(m.body(i)))
	}
	box := styles.Menu.Width(width + styles.Menu.GetHorizontalPadding())

	return lipgloss.JoinVertical(
		lipgloss.Center,
		styles.Title.Render("SETTINGS"),
		box.Render(m.body(m.section)),
		styles.Help.Render("tab next game • ↑/↓ move • ←/→ change • enter play • esc back"),
	)
}

// body renders the tab bar and the rows of section i.
func (m Model) body(i int) string {
	section := m.sections[i]

	var tabs []string
	for j, s := range m.sections {
		if j == i {
			tabs = append(tabs, styles.ActiveTab.Render(s.Title))
		} else {
			tabs = append(tabs, styles.Tab.Render(s.Title))
		}
	}
	lines := []string{"  " + strings.Join(tabs, styles.Tab.Render(" │ ")), ""}

	if len(section.Rows) == 0 {
		lines = append(lines, styles.Help.Render("  No settings yet."))
	}
	for j, r := range section.Rows {
		selected := i == m.section && j == m.cursor

		value := "  " + r.Value(m.values)
		if selected {
			value = "‹ " + r.Value(m.values) + " ›"
		}

		// Fixed-width columns keep rows aligned as values change.
		line := fmt.Sprintf("%-10s %-16s", r.Label, value)
		if selected {
			lines = append(lines, "> "+styles.Selected.Render(line))
		} else {
			lines = append(lines, "  "+line)
		}
	}

	play := "▶ Play " + section.Title
	if i == m.section && m.cursor == len(section.Rows) {
		play = "> " + styles.Selected.Render(play)
	} else {
		play = "  " + play
	}
	lines = append(lines, "", play)

	return strings.Join(lines, "\n")
}
