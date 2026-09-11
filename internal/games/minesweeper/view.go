package minesweeper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/barttob/terminal-arcade/internal/engine"
	"github.com/barttob/terminal-arcade/internal/styles"
)

func (m *Model) View() string {
	var board strings.Builder
	for y := 0; y < m.height; y++ {
		for x := 0; x <= m.width; x++ {
			board.WriteString(m.separator(x, y))
			if x < m.width {
				board.WriteString(m.renderCell(engine.Point{X: x, Y: y}))
			}
		}
		if y < m.height-1 {
			board.WriteByte('\n')
		}
	}

	var footer string
	switch {
	case m.won:
		footer = styles.Win.Render("You cleared the field!") + styles.Help.Render("  r play again · esc menu")
	case m.gameOver:
		footer = styles.Lose.Render("Boom!") + styles.Help.Render("  r try again · esc menu")
	default:
		footer = styles.Help.Render("arrows/WASD move · space reveal · f flag · r restart · esc menu · q quit")
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		styles.GameHeader.Render("MINESWEEPER"),
		fmt.Sprintf("Mines left: %d", m.mines-m.flags),
		styles.Board.Render(board.String()),
		footer,
	)
}

// separator returns the gap to the left of column x, bracketing the cursor cell.
func (m *Model) separator(x, y int) string {
	if m.gameOver || y != m.cursor.Y {
		return " "
	}
	switch x {
	case m.cursor.X:
		return styles.MineCursor.Render("[")
	case m.cursor.X + 1:
		return styles.MineCursor.Render("]")
	}
	return " "
}

func (m *Model) renderCell(p engine.Point) string {
	c := m.at(p)
	switch {
	case c.flagged && m.gameOver && !c.mine:
		return styles.MineWrongFlag.Render("x")
	case c.flagged, m.won && c.mine:
		return styles.MineFlag.Render("⚑")
	case m.gameOver && c.mine:
		if p == m.exploded {
			return styles.MineExploded.Render("*")
		}
		return styles.Mine.Render("*")
	case !c.revealed:
		return styles.MineHidden.Render("■")
	case c.adjacent == 0:
		return styles.MineEmpty.Render("·")
	default:
		return styles.MineNumbers[c.adjacent].Render(strconv.Itoa(c.adjacent))
	}
}
