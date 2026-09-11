package snake

import (
	"math/rand"
	"time"

	"github.com/barttob/terminal-arcade/internal/engine"
)

type Model struct {
	width    int
	height   int
	tickRate time.Duration
	walls    bool

	// snake runs from head to tail.
	snake []engine.Point

	// direction is the heading of the last move; turns holds key presses not
	// yet applied, one per tick, so quick successive turns aren't lost.
	direction engine.Direction
	turns     []engine.Direction

	food engine.Point

	score    int
	paused   bool
	gameOver bool
	won      bool

	// tickID is replaced on unpause so ticks from before the pause are dropped.
	tickID int
}

func NewModel(width, height int, tickRate time.Duration, walls bool) *Model {
	m := Model{
		width:    width,
		height:   height,
		tickRate: tickRate,
		walls:    walls,
		snake: []engine.Point{
			{X: width / 2, Y: height / 2},
			{X: width/2 - 1, Y: height / 2},
			{X: width/2 - 2, Y: height / 2},
		},
		direction: engine.Right,
		tickID:    engine.NewTickID(),
	}
	m.spawnFood()
	return &m
}

func (m *Model) inBounds(p engine.Point) bool {
	return p.X >= 0 && p.X < m.width && p.Y >= 0 && p.Y < m.height
}

// spawnFood places food on a random free cell, returning false if the snake
// fills the board.
func (m *Model) spawnFood() bool {
	var free []engine.Point
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			p := engine.Point{X: x, Y: y}
			if !engine.ContainsPoint(m.snake, p) {
				free = append(free, p)
			}
		}
	}
	if len(free) == 0 {
		return false
	}

	m.food = free[rand.Intn(len(free))]
	return true
}
