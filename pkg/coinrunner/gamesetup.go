package coinrunner

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

// var choices = []string{"Start Game", "Quit"}

func InitializeGame(cfg Config) {
	generalModel := GeneralModel{
		WorldData: InitWorld(),
		GameData: GameData{
			Token:        nil,
			CurrentState: cfg.GetGameState("startroom"),
		},
		UIData: UIData{},
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
			if m.UIData.Cursor >= len(m.WorldData.Rooms[m.GameData.CurrentState].Choices) {
				m.UIData.Cursor = len(m.WorldData.Rooms[m.GameData.CurrentState].Choices) - 1
			}
		case "enter", "space":
			// no choices? always have move forward choice? what about last room
			switch m.WorldData.Rooms[m.GameData.CurrentState].Choices[m.UIData.Cursor] {
			case StartAction:
				m.GameData.CurrentState++
				if int(m.GameData.CurrentState) >= len(m.WorldData.Rooms) {
					m.GameData.CurrentState--
				}
			case QuitAction:
				return m, tea.Quit
			default:
				return m, nil
			}

		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.UIData.WindowWidth = msg.Width
		m.UIData.WindowHeight = msg.Height
		DefaultStyle.Width(msg.Width)
	}

	return m, nil
}

func (m GeneralModel) View() string {

	title := ""
	content := ""
	if m.GameData.CurrentState != StartPage {
		title += HeaderStyle.Render("You arrived now at " + m.WorldData.Rooms[m.GameData.CurrentState].Name + "\n")
		content += m.WorldData.Rooms[m.GameData.CurrentState].Description + "\n\n\n"
	}
	for i, v := range m.WorldData.Rooms[m.GameData.CurrentState].Choices {
		if m.UIData.Cursor == i && !m.UIData.Flicker {
			content += "[•] "
		} else {
			content += "[ ] "
		}
		content += v.String()
		content += "\n"
	}
	content = DefaultStyle.Render(content)

	return lipgloss.Place(
		m.UIData.WindowWidth,
		m.UIData.WindowHeight,
		lipgloss.Center,
		lipgloss.Center,
		title+content,
		lipgloss.WithWhitespaceBackground(DefaultStyle.GetBackground()),
	)
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*250, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
