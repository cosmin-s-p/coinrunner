package coinrunner

import (
	"coinrunner/pkg/helpers"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderRoom(m GeneralModel) string {
	titleStyle := HeaderStyle.Width(m.UIData.WindowWidth).
		Height(m.UIData.TitleHeight).
		Padding(0)

	sideWindowStyle := DefaultStyle.
		Width(m.UIData.WindowWidth/4 - 2).
		Height(m.UIData.WindowHeight - titleStyle.GetHeight()).
		Padding(0).
		Border(lipgloss.NormalBorder())
	mainWindowStyle := DefaultStyle.
		Width(m.UIData.WindowWidth/2 - 1).
		Height(m.UIData.WindowHeight - titleStyle.GetHeight()).
		Padding(0).
		Border(lipgloss.NormalBorder())

	title := titleStyle.Render("\n\n\n\nYou arrived now at " + m.WorldData.Rooms[m.GameData.CurrentState].Name + "\n")

	leftWindow := ""
	for i, v := range m.WorldData.Rooms[m.GameData.CurrentState].Choices {
		if m.UIData.Cursor == i && !m.UIData.Flicker {
			leftWindow += "[•] "
		} else {
			leftWindow += "[ ] "
		}
		leftWindow += v.String()
		leftWindow += "\n"
	}

	leftWindow = sideWindowStyle.Render(leftWindow)

	mainContent := m.WorldData.Rooms[m.GameData.CurrentState].Description
	if m.GameData.CurrentState == MerchantGate {
		mainContent = strings.Replace(mainContent+"\n", "<insert-favorite-item>", m.GameData.FavoriteItem, 1)
	}
	if m.GameData.LatestDialogue != "" {
		mainContent += "\n\n" + m.GameData.LatestDialogue + "\n"
	}
	mainWindow := mainWindowStyle.Render(mainContent)

	rightWindow := fmt.Sprintf("%s\n%s\n%s\n",
		helpers.DialogueHistoryHeaderView(m.UIData.SidePanelWidth),
		m.UIData.Viewport.View(),
		helpers.DialogueHistoryFooterView(m.UIData.SidePanelWidth),
	)

	bottom := lipgloss.JoinHorizontal(lipgloss.Top, leftWindow, mainWindow, rightWindow)
	content := lipgloss.JoinVertical(lipgloss.Left, title, bottom)

	return content
}

func RenderStartPage(m GeneralModel) string {
	title := HeaderStyle.Render(m.WorldData.Rooms[m.GameData.CurrentState].Description + "\n")
	content := ""
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

func RenderProloguePage(m GeneralModel) string {

	content := ""
	content += m.WorldData.Rooms[m.GameData.CurrentState].Description + "\n\n\n"
	content += m.UIData.TextInput.View()
	content = DefaultStyle.Render(content)

	return lipgloss.Place(
		m.UIData.WindowWidth,
		m.UIData.WindowHeight,
		lipgloss.Center,
		lipgloss.Center,
		content,
		lipgloss.WithWhitespaceBackground(DefaultStyle.GetBackground()),
	)
}
