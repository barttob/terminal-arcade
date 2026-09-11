package app

import (
	"github.com/barttob/terminal-arcade/internal/engine"
	"github.com/barttob/terminal-arcade/internal/games"
	"github.com/barttob/terminal-arcade/internal/menu"
	"github.com/barttob/terminal-arcade/internal/settings"
	tea "github.com/charmbracelet/bubbletea"
)

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
	case engine.StartGameMsg:
		return m.startGame(msg.GameID)
	case engine.OpenSettingsMsg:
		m.settings = m.settings.Focus(msg.GameID)
		m.state = StateSettings
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	if m.state == StateMainMenu {
		updatedMenu, cmd := m.menu.Update(msg)
		m.menu = updatedMenu.(menu.Model)
		return m, cmd
	}

	if m.state == StateSettings {
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
			// Resume the game settings were opened from, if any.
			m.state = StateMainMenu
			if m.currentGame != nil {
				m.state = StatePlaying
			}
			return m, nil
		}

		updatedSettings, cmd := m.settings.Update(msg)
		m.settings = updatedSettings.(settings.Model)
		return m, cmd
	}

	if m.state == StatePlaying {
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
			m.currentGame = nil
			m.state = StateMainMenu
			return m, nil
		}

		updatedGame, cmd := m.currentGame.Update(msg)
		m.currentGame = updatedGame
		return m, cmd
	}

	return m, nil
}

func (m appModel) startGame(id string) (tea.Model, tea.Cmd) {
	entry, ok := games.Find(id)
	if !ok {
		return m, nil
	}

	m.currentGame = entry.Factory(m.settings.Values())
	m.state = StatePlaying
	return m, m.currentGame.Init()
}
