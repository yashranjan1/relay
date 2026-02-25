package styles

import "github.com/charmbracelet/lipgloss"

func RequestLayout(height, width int) func(...string) string {
	return lipgloss.NewStyle().
		Height(height).
		Width(width).
		Margin(0, 0, 0, 5).
		Padding(1, 0, 0, 10).
		Border(lipgloss.NormalBorder(), false, false, false, true).
		Render
}

func UrlInputStyle(active bool) lipgloss.Style {
	var color lipgloss.Color

	if active {
		color = accent
	}

	return lipgloss.NewStyle().
		Foreground(color).
		Margin(0, 0, 0, 5).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(color)
}

var (
	ActiveRequestItem   = lipgloss.NewStyle().Foreground(accent)
	InactiveRequestItem = lipgloss.NewStyle()
)
