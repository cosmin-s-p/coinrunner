package coinrunner

import (
	tea "github.com/charmbracelet/bubbletea"
)

func HandleAction(g GameData, w WorldData, c Choice) (GameData, tea.Cmd) {
	var cmd tea.Cmd
	var updDiagCmd tea.Cmd

	if g.CurrentState == ProloguePage {
		g.DialogueHistory = append(g.DialogueHistory, "\t> Somebody just bought: "+g.FavoriteItem)
		updDiagCmd = dialogUpdateCmd()
	} else if g.CurrentState != StartPage {
		g.DialogueHistory = append(g.DialogueHistory, "\t> You chose to: "+c.String())
		updDiagCmd = dialogUpdateCmd()
	}

	switch c {
	case StartAction, MoveForwardAction:
		cmd = roomChangeCmd(g.CurrentState, w.Rooms[g.CurrentState].NextRoom)
		g.CurrentState = w.Rooms[g.CurrentState].NextRoom
	case QuitAction:
		cmd = tea.Quit
	}

	return g, tea.Batch(cmd, updDiagCmd)
}
