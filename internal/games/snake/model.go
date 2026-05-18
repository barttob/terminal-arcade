package snake

import (
	"math/rand"

	"github.com/barttob/terminal-arcade/internal/engine"
)

type Model struct {
	width  int
	height int

	snake []engine.Point

	direction engine.Direction

	food engine.Point

	score    int
	gameOver bool
}

func NewModel(width, height int) *Model {
	m := Model{
		width:  width,
		height: height,
		snake: []engine.Point{
			{X: width / 2, Y: height / 2},
			{X: width/2 - 1, Y: height / 2},
			{X: width/2 - 2, Y: height / 2},
		},
		direction: engine.Right,
		food:      engine.Point{},
		score:     0,
		gameOver:  false,
	}
	m.food = m.spawnFood()
	return &m
}

func (m *Model) spawnFood() engine.Point {
	for {
		p := engine.Point{X: rand.Intn(m.width), Y: rand.Intn(m.height)}
		if !engine.ContainsPoint(m.snake, p) {
			return p
		}
	}
}
