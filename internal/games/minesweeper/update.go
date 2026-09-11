package minesweeper

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/barttob/terminal-arcade/internal/engine"
)

func (m *Model) Update(msg tea.Msg) (engine.Game, tea.Cmd) {
	return m, nil
}

func (m *Model) step() {

}
