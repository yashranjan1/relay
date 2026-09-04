package styles

import (
	"image/color"
)

type Theme struct {
	FooterNameBG      color.Color
	FooterNameFGFrom  color.Color
	FooterNameFGTo    color.Color
	Accent            color.Color
	HeadingForeground color.Color
	FooterSegmentBG   color.Color
	FooterSegmentFG   color.Color
	HelpFG            color.Color
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
