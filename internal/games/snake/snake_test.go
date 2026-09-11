package snake

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/barttob/terminal-arcade/internal/engine"
)

// newTestModel builds a walled board with a fixed snake (head first) and food.
func newTestModel(width, height int, direction engine.Direction, food engine.Point, body ...engine.Point) *Model {
	m := NewModel(width, height, time.Millisecond, true)
	m.snake = body
	m.direction = direction
	m.food = food
	return m
}

func TestEatingGrows(t *testing.T) {
	m := newTestModel(10, 10, engine.Right, engine.Point{X: 5, Y: 5},
		engine.Point{X: 4, Y: 5}, engine.Point{X: 3, Y: 5}, engine.Point{X: 2, Y: 5})
	m.step()

	if len(m.snake) != 4 {
		t.Fatalf("length = %d, want 4", len(m.snake))
	}
	if m.score != FoodScore {
		t.Fatalf("score = %d, want %d", m.score, FoodScore)
	}
	if engine.ContainsPoint(m.snake, m.food) {
		t.Fatalf("food respawned under the snake at %v", m.food)
	}
}

func TestWallEndsGame(t *testing.T) {
	m := newTestModel(5, 5, engine.Right, engine.Point{X: 0, Y: 0},
		engine.Point{X: 4, Y: 2}, engine.Point{X: 3, Y: 2}, engine.Point{X: 2, Y: 2})
	m.step()

	if !m.gameOver || m.won {
		t.Fatal("expected a loss after hitting the wall")
	}
	if !m.inBounds(m.snake[0]) {
		t.Fatalf("head moved off the board to %v", m.snake[0])
	}
}

func TestNoWallsWrapsAround(t *testing.T) {
	m := newTestModel(5, 5, engine.Right, engine.Point{X: 2, Y: 0},
		engine.Point{X: 4, Y: 2}, engine.Point{X: 3, Y: 2}, engine.Point{X: 2, Y: 2})
	m.walls = false
	m.step()

	if m.gameOver {
		t.Fatal("crossing the edge without walls ended the game")
	}
	if want := (engine.Point{X: 0, Y: 2}); m.snake[0] != want {
		t.Fatalf("head = %v, want %v", m.snake[0], want)
	}
}

func TestSelfCollisionEndsGame(t *testing.T) {
	// Heading up into the body of a hooked snake.
	m := newTestModel(10, 10, engine.Up, engine.Point{X: 0, Y: 0},
		engine.Point{X: 5, Y: 5}, engine.Point{X: 5, Y: 6}, engine.Point{X: 6, Y: 6},
		engine.Point{X: 6, Y: 5}, engine.Point{X: 6, Y: 4}, engine.Point{X: 5, Y: 4}, engine.Point{X: 4, Y: 4})
	m.step()

	if !m.gameOver {
		t.Fatal("expected a loss after running into the body")
	}
}

func TestChasingTailIsAllowed(t *testing.T) {
	// A 2×2 loop: the head moves into the cell the tail is leaving.
	m := newTestModel(5, 5, engine.Up, engine.Point{X: 4, Y: 4},
		engine.Point{X: 1, Y: 2}, engine.Point{X: 2, Y: 2}, engine.Point{X: 2, Y: 1}, engine.Point{X: 1, Y: 1})
	m.step()

	if m.gameOver {
		t.Fatal("moving into the vacating tail ended the game")
	}
}

func TestQueuedTurns(t *testing.T) {
	m := newTestModel(10, 10, engine.Right, engine.Point{X: 0, Y: 0},
		engine.Point{X: 5, Y: 5}, engine.Point{X: 4, Y: 5}, engine.Point{X: 3, Y: 5})

	m.queueTurn(engine.Left) // reverses the heading
	m.queueTurn(engine.Up)
	m.queueTurn(engine.Left) // valid once the up turn has been taken

	m.step()
	if m.direction != engine.Up {
		t.Fatalf("direction after first step = %v, want up", m.direction)
	}
	m.step()
	if m.direction != engine.Left {
		t.Fatalf("direction after second step = %v, want left", m.direction)
	}
	if m.gameOver {
		t.Fatal("queued turns reversed the snake into itself")
	}
}

func TestFillingBoardWins(t *testing.T) {
	// The snake fills all but the food cell of a 2×2 board.
	m := newTestModel(2, 2, engine.Down, engine.Point{X: 0, Y: 1},
		engine.Point{X: 0, Y: 0}, engine.Point{X: 1, Y: 0}, engine.Point{X: 1, Y: 1})
	m.step()

	if !m.won || !m.gameOver {
		t.Fatal("expected a win after filling the board")
	}
}

func key(s string) tea.KeyMsg {
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(s)}
}

func TestUnpauseDropsTicksFromBeforePause(t *testing.T) {
	m := newTestModel(10, 10, engine.Right, engine.Point{X: 0, Y: 0},
		engine.Point{X: 5, Y: 5}, engine.Point{X: 4, Y: 5}, engine.Point{X: 3, Y: 5})
	before := m.tickID

	m.Update(key("p"))
	if _, cmd := m.Update(engine.TickMsg{ID: before}); cmd != nil {
		t.Fatal("a tick while paused re-armed the chain")
	}

	// Resuming before the old tick lands must not leave two chains running.
	if _, cmd := m.Update(key("p")); cmd == nil {
		t.Fatal("unpausing didn't start a new tick")
	}
	if _, cmd := m.Update(engine.TickMsg{ID: before}); cmd != nil {
		t.Fatal("a tick from before the pause was accepted after resuming")
	}
}

func TestOpenSettingsPauses(t *testing.T) {
	m := NewModel(10, 10, time.Millisecond, true)
	_, cmd := m.Update(key("o"))

	if !m.paused {
		t.Fatal("opening settings didn't pause the game")
	}
	if msg, ok := cmd().(engine.OpenSettingsMsg); !ok || msg.GameID != ID {
		t.Fatalf("o produced %#v, want OpenSettingsMsg for snake", cmd())
	}
}

func TestNewClampsConfig(t *testing.T) {
	m := New(Config{Width: 1, Height: 100, Speed: 99}).(*Model)

	if m.width != MinSize || m.height != MaxHeight {
		t.Fatalf("board = %d×%d, want %d×%d", m.width, m.height, MinSize, MaxHeight)
	}
	if want := Speeds[len(Speeds)-1].TickRate; m.tickRate != want {
		t.Fatalf("tick rate = %v, want %v", m.tickRate, want)
	}
}

func TestStaleTickIgnored(t *testing.T) {
	m := NewModel(10, 10, time.Millisecond, true)
	head := m.snake[0]

	_, cmd := m.Update(engine.TickMsg{ID: m.tickID - 1})
	if cmd != nil || m.snake[0] != head {
		t.Fatal("a tick from another game moved the snake")
	}

	_, cmd = m.Update(engine.TickMsg{ID: m.tickID})
	if cmd == nil || m.snake[0] == head {
		t.Fatal("the game's own tick didn't move the snake and re-arm")
	}
}
