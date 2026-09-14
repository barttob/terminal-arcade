package flappy

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/barttob/terminal-arcade/internal/styles"
)

func (m *Model) View() string {
	cells := make([][]string, m.height)
	for y := range cells {
		cells[y] = make([]string, m.width)
		for x := range cells[y] {
			cells[y][x] = " "
		}
	}

	pipeCell := styles.FlappyPipe.Render("█")
	for _, p := range m.pipes {
		left := int(math.Floor(p.x))
		for x := max(0, left); x < min(m.width, left+PipeWidth); x++ {
			for y := range cells {
				if p.occupies(x, y, m.gap) {
					cells[y][x] = pipeCell
				}
			}
		}
	}

	bird := styles.FlappyBird
	if m.gameOver {
		bird = styles.FlappyDead
	}
	cells[m.birdRow()][BirdX] = bird.Render("●")

	rows := make([]string, m.height)
	for y, row := range cells {
		rows[y] = strings.Join(row, "")
	}

	var footer string
	switch {
	case m.gameOver:
		footer = styles.Lose.Render("Game over!") + styles.Help.Render("  r play again · o settings · esc menu")
	case m.paused:
		footer = styles.Paused.Render("Paused") + styles.Help.Render("  p resume · o settings · r restart · esc menu")
	case !m.started:
		footer = styles.Help.Render("space/↑/W flap to start · o settings · esc menu · q quit")
	default:
		footer = styles.Help.Render("space/↑/W flap · p pause · o settings · r restart · esc menu · q quit")
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		styles.GameHeader.Render("FLAPPY BIRD"),
		fmt.Sprintf("Score: %d   Best: %d", m.score, m.best),
		styles.Board.Render(strings.Join(rows, "\n")),
		footer,
	)
}
