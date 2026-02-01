package coinrunner

import tea "github.com/charmbracelet/bubbletea"

func HandleAction(g GameData, w WorldData, c Choice) (GameData, tea.Cmd) {
	var cmd tea.Cmd
	switch c {
	case StartAction, MoveForwardAction:
		cmd = roomChangeCmd(g.CurrentState, w.Rooms[g.CurrentState].NextRoom)
		g.CurrentState = w.Rooms[g.CurrentState].NextRoom
	case QuitAction:
		cmd = tea.Quit
	}

	return g, cmd
}
