package views

import (
	"charm.land/bubbles/v2/list"
	"github.com/yashranjan1/relay/internal/tui/styles"
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
