package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/barttob/terminal-arcade/internal/engine"
	"github.com/barttob/terminal-arcade/internal/games"
	"github.com/barttob/terminal-arcade/internal/menu"
	"github.com/barttob/terminal-arcade/internal/settings"
)

type appModel struct {
	state    AppState
	menu     menu.Model
	settings settings.Model
	// lobby       lobby.Model
	currentGame engine.Game
	//    playerID    engine.PlayerID
	terminalWidth  int
	terminalHeight int

	remote bool
}

func New() appModel {
	return appModel{
		state:    StateMainMenu,
		menu:     menu.MenuModel(),
		settings: settings.New(settings.Default(), games.SettingsSections()),
	}
}

func (m appModel) Init() tea.Cmd {
	return nil
}

func center(s string, width, height int) string {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Render(s)
}
