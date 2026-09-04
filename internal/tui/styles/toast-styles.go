package styles

import "charm.land/lipgloss/v2"

func ErrorToast(msg string, width int, height int) string {
	header := lipgloss.NewStyle().Foreground(AppTheme.Error).Render("󰀦 Error")
	joined := lipgloss.JoinVertical(lipgloss.Left, header, msg)

	return lipgloss.
		NewStyle().
		PaddingLeft(1).
		Height(height).
		Width(width).
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(AppTheme.Error).
		Background(AppTheme.ToastBG).
		Render(joined)
}

func WarnToast(msg string, width int, height int) string {
	header := lipgloss.NewStyle().Foreground(AppTheme.Warn).Render("󰀦 Warn")
	joined := lipgloss.JoinVertical(lipgloss.Left, header, msg)

	return lipgloss.NewStyle().
		PaddingLeft(1).
		Height(height).
		Width(width).
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(AppTheme.Warn).
		Background(AppTheme.ToastBG).
		Render(joined)
}

func InfoToast(msg string, width int, height int) string {
	header := lipgloss.NewStyle().Foreground(AppTheme.Info).Render(" Info")
	joined := lipgloss.JoinVertical(lipgloss.Left, header, msg)

	return lipgloss.NewStyle().
		PaddingLeft(1).
		Height(height).
		Width(width).
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(AppTheme.Info).
		Background(AppTheme.ToastBG).
		Render(joined)
}
