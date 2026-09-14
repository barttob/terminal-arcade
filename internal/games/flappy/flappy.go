package flappy

import (
	"time"

	"github.com/barttob/terminal-arcade/internal/engine"
	tea "github.com/charmbracelet/bubbletea"
)

const ID = "flappy"

// Board size; the view adds 5 rows and 2 columns, so it fits an 80×24
// terminal.
const (
	DefaultWidth  = 60
	DefaultHeight = 18
)

// DefaultTickRate is the physics step, about 30 frames a second.
const DefaultTickRate = 33 * time.Millisecond

// Bird physics, in rows per second (squared for gravity). A flap sets the
// velocity rather than adding to it, as in the original.
const (
	Gravity      = 60.0
	FlapVelocity = -15.0
	MaxFallSpeed = 20.0
)

// Pipe layout, in cells.
const (
	BirdX       = 12
	PipeWidth   = 4
	PipeSpacing = 24
	// MaxGapShift limits how far a gap moves from the previous pipe's, so the
	// next gap is always reachable.
	MaxGapShift = 5
)

// Speed sets how fast pipes scroll. It stays under one column per tick so a
// pipe can't skip past the bird.
type Speed struct {
	Name      string
	PipeSpeed float64 // columns per second
}

var Speeds = []Speed{
	{Name: "Slow", PipeSpeed: 12},
	{Name: "Normal", PipeSpeed: 16},
	{Name: "Fast", PipeSpeed: 21},
}

type Gap struct {
	Name string
	Rows int
}

var Gaps = []Gap{
	{Name: "Wide", Rows: 7},
	{Name: "Normal", Rows: 5},
	{Name: "Narrow", Rows: 4},
}

type Config struct {
	// Speed indexes Speeds.
	Speed int
	// Gap indexes Gaps.
	Gap int
}

var DefaultConfig = Config{Speed: 1, Gap: 1}

// Clamped returns c with the speed and gap forced into valid ranges.
func (c Config) Clamped() Config {
	c.Speed = max(0, min(c.Speed, len(Speeds)-1))
	c.Gap = max(0, min(c.Gap, len(Gaps)-1))
	return c
}

func (m *Model) Name() string {
	return "Flappy Bird"
}

func (m *Model) Description() string {
	return "Flap through the gaps between pipes without touching them or the ground."
}

// Init arms nothing: the bird hovers until the first flap starts the ticks.
func (m *Model) Init() tea.Cmd {
	return nil
}

// Reset starts a new round, carrying over the best score of the session.
func (m *Model) Reset() engine.Game {
	game := NewModel(m.width, m.height, m.pipeSpeed, m.gap)
	game.best = m.best
	return game
}

func (m *Model) IsGameOver() bool {
	return m.gameOver
}

// Score is the number of pipes passed.
func (m *Model) Score() int {
	return m.score
}

func New(config Config) engine.Game {
	config = config.Clamped()
	return NewModel(DefaultWidth, DefaultHeight, Speeds[config.Speed].PipeSpeed, Gaps[config.Gap].Rows)
}
