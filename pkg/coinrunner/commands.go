package coinrunner

import (
	"math/rand/v2"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// used to show blinking cursor for choices
func flickerTickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*250, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func roomChangeCmd(o GameState, n GameState) tea.Cmd {
	return func() tea.Msg {
		return roomChangeMsg{
			PreviousRoom: o,
			NewRoom:      n,
		}
	}
}

func dialogUpdateCmd() tea.Cmd {
	return func() tea.Msg {
		return dialogueUpdateMsg(true)
	}
}

func canNotMoveForwardCmd() tea.Cmd {
	return func() tea.Msg {
		return canNotMoveForwardMsg(true)
	}
}

func spawnCreaturesCmd(creaturesAvailable []Creature) tea.Cmd {

	// consider spawning a random creature from the ones available every 2 seconds
	return func() tea.Msg {
		time.Sleep(time.Second * 2)

		randomCreature := creaturesAvailable[rand.IntN(len(creaturesAvailable))]
		spawned := false

		// consider it spawned in 60% of the cases
		if rand.Float32()*100 < 60 {
			spawned = true
		}

		return creatureSpawnedMsg{spawned, randomCreature}
	}
}

func simulateParallelRequestsCmd() tea.Cmd {

	// create a new random token every 2 seconds to simulate other requests being processed in parallel
	return func() tea.Msg {
		time.Sleep(time.Second * 2)

		token := InitializeRandomToken()
		Memory = append(Memory, token)

		return simulateParallelRequestsMsg(true)
	}
}
