package styles

import "github.com/charmbracelet/lipgloss"

var (
	SelectedInactiveListStyle = lipgloss.NewStyle().PaddingLeft(1).Border(lipgloss.NormalBorder(), false, false, false, true)
	SelectedActiveListStyle   = lipgloss.NewStyle().Foreground(accent).PaddingLeft(1).Border(lipgloss.NormalBorder(), false, false, false, true).BorderForeground(accent)
	UnselectedListStyle       = lipgloss.NewStyle().PaddingLeft(2)
	InputStyle                = lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.NormalBorder(), false, false, false, true).Margin(1, 0)
	// Grad = lipg
)
