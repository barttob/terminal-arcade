package settings

import tea "github.com/charmbracelet/bubbletea"

// Section is one game's tab on the settings screen.
type Section struct {
	GameID string
	Title  string
	Rows   []Row
}

type Model struct {
	values   Settings
	sections []Section
	section  int
	// cursor indexes the current section's rows; len(Rows) is the Play button.
	cursor int
}

func New(values Settings, sections []Section) Model {
	return Model{values: values, sections: sections}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Values() Settings {
	return m.values
}

// Focus switches to gameID's tab. An unknown or empty ID keeps the current tab.
func (m Model) Focus(gameID string) Model {
	for i, s := range m.sections {
		if s.GameID == gameID {
			m.section = i
			m.cursor = 0
		}
	}
	return m
}

func (m Model) current() Section {
	return m.sections[m.section]
}
