package styles

import "charm.land/lipgloss/v2"

var (
	SelectedInactiveListStyle = lipgloss.NewStyle().PaddingLeft(1).Border(lipgloss.NormalBorder(), false, false, false, true)
	SelectedActiveListStyle   = lipgloss.NewStyle().Foreground(AppTheme.Accent).PaddingLeft(1).Border(lipgloss.NormalBorder(), false, false, false, true).BorderForeground(AppTheme.Accent)
	UnselectedListStyle       = lipgloss.NewStyle().PaddingLeft(2)
	InputStyle                = lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.NormalBorder(), false, false, false, true).Margin(1, 0)
	// Grad = lipg
)

func initCollectionStyles() {

	SelectedInactiveListStyle = lipgloss.NewStyle().PaddingLeft(1).Border(lipgloss.NormalBorder(), false, false, false, true)
	SelectedActiveListStyle = lipgloss.NewStyle().Foreground(AppTheme.Accent).PaddingLeft(1).Border(lipgloss.NormalBorder(), false, false, false, true).BorderForeground(AppTheme.Accent)
	UnselectedListStyle = lipgloss.NewStyle().PaddingLeft(2)
	InputStyle = lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.NormalBorder(), false, false, false, true).Margin(1, 0)
	// Grad = lipg
}
