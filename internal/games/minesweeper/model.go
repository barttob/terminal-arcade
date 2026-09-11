package minesweeper

import (
	"math/rand"

	"github.com/barttob/terminal-arcade/internal/engine"
)

type cell struct {
	mine     bool
	revealed bool
	flagged  bool
	adjacent int
}

type Model struct {
	width  int
	height int
	mines  int

	board  [][]cell
	cursor engine.Point

	// Mines are placed on the first reveal so the opening move is always safe.
	minesPlaced bool
	flags       int
	revealed    int

	gameOver bool
	won      bool
	exploded engine.Point
}

func NewModel(width, height, mines int) *Model {
	board := make([][]cell, height)
	for y := range board {
		board[y] = make([]cell, width)
	}

	return &Model{
		width:  width,
		height: height,
		mines:  mines,
		board:  board,
		cursor: engine.Point{X: width / 2, Y: height / 2},
	}
}

func (m *Model) at(p engine.Point) *cell {
	return &m.board[p.Y][p.X]
}

func (m *Model) inBounds(p engine.Point) bool {
	return p.X >= 0 && p.X < m.width && p.Y >= 0 && p.Y < m.height
}

func (m *Model) neighbors(p engine.Point) []engine.Point {
	var result []engine.Point
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			n := engine.Point{X: p.X + dx, Y: p.Y + dy}
			if (dx != 0 || dy != 0) && m.inBounds(n) {
				result = append(result, n)
			}
		}
	}
	return result
}

// placeMines scatters mines randomly, keeping safe and its neighbors clear.
func (m *Model) placeMines(safe engine.Point) {
	safeZone := append(m.neighbors(safe), safe)

	var candidates []engine.Point
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			p := engine.Point{X: x, Y: y}
			if !engine.ContainsPoint(safeZone, p) {
				candidates = append(candidates, p)
			}
		}
	}

	rand.Shuffle(len(candidates), func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})

	m.mines = min(m.mines, len(candidates))
	for _, p := range candidates[:m.mines] {
		m.setMine(p)
	}
	m.minesPlaced = true
}

func (m *Model) setMine(p engine.Point) {
	m.at(p).mine = true
	for _, n := range m.neighbors(p) {
		m.at(n).adjacent++
	}
}
