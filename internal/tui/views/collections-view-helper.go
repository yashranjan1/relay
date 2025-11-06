package views

import (
	"github.com/charmbracelet/bubbles/list"
	optionsProvider "github.com/maniac-en/req/internal/tui/components/OptionsProvider"
	"github.com/maniac-en/req/internal/tui/keybinds"
	"github.com/maniac-en/req/internal/tui/styles"
)

func createCollectionDelegate() list.ItemDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = styles.SelectedListStyle
	d.Styles.SelectedDesc = styles.SelectedListStyle

	return d
}

func defaultListConfig[T, U any](binds *keybinds.ListKeyMap, delegateCreator func() list.ItemDelegate) *optionsProvider.ListConfig[T, U] {
	config := optionsProvider.ListConfig[T, U]{
		ShowPagination:   true,
		ShowStatusBar:    false,
		ShowHelp:         false,
		ShowTitle:        false,
		FilteringEnabled: true,
		Delegate:         delegateCreator(),
		KeyMap: list.KeyMap{
			CursorUp:             binds.CursorUp,
			CursorDown:           binds.CursorDown,
			NextPage:             binds.NextPage,
			PrevPage:             binds.PrevPage,
			Filter:               binds.Filter,
			ClearFilter:          binds.ClearFilter,
			CancelWhileFiltering: binds.CancelWhileFiltering,
			AcceptWhileFiltering: binds.AcceptWhileFiltering,
		},
	}
	return &config
}
