package minesweeper

import (
	// "math/rand"

	// "github.com/barttob/terminal-arcade/internal/engine"
)

type Model struct {
	width  int
	height int

	score    int
	gameOver bool
}

func NewModel(width, height int) *Model {
	m := Model{
		width:  width,
		height: height,

		score:     0,
		gameOver:  false,
	}
	return &m
}
