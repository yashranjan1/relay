package styles

import (
	"image/color"
	"strconv"

	"charm.land/lipgloss/v2"
)

func RequestLayout(height, width int) func(...string) string {
	return lipgloss.NewStyle().
		Height(height).
		Width(width).
		Padding(1, 0, 0, 1).
		Render
}

func UrlInputStyle(active bool) lipgloss.Style {
	var fg color.Color

	if active {
		fg = AppTheme.Accent
	}

	return lipgloss.NewStyle().
		Foreground(fg).
		Margin(0, 0, 0, 5).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(fg)
}

func ResponseContentStyle(height, width int) func(...string) string {
	return lipgloss.NewStyle().Padding(1, 0, 1, 1).Height(height).Width(width).Render
}

func ResponseStyle(active bool) func(...string) string {
	if active {
		return lipgloss.NewStyle().BorderForeground(AppTheme.Accent).Border(lipgloss.RoundedBorder(), true).Render
	}
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true).Render
}

func ResponseFooterStyle(width int) func(...string) string {
	return lipgloss.NewStyle().
		Width(width).
		Align(lipgloss.Right).
		Render
}

func ResponseStatus(code int) string {
	var style = lipgloss.NewStyle()
	stringCode := strconv.Itoa(code)
	switch {
	case code < 300:
		style = style.Foreground(AppTheme.Success)
	case code < 400:
		style = style.Foreground(AppTheme.Info)
	case code < 600:
		style = style.Foreground(AppTheme.Error)
	default:
		style = style.Foreground(AppTheme.Success)
	}
	return style.Render(stringCode)
}

var (
	ActiveRequestItem   = lipgloss.NewStyle().Foreground(AppTheme.Accent)
	InactiveRequestItem = lipgloss.NewStyle()
)

func initRequestStyles() {
	ActiveRequestItem = lipgloss.NewStyle().Foreground(AppTheme.Accent)
	InactiveRequestItem = lipgloss.NewStyle()

}
