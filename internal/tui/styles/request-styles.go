package styles

import "github.com/charmbracelet/lipgloss"

func RequestLayout(height, width int) func(...string) string {
	return lipgloss.NewStyle().
		Height(height).
		Width(width).
		Margin(0, 0, 0, 1).
		Padding(1, 0, 0, 0).
		Render
}

func UrlInputStyle(active bool) lipgloss.Style {
	var color lipgloss.Color

	if active {
		color = AppTheme.Accent
	}

	return lipgloss.NewStyle().
		Foreground(color).
		Margin(0, 0, 0, 5).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(color)
}

func ResponseStyle(height, width int) func(...string) string {
	return lipgloss.NewStyle().Margin(0, 0, 0, 5).Padding(1, 0, 0, 5).Height(height).Width(width).Render
}

var (
	ActiveRequestItem   = lipgloss.NewStyle().Foreground(AppTheme.Accent)
	InactiveRequestItem = lipgloss.NewStyle()
)

func initRequestStyles() {
	ActiveRequestItem = lipgloss.NewStyle().Foreground(AppTheme.Accent)
	InactiveRequestItem = lipgloss.NewStyle()
}
