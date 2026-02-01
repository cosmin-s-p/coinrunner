package helpers

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

func CursorIncrease(cursor int, length int) int {
	if cursor >= length-1 {
		return length - 1
	}

	return cursor + 1
}

func CursorDecrease(cursor int) int {
	if cursor <= 0 {
		return 0
	}

	return cursor - 1
}

func DialogueHistoryHeaderView(width int) string {
	title := "Dialogue History"
	line := strings.Repeat("─", width)

	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func GetDialogueHistoryViewport(content []string, width int, height int) viewport.Model {
	v := viewport.New(width, height)
	v.SetContent(" \n" + strings.Join(content, "\n"))

	return v
}

func DialogueHistoryFooterView(width int) string {
	line := strings.Repeat("─", width)

	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}
