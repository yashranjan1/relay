package styles

import (
	"charm.land/lipgloss/v2"
)

var (
	LoaderMessageStyle = lipgloss.NewStyle().
		PaddingRight(2).
		Foreground(AppTheme.Accent).
		Background(AppTheme.FooterSegmentBG)
)

func initLoaderStyles() {
	LoaderMessageStyle = lipgloss.NewStyle().
		PaddingRight(2).
		Foreground(AppTheme.Accent).
		Background(AppTheme.FooterSegmentBG)
}
