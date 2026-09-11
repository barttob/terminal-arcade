package minesweeper

import (
	"testing"

	"github.com/barttob/terminal-arcade/internal/engine"
)

// newTestModel builds a board with mines at fixed positions.
func newTestModel(width, height int, mines ...engine.Point) *Model {
	m := NewModel(width, height, len(mines))
	for _, p := range mines {
		m.setMine(p)
	}
	m.minesPlaced = true
	return m
}

func TestFirstRevealIsSafe(t *testing.T) {
	for i := 0; i < 100; i++ {
		m := NewModel(9, 9, 60)
		start := engine.Point{X: 4, Y: 4}
		m.reveal(start)

		if m.gameOver && !m.won {
			t.Fatal("first reveal hit a mine")
		}
		for _, p := range append(m.neighbors(start), start) {
			if m.at(p).mine {
				t.Fatalf("mine placed next to first reveal at %v", p)
			}
		}
	}
}

func TestFloodFillStopsAtNumbers(t *testing.T) {
	m := newTestModel(5, 5, engine.Point{X: 4, Y: 4})
	m.reveal(engine.Point{X: 0, Y: 0})

	if !m.won {
		t.Fatalf("expected win after flood fill, revealed %d of 24", m.revealed)
	}
	if m.at(engine.Point{X: 4, Y: 4}).revealed {
		t.Fatal("flood fill revealed the mine")
	}
}

func TestRevealMineEndsGame(t *testing.T) {
	mine := engine.Point{X: 1, Y: 1}
	m := newTestModel(3, 3, mine)
	m.reveal(mine)

	if !m.gameOver || m.won {
		t.Fatal("expected a loss after revealing a mine")
	}
	if m.exploded != mine {
		t.Fatalf("exploded = %v, want %v", m.exploded, mine)
	}
}

func TestFlagBlocksReveal(t *testing.T) {
	m := newTestModel(3, 3, engine.Point{X: 2, Y: 2})
	p := engine.Point{X: 0, Y: 0}
	m.toggleFlag(p)
	m.reveal(p)

	if m.at(p).revealed {
		t.Fatal("flagged cell was revealed")
	}
	if m.flags != 1 {
		t.Fatalf("flags = %d, want 1", m.flags)
	}
}
