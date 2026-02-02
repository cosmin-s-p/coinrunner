package coinrunner

import (
	tea "github.com/charmbracelet/bubbletea"
)

func HandleAction(g GameData, w WorldData, c Choice) (GameData, tea.Cmd) {
	var commands []tea.Cmd

	// special cases
	if g.CurrentState == ProloguePage {
		g.DialogueHistory = append(g.DialogueHistory, "\t> Somebody just bought: "+g.FavoriteItem)
		commands = append(commands, dialogUpdateCmd())
	} else if g.CurrentState != StartPage {
		g.DialogueHistory = append(g.DialogueHistory, "\t> You chose to: "+c.String())
		commands = append(commands, dialogUpdateCmd())
	}

	switch c {
	case StartAction:
		commands = append(commands, roomChangeCmd(g.CurrentState, w.Rooms[g.CurrentState].NextRoom))
		g.CurrentState = w.Rooms[g.CurrentState].NextRoom

	case MoveForwardAction:
		// check if action is possible; player must be idle and able to proceed
		if !g.IsIdle || !g.CanMoveForward {
			g.LatestDialogue = "\t> You cannot move forward at this time."
			g.DialogueHistory = append(g.DialogueHistory, g.LatestDialogue)
			commands = append(commands, dialogUpdateCmd())
			return g, tea.Batch(commands...)
		}

		// check if this is a duplicate request
		if CheckInMemoryData(*g.Token) {
			g.CurrentState = GameOver
			g.LatestDialogue = "\t> You are a duplicate token and have been terminated!."
			return g, nil
		}

		// trigger room change cmd
		commands = append(commands, roomChangeCmd(g.CurrentState, w.Rooms[g.CurrentState].NextRoom))

		// update game state
		g.CurrentState = w.Rooms[g.CurrentState].NextRoom
		g.IsIdle = true
		g.CanMoveForward = true

		// update dialogue history and trigger dialog update cmd
		g.DialogueHistory = append(g.DialogueHistory, "\t> You arrived at "+g.CurrentState.String())
		commands = append(commands, dialogUpdateCmd())

	case QuitAction:
		return g, tea.Quit
	}

	return g, tea.Batch(commands...)
}
