package games

import (
	"github.com/barttob/terminal-arcade/internal/engine"
	"github.com/barttob/terminal-arcade/internal/games/minesweeper"
	"github.com/barttob/terminal-arcade/internal/games/snake"
	"github.com/barttob/terminal-arcade/internal/settings"
)

type Entry struct {
	ID       string
	Name     string
	Factory  func(settings.Settings) engine.Game
	Settings []settings.Row
}

// Registry is the single source of truth for which games exist. Each entry
// gets a main menu item and its own tab on the settings screen.
var Registry = []Entry{
	{
		ID:      snake.ID,
		Name:    "Snake",
		Factory: func(settings.Settings) engine.Game { return snake.New() },
	},
	{
		ID:       minesweeper.ID,
		Name:     "Minesweeper",
		Factory:  func(s settings.Settings) engine.Game { return minesweeper.New(s.Minesweeper) },
		Settings: settings.MinesweeperRows,
	},
	// {
	// 	ID:   tetris.ID,
	// 	Name: "Tetris",
	// 	Factory: func(s settings.Settings) engine.Game {
	// 		// start Tetris game
	// 	},
	// },
}

func Find(id string) (Entry, bool) {
	for _, entry := range Registry {
		if entry.ID == id {
			return entry, true
		}
	}
	return Entry{}, false
}

func SettingsSections() []settings.Section {
	sections := make([]settings.Section, len(Registry))
	for i, entry := range Registry {
		sections[i] = settings.Section{GameID: entry.ID, Title: entry.Name, Rows: entry.Settings}
	}
	return sections
}
