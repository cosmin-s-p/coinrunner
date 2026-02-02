package coinrunner

import (
	"coinrunner/pkg/helpers"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func GameRoomUpdate(m GeneralModel, msg tea.Msg) (GeneralModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w":
			m.UIData.Cursor = helpers.CursorDecrease(m.UIData.Cursor)
		case "down", "s":
			length := len(m.WorldData.Rooms[m.GameData.CurrentState].Choices)
			m.UIData.Cursor = helpers.CursorIncrease(m.UIData.Cursor, length)
		case "enter", " ":
			// check player choice and do the corresponding action
			action := m.WorldData.Rooms[m.GameData.CurrentState].Choices[m.UIData.Cursor]

			m.GameData, cmd = HandleAction(m.GameData, m.WorldData, action)

			return m, cmd

		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case dialogueUpdateMsg:
		// update the scrollable viewport content
		m.UIData.Viewport.SetContent(" \n" + strings.Join(m.GameData.DialogueHistory, "\n"))
	}

	// component specific update method
	m.UIData.Viewport, cmd = m.UIData.Viewport.Update(msg)

	return m, cmd
}

func StartRoomUpdate(m GeneralModel, msg tea.Msg) (GeneralModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w":
			m.UIData.Cursor = helpers.CursorDecrease(m.UIData.Cursor)
		case "down", "s":
			length := len(m.WorldData.Rooms[m.GameData.CurrentState].Choices)
			m.UIData.Cursor = helpers.CursorIncrease(m.UIData.Cursor, length)
		case "enter", " ":
			// check player choice and do the corresponding action
			action := m.WorldData.Rooms[m.GameData.CurrentState].Choices[m.UIData.Cursor]

			m.GameData, cmd = HandleAction(m.GameData, m.WorldData, action)

			return m, cmd

		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, cmd
}

func PrologueRoomUpdate(m GeneralModel, msg tea.Msg) (GeneralModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			// save the text input to the game state and reset the input
			m.GameData.FavoriteItem = m.UIData.TextInput.Value()
			m.UIData.TextInput = textinput.New()

			// check player choice and do the corresponding action
			action := m.WorldData.Rooms[m.GameData.CurrentState].Choices[m.UIData.Cursor]

			m.GameData, cmd = HandleAction(m.GameData, m.WorldData, action)

			return m, cmd

		case "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	// component specific update method
	m.UIData.TextInput, cmd = m.UIData.TextInput.Update(msg)

	return m, cmd
}

func RoomChangeUIReset(m GeneralModel, msg roomChangeMsg) (GeneralModel, tea.Cmd) {
	m.UIData.Cursor = 0

	// initialize text input upon entering the prologue page
	if msg.NewRoom == ProloguePage {
		ti := textinput.New()
		ti.Placeholder = " "
		ti.Focus()
		ti.CharLimit = 156
		ti.Width = 20

		m.UIData.TextInput = ti

		return m, textinput.Blink
	}

	return m, nil
}

func WindowSizeUpdate(m GeneralModel, msg tea.WindowSizeMsg) GeneralModel {

	m.UIData.WindowWidth = msg.Width
	m.UIData.WindowHeight = msg.Height
	m.UIData.SidePanelWidth = m.UIData.WindowWidth / 4
	m.UIData.SidePanelHeight = m.UIData.WindowHeight - m.UIData.TitleHeight

	// also set viewport size
	m.UIData.Viewport.Width = m.UIData.SidePanelWidth
	m.UIData.Viewport.Height = m.UIData.SidePanelHeight

	return m
}
