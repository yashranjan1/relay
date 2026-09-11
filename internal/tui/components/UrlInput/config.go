package urlinput

import "charm.land/lipgloss/v2"

type UrlInputConfig struct {
	StyleGenner func(bool) lipgloss.Style
	Width       int
	Height      int
}
