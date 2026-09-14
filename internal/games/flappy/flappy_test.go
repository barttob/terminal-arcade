package flappy

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/barttob/terminal-arcade/internal/engine"
)

// newTestModel builds a started game with a single pipe.
func newTestModel(birdY, velocity float64, p pipe) *Model {
	m := NewModel(DefaultWidth, DefaultHeight, Speeds[1].PipeSpeed, Gaps[1].Rows)
	m.started = true
	m.birdY = birdY
	m.velocity = velocity
	m.pipes = []pipe{p}
	return m
}

// farPipe is off to the right, out of the bird's way.
var farPipe = pipe{x: DefaultWidth, gapTop: 5}

func key(s string) tea.KeyMsg {
	if s == " " {
		return tea.KeyMsg{Type: tea.KeySpace, Runes: []rune(s)}
	}
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(s)}
}

func TestFirstFlapStartsTicks(t *testing.T) {
	m := NewModel(DefaultWidth, DefaultHeight, Speeds[1].PipeSpeed, Gaps[1].Rows)
	y := m.birdY

	if _, cmd := m.Update(engine.TickMsg{ID: m.tickID}); cmd != nil || m.birdY != y {
		t.Fatal("a tick moved the bird before the first flap")
	}

	_, cmd := m.Update(key(" "))
	if !m.started || cmd == nil {
		t.Fatal("the first flap didn't start the game")
	}
	if m.velocity != FlapVelocity {
		t.Fatalf("velocity = %v, want %v", m.velocity, FlapVelocity)
	}

	if _, cmd := m.Update(key("w")); cmd != nil {
		t.Fatal("a later flap started a second tick chain")
	}
}

func TestGravityPullsBirdDown(t *testing.T) {
	m := newTestModel(5, 0, farPipe)
	m.step()

	if m.velocity <= 0 || m.birdY <= 5 {
		t.Fatalf("bird at %v with velocity %v, want it falling", m.birdY, m.velocity)
	}
}

func TestGroundEndsGame(t *testing.T) {
	m := newTestModel(DefaultHeight-0.1, MaxFallSpeed, farPipe)
	m.step()

	if !m.gameOver {
		t.Fatal("expected a loss after hitting the ground")
	}
	if row := m.birdRow(); row != DefaultHeight-1 {
		t.Fatalf("bird row = %d, want it resting on the bottom row", row)
	}
}

func TestCeilingStopsBird(t *testing.T) {
	m := newTestModel(0.1, FlapVelocity, farPipe)
	m.step()

	if m.gameOver {
		t.Fatal("hitting the ceiling ended the game")
	}
	if m.birdY != 0 || m.velocity != 0 {
		t.Fatalf("bird at %v with velocity %v, want stopped at the ceiling", m.birdY, m.velocity)
	}
}

func TestPipeEndsGame(t *testing.T) {
	m := newTestModel(2, 0, pipe{x: BirdX, gapTop: 10})
	m.step()

	if !m.gameOver {
		t.Fatal("expected a loss after flying into a pipe")
	}
}

func TestPassingPipeScores(t *testing.T) {
	gapTop := 8
	m := newTestModel(float64(gapTop)+2.5, 0, pipe{x: BirdX - PipeWidth + 0.1, gapTop: gapTop})
	m.step()

	if m.gameOver {
		t.Fatal("flying through the gap ended the game")
	}
	if m.score != 1 || m.best != 1 {
		t.Fatalf("score = %d, best = %d, want 1 and 1", m.score, m.best)
	}

	m.step()
	if m.score != 1 {
		t.Fatalf("score = %d after another step, want the pipe counted once", m.score)
	}
}

func TestPipesSpawnAndDrop(t *testing.T) {
	m := newTestModel(5, 0, pipe{x: -PipeWidth, gapTop: 5})
	m.pipes = append(m.pipes, pipe{x: DefaultWidth - PipeSpacing, gapTop: 5})
	m.spawnPipes()

	if len(m.pipes) != 2 {
		t.Fatalf("pipes = %d, want the old one dropped and a new one added", len(m.pipes))
	}
	if got, want := m.pipes[1].x, float64(DefaultWidth); got != want {
		t.Fatalf("new pipe at x = %v, want %v", got, want)
	}
}

func TestGapsStayReachable(t *testing.T) {
	m := newTestModel(5, 0, farPipe)
	for _, prev := range []int{-1, 1, 6, m.height - m.gap - 1} {
		for i := 0; i < 200; i++ {
			top := m.randomGapTop(prev)
			if top < 1 || top+m.gap > m.height-1 {
				t.Fatalf("gap top %d leaves no pipe at an edge", top)
			}
			if prev >= 0 && (top < prev-MaxGapShift || top > prev+MaxGapShift) {
				t.Fatalf("gap top %d is more than %d rows from %d", top, MaxGapShift, prev)
			}
		}
	}
}

func TestUnpauseDropsTicksFromBeforePause(t *testing.T) {
	m := newTestModel(5, 0, farPipe)
	before := m.tickID

	m.Update(key("p"))
	if _, cmd := m.Update(engine.TickMsg{ID: before}); cmd != nil {
		t.Fatal("a tick while paused re-armed the chain")
	}
	if _, cmd := m.Update(key(" ")); cmd != nil || m.velocity != 0 {
		t.Fatal("flapping while paused moved the bird")
	}

	if _, cmd := m.Update(key("p")); cmd == nil {
		t.Fatal("unpausing didn't start a new tick")
	}
	if _, cmd := m.Update(engine.TickMsg{ID: before}); cmd != nil {
		t.Fatal("a tick from before the pause was accepted after resuming")
	}
}

func TestOpenSettingsPauses(t *testing.T) {
	m := newTestModel(5, 0, farPipe)
	_, cmd := m.Update(key("o"))

	if !m.paused {
		t.Fatal("opening settings didn't pause the game")
	}
	if msg, ok := cmd().(engine.OpenSettingsMsg); !ok || msg.GameID != ID {
		t.Fatalf("o produced %#v, want OpenSettingsMsg for flappy", cmd())
	}
}

func TestResetKeepsBest(t *testing.T) {
	m := newTestModel(5, 0, farPipe)
	m.score, m.best, m.gameOver = 3, 7, true

	game, _ := m.Update(key("r"))
	next := game.(*Model)
	if next.score != 0 || next.best != 7 || next.gameOver || next.started {
		t.Fatalf("after restart: score %d, best %d, over %v, started %v", next.score, next.best, next.gameOver, next.started)
	}
}

func TestNewClampsConfig(t *testing.T) {
	m := New(Config{Speed: 99, Gap: -1}).(*Model)

	if want := Speeds[len(Speeds)-1].PipeSpeed; m.pipeSpeed != want {
		t.Fatalf("pipe speed = %v, want %v", m.pipeSpeed, want)
	}
	if want := Gaps[0].Rows; m.gap != want {
		t.Fatalf("gap = %d, want %d", m.gap, want)
	}
}
