package styles

import "charm.land/lipgloss/v2"

func ErrorToast(msg string) string {
	header := lipgloss.NewStyle().Foreground(AppTheme.Error).Render("󰀦 Error")
	joined := lipgloss.JoinVertical(lipgloss.Left, header, msg)

	return lipgloss.NewStyle().PaddingLeft(1).Height(4).Width(20).BorderLeft(true).BorderForeground(AppTheme.Error).Render(joined)
}

func WarnToast(msg string) string {
	header := lipgloss.NewStyle().Foreground(AppTheme.Warn).Render("󰀦 Warn")
	joined := lipgloss.JoinVertical(lipgloss.Left, header, msg)

	return lipgloss.NewStyle().PaddingLeft(1).Height(4).Width(20).BorderLeft(true).BorderForeground(AppTheme.Warn).Render(joined)
}

func InfoToast(msg string) string {
	header := lipgloss.NewStyle().Foreground(AppTheme.Info).Render(" Info")
	joined := lipgloss.JoinVertical(lipgloss.Left, header, msg)

	return lipgloss.NewStyle().PaddingLeft(1).Height(4).Width(20).BorderLeft(true).BorderForeground(AppTheme.Info).Render(joined)
}
