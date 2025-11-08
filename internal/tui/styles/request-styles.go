package styles

import "github.com/charmbracelet/lipgloss"

func RequestLayout(height, width int) func(...string) string {
	return lipgloss.NewStyle().
		Height(height).
		Width(width).
		Padding(1, 0, 0, 1).
		BorderLeft(true).
		Render
}
