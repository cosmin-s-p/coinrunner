package coinrunner

import "time"

type tickMsg time.Time

type roomChangeMsg struct {
	PreviousRoom GameState
	NewRoom      GameState
}
