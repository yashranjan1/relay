package urlinput

import "github.com/charmbracelet/lipgloss"

type UrlInputConfig struct {
	StyleGenner func(bool) lipgloss.Style
	Width       int
	Height      int
}
