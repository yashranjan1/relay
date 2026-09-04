package styles

import (
	stdcolor "image/color"

	"charm.land/lipgloss/v2"
)

func gradientText(text string, startColor, endColor stdcolor.Color, base, additional lipgloss.Style) string {
	n := len(text)
	result := ""

	for i := range n {
		ratio := float64(i) / float64(n-1)
		color := interpolateColor(startColor, endColor, ratio)

		style := base.Foreground(color)
		result += style.Render(string(text[i]))
	}

	return additional.Render(result)
}

func interpolateColor(start, end stdcolor.Color, ratio float64) stdcolor.Color {
	r1, g1, b1, _ := start.RGBA()
	r2, g2, b2, _ := end.RGBA()

	r := uint8(float64(r1>>8) + (float64(r2>>8)-float64(r1>>8))*ratio)
	g := uint8(float64(g1>>8) + (float64(g2>>8)-float64(g1>>8))*ratio)
	b := uint8(float64(b1>>8) + (float64(b2>>8)-float64(b1>>8))*ratio)

	return stdcolor.RGBA{R: r, G: g, B: b, A: 0xff}
}

func ApplyGradientToFooter(text string) string {
	return gradientText("REQ", AppTheme.FooterNameFGFrom, AppTheme.FooterNameFGTo, footerNameStyle, footerNameBGStyle)
}
