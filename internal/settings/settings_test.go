package settings

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/barttob/terminal-arcade/internal/engine"
	"github.com/barttob/terminal-arcade/internal/games/minesweeper"
	"github.com/barttob/terminal-arcade/internal/games/snake"
)

func testModel() Model {
	return New(Default(), []Section{
		{GameID: snake.ID, Title: "Snake", Rows: SnakeRows},
		{GameID: minesweeper.ID, Title: "Minesweeper", Rows: MinesweeperRows},
	})
}

func TestFocusSelectsTab(t *testing.T) {
	m := testModel().Focus(minesweeper.ID)
	if m.current().GameID != minesweeper.ID {
		t.Fatalf("focused tab = %q, want %q", m.current().GameID, minesweeper.ID)
	}
	if m.Focus("").current().GameID != minesweeper.ID {
		t.Fatal("empty Focus should keep the current tab")
	}
}

func TestTabWrapsAndEnterStartsGame(t *testing.T) {
	var model tea.Model = testModel()
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyShiftTab})
	if got := model.(Model).current().GameID; got != minesweeper.ID {
		t.Fatalf("shift+tab from first tab = %q, want last tab", got)
	}

	_, cmd := model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("enter returned no command")
	}
	if msg, ok := cmd().(engine.StartGameMsg); !ok || msg.GameID != minesweeper.ID {
		t.Fatalf("enter produced %#v, want StartGameMsg for minesweeper", cmd())
	}
}

func TestCyclePresetWraps(t *testing.T) {
	first := minesweeper.Presets[0].Config
	last := minesweeper.Presets[len(minesweeper.Presets)-1].Config

	if got := cyclePreset(first, -1); got != last {
		t.Fatalf("cycling back from first = %+v, want %+v", got, last)
	}
	if got := cyclePreset(last, 1); got != first {
		t.Fatalf("cycling forward from last = %+v, want %+v", got, first)
	}
}

func TestCyclePresetFromCustom(t *testing.T) {
	custom := minesweeper.Config{Width: 12, Height: 7, Mines: 15}
	if presetName(custom) != "Custom" {
		t.Fatalf("presetName = %q, want Custom", presetName(custom))
	}
	if got := cyclePreset(custom, 1); got != minesweeper.Presets[0].Config {
		t.Fatalf("cycling forward from custom = %+v, want first preset", got)
	}
}

func TestSnakeRowsClampAndToggle(t *testing.T) {
	s := Default()
	speedRow, wallsRow := SnakeRows[0], SnakeRows[3]

	for i := 0; i < len(snake.Speeds)+2; i++ {
		speedRow.Adjust(&s, 1)
	}
	if got, want := speedRow.Value(s), snake.Speeds[len(snake.Speeds)-1].Name; got != want {
		t.Fatalf("speed = %q, want clamped to %q", got, want)
	}

	walls := s.Snake.Walls
	wallsRow.Adjust(&s, -1)
	if s.Snake.Walls == walls {
		t.Fatal("adjusting walls didn't toggle them")
	}
}

func TestShrinkingBoardClampsMines(t *testing.T) {
	s := Settings{Minesweeper: minesweeper.Config{Width: 5, Height: 5, Mines: 16}}
	widthRow := MinesweeperRows[1]

	widthRow.Adjust(&s, -1)
	if s.Minesweeper.Width != minesweeper.MinSize {
		t.Fatalf("width = %d, want clamped to %d", s.Minesweeper.Width, minesweeper.MinSize)
	}

	s.Minesweeper = minesweeper.Config{Width: 6, Height: 5, Mines: 21}
	widthRow.Adjust(&s, -1)
	if want := s.Minesweeper.MaxMines(); s.Minesweeper.Mines != want {
		t.Fatalf("mines = %d, want clamped to %d", s.Minesweeper.Mines, want)
	}
}
