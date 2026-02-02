package coinrunner

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func InitializeGame(cfg Config) {
	token := InitializeRandomToken()
	generalModel := GeneralModel{
		WorldData: InitWorld(),
		GameData: GameData{
			Token:          &token,
			CurrentState:   cfg.GetGameState("start-room"),
			IsIdle:         true,
			CanMoveForward: true,
		},
		UIData: UIData{
			TextInput:   textinput.New(),
			Viewport:    viewport.New(0, 0),
			TitleHeight: cfg.GetInt("title-height"),
		},
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
	var cmd tea.Cmd
	// common messages that apply to all rooms
	switch msg := msg.(type) {
	case tickMsg:
		m.UIData.Flicker = !m.UIData.Flicker
		return m, flickerTickCmd()
	case tea.WindowSizeMsg:
		m = WindowSizeUpdate(m, msg)
		return m, nil
	case roomChangeMsg:
		m, cmd = RoomChangeUIReset(m, msg)
		// add cmd to spawn creatures
		// if new room is merchant gate add command to start creating tokens
		return m, cmd
	case canNotMoveForwardMsg:
		// just update the ui
		return m, nil
	}

	// pick the proper update method based on current game state
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

	return RenderRoom(m)
}
