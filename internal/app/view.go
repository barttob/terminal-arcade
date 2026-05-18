package app

func (m appModel) View() string {
	switch m.state {
	case StateMainMenu:
		return center(m.menu.View(), m.terminalWidth, m.terminalHeight)
	case StatePlaying:
		return center(m.currentGame.View(), m.terminalWidth, m.terminalHeight)
	case StatePaused:
		return center("Paused", m.terminalWidth, m.terminalHeight)
	case StateGameOver:
		return center("Game Over", m.terminalWidth, m.terminalHeight)
	case StateSettings:
		return center("Settings", m.terminalWidth, m.terminalHeight)
	case StateHelp:
		return center("Help", m.terminalWidth, m.terminalHeight)
	case StateLobby:
		return center("Lobby", m.terminalWidth, m.terminalHeight)
	case StateRoom:
		return center("Room", m.terminalWidth, m.terminalHeight)
	default:
		return ""
	}
}
