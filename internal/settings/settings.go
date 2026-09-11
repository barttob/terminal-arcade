package settings

import (
	"strconv"

	"github.com/barttob/terminal-arcade/internal/games/minesweeper"
)

// Settings holds per-game options passed to game factories.
type Settings struct {
	Minesweeper minesweeper.Config
}

func Default() Settings {
	return Settings{
		Minesweeper: minesweeper.DefaultConfig,
	}
}

// Row is one adjustable line in a game's settings tab.
type Row struct {
	Label  string
	Value  func(Settings) string
	Adjust func(s *Settings, delta int)
}

var MinesweeperRows = []Row{
	{
		Label:  "Difficulty",
		Value:  func(s Settings) string { return presetName(s.Minesweeper) },
		Adjust: func(s *Settings, delta int) { s.Minesweeper = cyclePreset(s.Minesweeper, delta) },
	},
	{
		Label: "Width",
		Value: func(s Settings) string { return strconv.Itoa(s.Minesweeper.Width) },
		Adjust: func(s *Settings, delta int) {
			s.Minesweeper.Width += delta
			s.Minesweeper = s.Minesweeper.Clamped()
		},
	},
	{
		Label: "Height",
		Value: func(s Settings) string { return strconv.Itoa(s.Minesweeper.Height) },
		Adjust: func(s *Settings, delta int) {
			s.Minesweeper.Height += delta
			s.Minesweeper = s.Minesweeper.Clamped()
		},
	},
	{
		Label: "Mines",
		Value: func(s Settings) string { return strconv.Itoa(s.Minesweeper.Mines) },
		Adjust: func(s *Settings, delta int) {
			s.Minesweeper.Mines += delta
			s.Minesweeper = s.Minesweeper.Clamped()
		},
	},
}

func presetIndex(c minesweeper.Config) int {
	for i, p := range minesweeper.Presets {
		if p.Config == c {
			return i
		}
	}
	return -1
}

func presetName(c minesweeper.Config) string {
	if i := presetIndex(c); i >= 0 {
		return minesweeper.Presets[i].Name
	}
	return "Custom"
}

// cyclePreset steps through the presets, wrapping at either end. A custom
// board steps to the first or last preset.
func cyclePreset(c minesweeper.Config, delta int) minesweeper.Config {
	n := len(minesweeper.Presets)
	i := presetIndex(c)
	switch {
	case i >= 0:
		i = (i + delta + n) % n
	case delta > 0:
		i = 0
	default:
		i = n - 1
	}
	return minesweeper.Presets[i].Config
}
