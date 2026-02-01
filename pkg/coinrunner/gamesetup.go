package coinrunner

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

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
	return flickerTickCmd()
}

func (m GeneralModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// common messages that apply to all rooms
	switch msg := msg.(type) {
	case tickMsg:
		m.UIData.Flicker = !m.UIData.Flicker
		return m, flickerTickCmd()
	case tea.WindowSizeMsg:
		m = WindowSizeUpdate(m, msg)
		return m, nil
	case roomChangeMsg:
		m = RoomChangeUIReset(m)
		return m, nil
	}

	// pick to the proper update method based on current game state
	switch m.GameData.CurrentState {
	case StartPage:
		return StartRoomUpdate(m, msg)
	case ProloguePage:
		return PrologueRoomUpdate(m, msg)
	}
	return GameRoomUpdate(m, msg)
}

func (m GeneralModel) View() string {

	switch m.GameData.CurrentState {
	case StartPage:
		return RenderStartPage(m)
	case ProloguePage:
		return RenderProloguePage(m)
	}

	return RenderStartPage(m)
}
