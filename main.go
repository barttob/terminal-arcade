package main

import (
	"fmt"
	"os"

	"terminal-arcade/games"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type screen int

const (
	screenMenu screen = iota
	screenSettings
	screenGame
)

type appModel struct {
	screen screen
	cursor int

	width  int
	height int
	speed  int

	game games.SnakeModel
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			Border(lipgloss.DoubleBorder()).
			Padding(1, 4).
			Align(lipgloss.Center)

	menuStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 3).
			Foreground(lipgloss.Color("230"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("120")).
			Bold(true)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

func initialApp() appModel {
	return appModel{
		screen: screenMenu,
		width:  30,
		height: 15,
		speed:  120,
	}
}

func (m appModel) Init() tea.Cmd {
	return nil
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.screen == screenGame {
		updatedGame, cmd := m.game.Update(msg)
		m.game = updatedGame.(games.SnakeModel)

		if m.game.ExitToMenu {
			m.screen = screenMenu
			m.cursor = 0
			return m, nil
		}

		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	switch m.screen {
	case screenMenu:
		return updateMenu(m, msg)
	case screenSettings:
		return updateSettings(m, msg)
	}

	return m, nil
}

func updateMenu(m appModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	options := []string{"Play Snake", "Snake Settings", "Exit"}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(options)-1 {
				m.cursor++
			}

		case "enter":
			switch m.cursor {
			case 0:
				m.game = games.NewSnakeModel(m.width, m.height, m.speed)
				m.screen = screenGame
				return m, m.game.Init()

			case 1:
				m.screen = screenSettings
				m.cursor = 0

			case 2:
				return m, tea.Quit
			}

		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func updateSettings(m appModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	optionsCount := 4

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			m.screen = screenMenu
			m.cursor = 0

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < optionsCount-1 {
				m.cursor++
			}

		case "left", "h":
			switch m.cursor {
			case 0:
				if m.width > 15 {
					m.width -= 5
				}
			case 1:
				if m.height > 10 {
					m.height -= 5
				}
			case 2:
				if m.speed < 300 {
					m.speed += 20
				}
			}

		case "right", "l":
			switch m.cursor {
			case 0:
				if m.width < 60 {
					m.width += 5
				}
			case 1:
				if m.height < 30 {
					m.height += 5
				}
			case 2:
				if m.speed > 40 {
					m.speed -= 20
				}
			}

		case "enter":
			if m.cursor == 3 {
				m.screen = screenMenu
				m.cursor = 0
			}
		}
	}

	return m, nil
}

func (m appModel) View() string {
	switch m.screen {
	case screenMenu:
		return center(renderMenu(m))
	case screenSettings:
		return center(renderSettings(m))
	case screenGame:
		return m.game.View()
	default:
		return ""
	}
}

func renderMenu(m appModel) string {
	options := []string{"Play Snake", "Snake Settings", "Exit"}

	var body string

	for i, option := range options {
		prefix := "  "
		text := option

		if m.cursor == i {
			prefix = "➜ "
			text = selectedStyle.Render(option)
		}

		body += prefix + text + "\n"
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		titleStyle.Render("SNAKE TERMINAL"),
		menuStyle.Render(body),
		helpStyle.Render("↑/↓ move • enter select • q quit"),
	)
}

func renderSettings(m appModel) string {
	rows := []string{
		fmt.Sprintf("Board width:  %d", m.width),
		fmt.Sprintf("Board height: %d", m.height),
		fmt.Sprintf("Speed:        %dms", m.speed),
		"Back",
	}

	var body string

	for i, row := range rows {
		prefix := "  "
		text := row

		if m.cursor == i {
			prefix = "➜ "
			text = selectedStyle.Render(row)
		}

		body += prefix + text + "\n"
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		titleStyle.Render("SETTINGS"),
		menuStyle.Render(body),
		helpStyle.Render("←/→ change • esc back"),
	)
}

func center(s string) string {
	return lipgloss.NewStyle().
		Width(80).
		Height(24).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Render(s)
}

func main() {
	p := tea.NewProgram(initialApp(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
