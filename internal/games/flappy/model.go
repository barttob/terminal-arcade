package flappy

import (
	"math"
	"math/rand"

	"github.com/barttob/terminal-arcade/internal/engine"
)

// pipe is a top and bottom pipe sharing a column range, with a gap between.
type pipe struct {
	// x is the left edge in columns; it's fractional so pipes can scroll
	// slower than one column per tick.
	x      float64
	gapTop int
	passed bool
}

type Model struct {
	width     int
	height    int
	pipeSpeed float64
	gap       int

	// birdY is measured in rows from the top; the bird is drawn and collides
	// at birdRow.
	birdY    float64
	velocity float64

	// pipes run left to right.
	pipes []pipe

	score int
	best  int

	// started stays false until the first flap; the bird hovers until then.
	started  bool
	paused   bool
	gameOver bool

	// tickID is replaced on unpause so ticks from before the pause are dropped.
	tickID int
}

func NewModel(width, height int, pipeSpeed float64, gap int) *Model {
	m := Model{
		width:     width,
		height:    height,
		pipeSpeed: pipeSpeed,
		gap:       gap,
		birdY:     float64(height) / 2,
		tickID:    engine.NewTickID(),
	}
	m.pipes = []pipe{{x: float64(width), gapTop: m.randomGapTop(-1)}}
	return &m
}

func (m *Model) birdRow() int {
	return int(math.Floor(m.birdY))
}

// randomGapTop picks the top row of a new gap, leaving at least one row of
// pipe above and below it and, if prev is a gap top, staying within
// MaxGapShift of it.
func (m *Model) randomGapTop(prev int) int {
	lo, hi := 1, m.height-m.gap-1
	if prev >= 0 {
		lo = max(lo, prev-MaxGapShift)
		hi = min(hi, prev+MaxGapShift)
	}
	return lo + rand.Intn(hi-lo+1)
}

// occupies reports whether the pipe fills the cell at column x, row y.
func (p pipe) occupies(x, y, gap int) bool {
	left := int(math.Floor(p.x))
	if x < left || x >= left+PipeWidth {
		return false
	}
	return y < p.gapTop || y >= p.gapTop+gap
}
