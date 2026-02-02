package coinrunner

import "time"

type tickMsg time.Time

type roomChangeMsg struct {
	PreviousRoom GameState
	NewRoom      GameState
}

type dialogueUpdateMsg bool

type canNotMoveForwardMsg bool

type creatureSpawnedMsg struct {
	spawned  bool
	creature Creature
}

type simulateParallelRequestsMsg bool
