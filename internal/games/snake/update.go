package snake

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/barttob/terminal-arcade/internal/engine"
)

// maxQueuedTurns caps buffered input so holding a key doesn't pile up moves.
const maxQueuedTurns = 3

func (m *Model) Update(msg tea.Msg) (engine.Game, tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		// Pausing or ending the game lets the tick chain lapse.
		if msg.ID != m.tickID || m.paused || m.gameOver {
			return m, nil
		}
		m.step()
		return m, engine.Tick(m.tickID, m.tickRate)

	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			game := m.Reset()
			return game, game.Init()
		case "p":
			return m, m.togglePause()
		case "o":
			// Ticks aren't delivered while settings are open, so pause and
			// let the player resume with p on the way back.
			if !m.gameOver {
				m.paused = true
			}
			return m, func() tea.Msg { return engine.OpenSettingsMsg{GameID: ID} }
		}
		if m.paused || m.gameOver {
			return m, nil
		}

		if direction, ok := engine.DirectionFromKey(msg.String()); ok {
			m.queueTurn(direction)
		}
	}

	return m, nil
}

// togglePause pauses or resumes the game. Resuming starts a new tick chain
// under a new ID, so a tick still in flight from before the pause is dropped
// rather than doubling the speed.
func (m *Model) togglePause() tea.Cmd {
	if m.gameOver {
		return nil
	}

	m.paused = !m.paused
	if m.paused {
		return nil
	}
	m.tickID = engine.NewTickID()
	return m.Init()
}

// queueTurn buffers a turn, ignoring ones that repeat or reverse the heading
// the snake will have when the turn is applied.
func (m *Model) queueTurn(direction engine.Direction) {
	if len(m.turns) >= maxQueuedTurns {
		return
	}

	heading := m.direction
	if n := len(m.turns); n > 0 {
		heading = m.turns[n-1]
	}
	if direction == heading || direction == heading.Opposite() {
		return
	}
	m.turns = append(m.turns, direction)
}

func (m *Model) step() {
	if len(m.turns) > 0 {
		m.direction = m.turns[0]
		m.turns = m.turns[1:]
	}

	head := m.snake[0]
	next := engine.Point{X: head.X + m.direction.X, Y: head.Y + m.direction.Y}
	if !m.walls {
		next.X = (next.X + m.width) % m.width
		next.Y = (next.Y + m.height) % m.height
	}
	eating := engine.IsCollision(next, m.food)

	// The tail moves out of the way this tick unless the snake is growing.
	body := m.snake
	if !eating {
		body = body[:len(body)-1]
	}
	if !m.inBounds(next) || engine.ContainsPoint(body, next) {
		m.gameOver = true
		return
	}

	m.snake = append([]engine.Point{next}, body...)
	if !eating {
		return
	}

	m.score += FoodScore
	if !m.spawnFood() {
		m.won = true
		m.gameOver = true
	}
}
