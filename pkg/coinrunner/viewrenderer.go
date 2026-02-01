package coinrunner

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func RenderStartPage(m GeneralModel) string {
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
	content += "Favorite item: " + m.GameData.FavoriteItem + "\n"
	content += "Cursor: " + strconv.Itoa(m.UIData.Cursor) + "\n"
	content += "Cursor: " + strconv.FormatBool(m.UIData.Flicker) + "\n"
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
