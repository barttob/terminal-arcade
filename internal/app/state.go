package app

type AppState int

const (
	StateMainMenu AppState = iota
	StatePlaying
	StatePaused
	StateGameOver
	StateSettings
	StateHelp
	StateLobby
	StateRoom
)
