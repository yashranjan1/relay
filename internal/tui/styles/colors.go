package styles

import "charm.land/lipgloss/v2"

var ThemeMap = map[string]Theme{
	"default":     DefaultTheme,
	"cyberpunk":   CyberpunkTheme,
	"light-paper": LightPaperTheme,
	"forest":      ForestTheme,
}

var DefaultTheme = Theme{
	FooterNameBG:      lipgloss.Color("#1a1a1a"),
	FooterNameFGFrom:  lipgloss.Color("#41A0AE"),
	FooterNameFGTo:    lipgloss.Color("#77F07F"),
	Accent:            lipgloss.Color("#77F07F"),
	HeadingForeground: lipgloss.Color("#000000"),
	FooterSegmentBG:   lipgloss.Color("#262626"),
	FooterSegmentFG:   lipgloss.Color("#656565"),
	HelpFG:            lipgloss.Color("#3C3C3C"),
	Error:             lipgloss.Color("#FF5555"),
	Warn:              lipgloss.Color("#D4BD5D"),
	Info:              lipgloss.Color("#14578F"),
	ToastBG:           lipgloss.Color("#000000"),
}

var LightPaperTheme = Theme{
	FooterNameBG:      lipgloss.Color("#E5E5E0"),
	FooterNameFGFrom:  lipgloss.Color("#006680"),
	FooterNameFGTo:    lipgloss.Color("#2AA198"),
	Accent:            lipgloss.Color("#2AA198"),
	HeadingForeground: lipgloss.Color("#073642"),
	FooterSegmentBG:   lipgloss.Color("#D3D3CD"),
	FooterSegmentFG:   lipgloss.Color("#586E75"),
	HelpFG:            lipgloss.Color("#93A1A1"),
	Error:             lipgloss.Color("#FF5555"),
	Warn:              lipgloss.Color("#D4BD5D"),
	Info:              lipgloss.Color("#14578F"),
	ToastBG:           lipgloss.Color("#000000"),
}

var ForestTheme = Theme{
	FooterNameBG:      lipgloss.Color("#18201E"),
	FooterNameFGFrom:  lipgloss.Color("#40A02B"), // Rich forest green
	FooterNameFGTo:    lipgloss.Color("#179299"), // Deep teal
	Accent:            lipgloss.Color("#40A02B"), // Darker green with much higher contrast
	HeadingForeground: lipgloss.Color("#E6E9EF"),
	FooterSegmentBG:   lipgloss.Color("#25322F"),
	FooterSegmentFG:   lipgloss.Color("#97B6AC"),
	HelpFG:            lipgloss.Color("#4F6C64"),
	Error:             lipgloss.Color("#FF5555"),
	Warn:              lipgloss.Color("#D4BD5D"),
	Info:              lipgloss.Color("#14578F"),
	ToastBG:           lipgloss.Color("#000000"),
}

var CyberpunkTheme = Theme{
	FooterNameBG:      lipgloss.Color("#181425"),
	FooterNameFGFrom:  lipgloss.Color("#FF2A6D"),
	FooterNameFGTo:    lipgloss.Color("#05D9E8"),
	Accent:            lipgloss.Color("#05D9E8"),
	HeadingForeground: lipgloss.Color("#FFFFFF"),
	FooterSegmentBG:   lipgloss.Color("#241734"),
	FooterSegmentFG:   lipgloss.Color("#8A78A5"),
	HelpFG:            lipgloss.Color("#62447D"),
	Error:             lipgloss.Color("#FF5555"),
	Warn:              lipgloss.Color("#D4BD5D"),
	Info:              lipgloss.Color("#14578F"),
	ToastBG:           lipgloss.Color("#000000"),
}
