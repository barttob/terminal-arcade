package minesweeper

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/barttob/terminal-arcade/internal/engine"
)

func (m *Model) Update(msg tea.Msg) (engine.Game, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch key.String() {
	case "r":
		return m.Reset(), nil
	case "o":
		return m, func() tea.Msg { return engine.OpenSettingsMsg{GameID: ID} }
	}
	if m.gameOver {
		return m, nil
	}

	if direction, ok := engine.DirectionFromKey(key.String()); ok {
		next := engine.Point{X: m.cursor.X + direction.X, Y: m.cursor.Y + direction.Y}
		if m.inBounds(next) {
			m.cursor = next
		}
		return m, nil
	}

	switch key.String() {
	case " ", "enter":
		m.reveal(m.cursor)
	case "f":
		m.toggleFlag(m.cursor)
	}

	return m, nil
}

func (m *Model) reveal(p engine.Point) {
	c := m.at(p)
	if c.revealed || c.flagged {
		return
	}

	if !m.minesPlaced {
		m.placeMines(p)
	}

	c.revealed = true
	if c.mine {
		m.exploded = p
		m.gameOver = true
		return
	}
	m.revealed++

	// Flood-fill outward through cells with no adjacent mines.
	queue := []engine.Point{p}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if m.at(current).adjacent != 0 {
			continue
		}
		for _, n := range m.neighbors(current) {
			nc := m.at(n)
			if nc.revealed || nc.flagged {
				continue
			}
			nc.revealed = true
			m.revealed++
			queue = append(queue, n)
		}
	}

	if m.revealed == m.width*m.height-m.mines {
		m.won = true
		m.gameOver = true
	}
}

func (m *Model) toggleFlag(p engine.Point) {
	c := m.at(p)
	if c.revealed {
		return
	}

	c.flagged = !c.flagged
	if c.flagged {
		m.flags++
	} else {
		m.flags--
	}
}
