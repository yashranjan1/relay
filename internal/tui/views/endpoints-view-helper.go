package views

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/maniac-en/req/internal/tui/styles"
)

func createEndpointsDelegate(active bool) list.ItemDelegate {
	d := list.NewDefaultDelegate()
	if active {
		d.Styles.SelectedTitle = styles.SelectedActiveListStyle
	} else {
		d.Styles.SelectedTitle = styles.SelectedInactiveListStyle

	}
	d.Styles.SelectedDesc = styles.SelectedActiveListStyle

	return d
}
