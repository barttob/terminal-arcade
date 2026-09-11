package snake

import (
	"fmt"
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

	if !m.won {
		cells[m.food.Y][m.food.X] = styles.SnakeFood.Render("◆")
	}
	body := styles.SnakeBody.Render("●")
	for _, p := range m.snake[1:] {
		cells[p.Y][p.X] = body
	}
	head := styles.SnakeHead
	if m.gameOver && !m.won {
		head = styles.SnakeDead
	}
	cells[m.snake[0].Y][m.snake[0].X] = head.Render("●")

	// Cells are spaced out so the board's aspect ratio roughly matches the
	// terminal's tall character cells.
	var board strings.Builder
	for y, row := range cells {
		for _, cell := range row {
			board.WriteByte(' ')
			board.WriteString(cell)
		}
		board.WriteByte(' ')
		if y < m.height-1 {
			board.WriteByte('\n')
		}
	}

	var footer string
	switch {
	case m.won:
		footer = styles.Win.Render("You filled the board!") + styles.Help.Render("  r play again · o settings · esc menu")
	case m.gameOver:
		footer = styles.Lose.Render("Game over!") + styles.Help.Render("  r play again · o settings · esc menu")
	case m.paused:
		footer = styles.Paused.Render("Paused") + styles.Help.Render("  p resume · o settings · r restart · esc menu")
	default:
		footer = styles.Help.Render("arrows/WASD move · p pause · o settings · r restart · esc menu · q quit")
	}

	frame := styles.Board
	if !m.walls {
		frame = styles.SnakeOpenBoard
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		styles.GameHeader.Render("SNAKE"),
		fmt.Sprintf("Score: %d", m.score),
		frame.Render(board.String()),
		footer,
	)
}
