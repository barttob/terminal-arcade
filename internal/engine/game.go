package engine

import tea "github.com/charmbracelet/bubbletea"

type Game interface {
	Name() string
	Description() string
	Init() tea.Cmd
	Update(msg tea.Msg) (Game, tea.Cmd)
	View() string
	Reset() Game
	IsGameOver() bool
	Score() int
}

// StartGameMsg asks the app to start a new game using the current settings.
type StartGameMsg struct {
	GameID string
}

// OpenSettingsMsg asks the app to show the settings screen, on GameID's tab
// if set.
type OpenSettingsMsg struct {
	GameID string
}
