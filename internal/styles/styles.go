package styles

import "github.com/charmbracelet/lipgloss"

var (
	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Border(lipgloss.DoubleBorder()).
		Padding(1, 4).
		Align(lipgloss.Center)
	Menu = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 3).
		Foreground(lipgloss.Color("230"))
	Selected = lipgloss.NewStyle().
			Foreground(lipgloss.Color("120")).
			Bold(true)
	Help = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Align(lipgloss.Center)
)
