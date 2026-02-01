package coinrunner

import (
	"coinrunner/pkg/helpers"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func GameRoomUpdate(m GeneralModel, msg tea.Msg) (GeneralModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tickMsg:
		m.UIData.Flicker = !m.UIData.Flicker
		return m, flickerTickCmd()
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w":
			m.UIData.Cursor = helpers.CursorDecrease(m.UIData.Cursor)
		case "down", "s":
			length := len(m.WorldData.Rooms[m.GameData.CurrentState].Choices)
			m.UIData.Cursor = helpers.CursorIncrease(m.UIData.Cursor, length)
		case "enter", " ":
			action := m.WorldData.Rooms[m.GameData.CurrentState].Choices[m.UIData.Cursor]

			m.GameData, cmd = HandleAction(m.GameData, m.WorldData, action)

			return m, cmd

		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m = WindowSizeUpdate(m, msg)
	case roomChangeMsg:
		m = RoomChangeUIReset(m)
	}

	return m, cmd
}

func StartRoomUpdate(m GeneralModel, msg tea.Msg) (GeneralModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tickMsg:
		m.UIData.Flicker = !m.UIData.Flicker
		return m, flickerTickCmd()
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w":
			m.UIData.Cursor = helpers.CursorDecrease(m.UIData.Cursor)
		case "down", "s":
			length := len(m.WorldData.Rooms[m.GameData.CurrentState].Choices)
			m.UIData.Cursor = helpers.CursorIncrease(m.UIData.Cursor, length)
		case "enter", " ":
			action := m.WorldData.Rooms[m.GameData.CurrentState].Choices[m.UIData.Cursor]

			m.GameData, cmd = HandleAction(m.GameData, m.WorldData, action)

			return m, cmd

		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m = WindowSizeUpdate(m, msg)
	case roomChangeMsg:
		m = RoomChangeUIReset(m)
	}

	return m, cmd
}

func PrologueRoomUpdate(m GeneralModel, msg tea.Msg) (GeneralModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tickMsg:
		m.UIData.Flicker = !m.UIData.Flicker
		return m, flickerTickCmd()
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			m.GameData.FavoriteItem = m.UIData.TextInput.Value()

			action := m.WorldData.Rooms[m.GameData.CurrentState].Choices[m.UIData.Cursor]

			m.GameData, cmd = HandleAction(m.GameData, m.WorldData, action)

			return m, cmd

		case "esc", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m = WindowSizeUpdate(m, msg)
	case roomChangeMsg:
		m = RoomChangeUIReset(m)
	}

	m.UIData.TextInput, cmd = m.UIData.TextInput.Update(msg)

	return m, cmd
}

func RoomChangeUIReset(m GeneralModel) GeneralModel {
	m.UIData.Cursor = 0
	m.UIData.TextInput = textinput.New()

	if m.GameData.CurrentState == ProloguePage {
		ti := textinput.New()
		ti.Placeholder = " "
		ti.Focus()
		ti.CharLimit = 156
		ti.Width = 20

		m.UIData.TextInput = ti
	}

	return m
}

func WindowSizeUpdate(m GeneralModel, msg tea.WindowSizeMsg) GeneralModel {

	m.UIData.WindowWidth = msg.Width
	m.UIData.WindowHeight = msg.Height

	return m
}
