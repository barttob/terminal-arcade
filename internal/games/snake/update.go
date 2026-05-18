package snake

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/barttob/terminal-arcade/internal/engine"
)

func (m *Model) Update(msg tea.Msg) (engine.Game, tea.Cmd) {
	direction, ok := engine.DirectionFromKey(msg.(tea.KeyMsg).String())
	if ok {
		if direction != m.direction.Opposite() {
			m.direction = direction
		}
	}

	return m, nil
}

func (m *Model) step() {

}
