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

	Tab = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))
	ActiveTab = lipgloss.NewStyle().
			Bold(true).
			Underline(true).
			Foreground(lipgloss.Color("205"))

	GameHeader = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))
	Board = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("241"))
	Win = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("120"))
	Lose = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("196"))

	// Minesweeper
	MineCursor = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("120"))
	MineHidden = lipgloss.NewStyle().
			Foreground(lipgloss.Color("244"))
	MineEmpty = lipgloss.NewStyle().
			Foreground(lipgloss.Color("238"))
	MineFlag = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("214"))
	MineWrongFlag = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("196"))
	Mine = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("196"))
	MineExploded = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("231")).
			Background(lipgloss.Color("196"))
	// MineNumbers is indexed by adjacent mine count (index 0 is unused).
	MineNumbers = [9]lipgloss.Style{
		lipgloss.NewStyle(),
		lipgloss.NewStyle().Foreground(lipgloss.Color("39")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("76")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("203")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("63")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("130")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("44")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("245")),
	}
)
