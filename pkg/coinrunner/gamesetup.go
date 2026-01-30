package coinrunner

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

var choices = []string{"Start Game", "Quit"}

func InitializeGame() {
	generalModel := GeneralModel{
		GameData: GameData{
			Token:        nil,
			CurrentState: StartPage,
		},
		UIData: UIData{
			Style: lipgloss.NewStyle().Bold(true).
				Foreground(lipgloss.Color("#FAFAFA")).
				Background(lipgloss.Color("#040011")).
				Align(lipgloss.Center).
				PaddingTop(2),
		},
	}

	p := tea.NewProgram(generalModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (m GeneralModel) Init() tea.Cmd {
	return tickCmd()
}

func (m GeneralModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.UIData.Flicker = !m.UIData.Flicker
		return m, tickCmd()
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w":
			m.UIData.Cursor--
			if m.UIData.Cursor < 0 {
				m.UIData.Cursor = 0
			}
		case "down", "s":
			m.UIData.Cursor++
			if m.UIData.Cursor >= len(choices) {
				m.UIData.Cursor = len(choices) - 1
			}
		case "enter", "space":
			// wip
			m.GameData.CurrentState++
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.UIData.WindowWidth = msg.Width
		m.UIData.WindowHeight = msg.Height
	}

	return m, nil
}

func (m GeneralModel) View() string {

	s := ""
	if m.GameData.CurrentState != StartPage {
		s += "You are now at " + m.GameData.CurrentState.String() + "\n\n\n"
	}
	for i := 0; i < len(choices); i++ {
		if m.GameData.CurrentState != StartPage {
			// wip
			continue
		}
		if m.UIData.Cursor == i && !m.UIData.Flicker {
			s += "[•] "
		} else {
			s += "[ ] "
		}
		s += choices[i]
		s += "\n"
	}
	s = m.UIData.Style.Render(s)

	return lipgloss.Place(
		m.UIData.WindowWidth,
		m.UIData.WindowHeight,
		lipgloss.Center,
		lipgloss.Center,
		s,
		lipgloss.WithWhitespaceBackground(m.UIData.Style.GetBackground()),
	)
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*250, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
