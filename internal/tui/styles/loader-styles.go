package styles

import (
	"charm.land/lipgloss/v2"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

func ParseHex(s string) (color.RGBA, error) {
	s = strings.TrimPrefix(s, "#")

	if len(s) == 3 {
		s = string([]byte{s[0], s[0], s[1], s[1], s[2], s[2]})
	}

	if len(s) != 6 {
		return color.RGBA{}, fmt.Errorf("invalid hex color: %q", s)
	}

	val, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("invalid hex color: %q", s)
	}

	return color.RGBA{
		R: uint8(val >> 16),
		G: uint8(val >> 8),
		B: uint8(val),
		A: 255,
	}, nil
}

func FakeOpacity(fg, bg color.RGBA, alpha float64) (string, error) {

	blend := func(fg, bg uint8) uint8 {
		return uint8(float64(fg)*alpha + float64(bg)*(1-alpha))
	}

	result := color.RGBA{
		R: blend(fg.R, bg.R),
		G: blend(fg.G, bg.G),
		B: blend(fg.B, bg.B),
	}

	return fmt.Sprintf("#%02X%02X%02X", result.R, result.G, result.B), nil
}

func GetLoaderCharStyle(alpha float64) func(...string) string {
	fgR, fgG, fgB, _ := AppTheme.Accent.RGBA()

	fg := color.RGBA{
		R: uint8(fgR),
		G: uint8(fgG),
		B: uint8(fgB),
	}

	bgR, bgG, bgB, _ := AppTheme.FooterSegmentBG.RGBA()

	bg := color.RGBA{
		R: uint8(bgR),
		G: uint8(bgG),
		B: uint8(bgB),
	}

	color, _ := FakeOpacity(fg, bg, alpha)

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(color)).
		Background(AppTheme.FooterSegmentBG).
		Render
}

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
