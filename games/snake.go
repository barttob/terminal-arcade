package games

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type point struct {
	x, y int
}

type SnakeModel struct {
	width  int
	height int
	speed  int

	snake []point
	dir   point
	food  point

	score    int
	gameOver bool

	ExitToMenu bool
}

type tickMsg time.Time

var (
	boardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(0, 1)

	snakeHeadStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("120")).
			Bold(true)

	snakeBodyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42"))

	foodStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("203")).
			Bold(true)

	gameTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true)

	gameHelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	gameOverStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("196")).
			Border(lipgloss.DoubleBorder()).
			Padding(1, 4)
)

func NewSnakeModel(width, height, speed int) SnakeModel {
	rand.Seed(time.Now().UnixNano())

	m := SnakeModel{
		width:  width,
		height: height,
		speed:  speed,
		snake: []point{
			{width / 2, height / 2},
			{width/2 - 1, height / 2},
			{width/2 - 2, height / 2},
		},
		dir: point{1, 0},
	}

	m.food = randomFood(m)

	return m
}

func (m SnakeModel) Init() tea.Cmd {
	return snakeTick(m.speed)
}

func snakeTick(speed int) tea.Cmd {
	return tea.Tick(time.Duration(speed)*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m SnakeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.gameOver {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "r":
				newGame := NewSnakeModel(m.width, m.height, m.speed)
				return newGame, newGame.Init()

			case "esc", "q":
				m.ExitToMenu = true
				return m, nil
			}
		}

		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.ExitToMenu = true
			return m, nil

		case "up", "w":
			if m.dir.y != 1 {
				m.dir = point{0, -1}
			}

		case "down", "s":
			if m.dir.y != -1 {
				m.dir = point{0, 1}
			}

		case "left", "a":
			if m.dir.x != 1 {
				m.dir = point{-1, 0}
			}

		case "right", "d":
			if m.dir.x != -1 {
				m.dir = point{1, 0}
			}
		}

	case tickMsg:
		m = m.move()

		if !m.gameOver {
			return m, snakeTick(m.speed)
		}
	}

	return m, nil
}

func (m SnakeModel) move() SnakeModel {
	head := m.snake[0]

	newHead := point{
		x: head.x + m.dir.x,
		y: head.y + m.dir.y,
	}

	if newHead.x < 0 ||
		newHead.x >= m.width ||
		newHead.y < 0 ||
		newHead.y >= m.height {
		m.gameOver = true
		return m
	}

	for _, part := range m.snake {
		if part == newHead {
			m.gameOver = true
			return m
		}
	}

	m.snake = append([]point{newHead}, m.snake...)

	if newHead == m.food {
		m.score++
		m.food = randomFood(m)
	} else {
		m.snake = m.snake[:len(m.snake)-1]
	}

	return m
}

func randomFood(m SnakeModel) point {
	for {
		p := point{
			x: rand.Intn(m.width),
			y: rand.Intn(m.height),
		}

		if !containsPoint(m.snake, p) {
			return p
		}
	}
}

func (m SnakeModel) View() string {
	if m.gameOver {
		text := fmt.Sprintf(
			"GAME OVER\n\nScore: %d\n\nr - restart\nq/esc - menu",
			m.score,
		)

		return center(gameOverStyle.Render(text))
	}

	header := gameTitleStyle.Render(fmt.Sprintf("SNAKE  Score: %d", m.score))
	board := boardStyle.Render(renderBoard(m))
	help := gameHelpStyle.Render("WASD/arrows move • esc menu")

	return center(lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		board,
		help,
	))
}

func renderBoard(m SnakeModel) string {
	var b strings.Builder

	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			p := point{x, y}

			switch {
			case p == m.snake[0]:
				b.WriteString(snakeHeadStyle.Render("██"))

			case p == m.food:
				b.WriteString(foodStyle.Render("▓▓"))

			case containsPoint(m.snake[1:], p):
				b.WriteString(snakeBodyStyle.Render("██"))

			default:
				b.WriteString("  ")
			}
		}

		if y < m.height-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

func containsPoint(points []point, p point) bool {
	for _, item := range points {
		if item == p {
			return true
		}
	}

	return false
}

func center(s string) string {
	return lipgloss.NewStyle().
		Width(80).
		Height(24).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Render(s)
}
