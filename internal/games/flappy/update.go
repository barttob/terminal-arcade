package flappy

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/barttob/terminal-arcade/internal/engine"
)

func (m *Model) Update(msg tea.Msg) (engine.Game, tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		// Pausing or ending the game lets the tick chain lapse.
		if msg.ID != m.tickID || !m.started || m.paused || m.gameOver {
			return m, nil
		}
		m.step()
		if m.gameOver {
			return m, nil
		}
		return m, engine.Tick(m.tickID, DefaultTickRate)

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
			if m.started && !m.gameOver {
				m.paused = true
			}
			return m, func() tea.Msg { return engine.OpenSettingsMsg{GameID: ID} }
		case " ", "up", "w":
			return m, m.flap()
		}
	}

	return m, nil
}

// flap launches the bird upward. The first flap starts the tick chain.
func (m *Model) flap() tea.Cmd {
	if m.paused || m.gameOver {
		return nil
	}

	m.velocity = FlapVelocity
	if m.started {
		return nil
	}
	m.started = true
	return engine.Tick(m.tickID, DefaultTickRate)
}

// togglePause pauses or resumes the game. Resuming starts a new tick chain
// under a new ID, so a tick still in flight from before the pause is dropped
// rather than doubling the speed.
func (m *Model) togglePause() tea.Cmd {
	if !m.started || m.gameOver {
		return nil
	}

	m.paused = !m.paused
	if m.paused {
		return nil
	}
	m.tickID = engine.NewTickID()
	return engine.Tick(m.tickID, DefaultTickRate)
}

func (m *Model) step() {
	dt := DefaultTickRate.Seconds()

	m.velocity = min(m.velocity+Gravity*dt, MaxFallSpeed)
	m.birdY += m.velocity * dt
	if m.birdY < 0 {
		// The ceiling stops the bird without ending the game.
		m.birdY = 0
		m.velocity = 0
	}
	if m.birdY >= float64(m.height) {
		m.birdY = float64(m.height - 1)
		m.gameOver = true
		return
	}

	for i := range m.pipes {
		m.pipes[i].x -= m.pipeSpeed * dt
	}
	m.spawnPipes()

	row := m.birdRow()
	for i := range m.pipes {
		p := &m.pipes[i]
		if p.occupies(BirdX, row, m.gap) {
			m.gameOver = true
			return
		}
		if !p.passed && p.x+PipeWidth <= BirdX {
			p.passed = true
			m.score++
			m.best = max(m.best, m.score)
		}
	}
}

// spawnPipes adds a pipe once the last one has scrolled PipeSpacing columns
// in, and drops pipes that have left the board.
func (m *Model) spawnPipes() {
	if last := m.pipes[len(m.pipes)-1]; last.x <= float64(m.width-PipeSpacing) {
		m.pipes = append(m.pipes, pipe{x: last.x + PipeSpacing, gapTop: m.randomGapTop(last.gapTop)})
	}
	for len(m.pipes) > 1 && m.pipes[0].x+PipeWidth <= 0 {
		m.pipes = m.pipes[1:]
	}
}
