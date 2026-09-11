package menu

import (
	"github.com/barttob/terminal-arcade/internal/engine"
	"github.com/barttob/terminal-arcade/internal/games/snake"
	"github.com/barttob/terminal-arcade/internal/games/minesweeper"
	tea "github.com/charmbracelet/bubbletea"
)

type gameEntry struct {
	Name    string
	Factory func() engine.Game
}

type Model struct {
	choices []gameEntry
	cursor  int

	terminalWidth  int
	terminalHeight int
}

var gameRegistry = []gameEntry{
	{
		Name:    "Snake",
		Factory: snake.New,
	},
	{
		Name:    "Minesweeper",
		Factory: minesweeper.New,
	},
	// {
	// 	Name: "Tetris",
	// 	Factory: func() engine.Game {
	// 		// start Tetris game
	// 	},
	// },
}

func MenuModel() Model {
	return Model{
		choices: gameRegistry,
		cursor:  0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
