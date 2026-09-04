package methodpicker

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/yashranjan1/relay/internal/tui/keybinds"
)

type MethodPickerConfig[T any] struct {
	Items            []T
	ItemMapper       func([]T) []list.Item
	FilteringEnabled bool

	Delegate          func(bool) list.ItemDelegate
	KeyMap            *keybinds.ListKeyMap
	AdditionalKeymaps *keybinds.ListKeyMap

	OnSelectAction tea.Msg

	ShowPagination bool
	ShowStatusBar  bool
	ShowHelp       bool
	ShowTitle      bool
	Width, Height  int
}

func config() {}
