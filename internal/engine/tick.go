package engine

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg time.Time

func Tick(d time.Duration) tea.Cmd {
	return tea.Tick(d*time.Millisecond, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
