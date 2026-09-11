package engine

import (
	"sync/atomic"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// TickMsg is delivered by Tick. ID tells a game whether it scheduled the tick,
// so it can drop ticks still in flight from a game that was restarted or left.
type TickMsg struct {
	ID   int
	Time time.Time
}

var lastTickID atomic.Int64

// NewTickID returns an ID no other game in this process has used.
func NewTickID() int {
	return int(lastTickID.Add(1))
}

func Tick(id int, d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return TickMsg{ID: id, Time: t}
	})
}
