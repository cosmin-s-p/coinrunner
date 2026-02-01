package coinrunner

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

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
