package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	footerNameStyle    = lipgloss.NewStyle().Bold(true).Background(AppTheme.FooterNameBG)
	footerNameBGStyle  = lipgloss.NewStyle().Background(AppTheme.FooterNameBG).Padding(0, 3, 0)
	FooterSegmentStyle = lipgloss.NewStyle().Background(AppTheme.FooterSegmentBG).PaddingLeft(2).Foreground(AppTheme.FooterSegmentFG)
	FooterVersionStyle = lipgloss.NewStyle().Background(AppTheme.FooterSegmentBG).AlignHorizontal(lipgloss.Right).PaddingRight(2).Foreground(AppTheme.FooterSegmentFG)
	TabHeadingInactive = lipgloss.NewStyle().Width(25).AlignHorizontal(lipgloss.Center).Border(lipgloss.NormalBorder(), false, false, false, true)
	TabHeadingActive   = lipgloss.NewStyle().Background(AppTheme.Accent).Foreground(AppTheme.HeadingForeground).Width(25).AlignHorizontal(lipgloss.Center).Border(lipgloss.NormalBorder(), false, false, false, true)
	HelpStyle          = lipgloss.NewStyle().Padding(1, 0, 1, 2)
	AppHelpStyle       = lipgloss.NewStyle().Padding(1, 0).Foreground(AppTheme.HelpFG)
	ErrorBarStyle      = lipgloss.NewStyle().Background(lipgloss.Color("#FF0000")).Foreground(lipgloss.Color("#FFFFFF")).Padding(0, 1)
)

func initAppStyles() {
	footerNameStyle = lipgloss.NewStyle().Bold(true).Background(AppTheme.FooterNameBG)
	footerNameBGStyle = lipgloss.NewStyle().Background(AppTheme.FooterNameBG).Padding(0, 3, 0)
	FooterSegmentStyle = lipgloss.NewStyle().Background(AppTheme.FooterSegmentBG).PaddingLeft(2).Foreground(AppTheme.FooterSegmentFG)
	FooterVersionStyle = lipgloss.NewStyle().Background(AppTheme.FooterSegmentBG).AlignHorizontal(lipgloss.Right).PaddingRight(2).Foreground(AppTheme.FooterSegmentFG)
	TabHeadingInactive = lipgloss.NewStyle().Width(25).AlignHorizontal(lipgloss.Center).Border(lipgloss.NormalBorder(), false, false, false, true)
	TabHeadingActive = lipgloss.NewStyle().Background(AppTheme.Accent).Foreground(AppTheme.HeadingForeground).Width(25).AlignHorizontal(lipgloss.Center).Border(lipgloss.NormalBorder(), false, false, false, true)
	HelpStyle = lipgloss.NewStyle().Padding(1, 0, 1, 2)
	AppHelpStyle = lipgloss.NewStyle().Padding(1, 0).Foreground(AppTheme.HelpFG)
	ErrorBarStyle = lipgloss.NewStyle().Background(lipgloss.Color("#FF0000")).Foreground(lipgloss.Color("#FFFFFF")).Padding(0, 1)
}
