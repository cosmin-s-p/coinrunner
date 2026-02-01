package coinrunner

import "github.com/charmbracelet/lipgloss"

var DefaultStyle = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#040011")).
	Align(lipgloss.Center, lipgloss.Center).
	PaddingTop(2)

var HeaderStyle = DefaultStyle.Foreground(lipgloss.Color("#d897c8"))
