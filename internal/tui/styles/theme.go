package styles

import (
	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	FooterNameBG      lipgloss.Color
	FooterNameFGFrom  lipgloss.Color
	FooterNameFGTo    lipgloss.Color
	Accent            lipgloss.Color
	HeadingForeground lipgloss.Color
	FooterSegmentBG   lipgloss.Color
	FooterSegmentFG   lipgloss.Color
	HelpFG            lipgloss.Color
}

var (
	AppTheme Theme
)

func init() {
	SetTheme(DefaultTheme)
}

func SetTheme(newTheme Theme) {
	AppTheme = newTheme

	// reset styles
	initRequestStyles()
	initCollectionStyles()
	initAppStyles()
}
