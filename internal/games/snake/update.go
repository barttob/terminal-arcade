package snake

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/barttob/terminal-arcade/internal/engine"
)

func (m *Model) Update(msg tea.Msg) (engine.Game, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	direction, ok := engine.DirectionFromKey(key.String())
	if ok {
		if direction != m.direction.Opposite() {
			m.direction = direction
		}
	}

	return m, nil
}

func (m *Model) step() {

}
