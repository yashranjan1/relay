package optionsProvider

import (
	"context"

	"charm.land/bubbles/v2/key"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/yashranjan1/relay/internal/tui/keybinds"
)

type ListConfig[T, U any] struct {
	OnSelectAction tea.Msg
	OnChangeAction func(int64)

	ShowPagination bool
	ShowStatusBar  bool
	ShowHelp       bool
	ShowTitle      bool
	Width, Height  int

	FilteringEnabled bool

	Placeholder       string
	Delegate          func(bool) list.ItemDelegate
	KeyMap            list.KeyMap
	AdditionalKeymaps *keybinds.ListKeyMap

	ItemMapper func([]T) []list.Item

	GetItemsFunc func(context.Context) ([]T, error)
	Source       string
	// Style    lipgloss.Style
}

type InputConfig struct {
	Prompt      string
	Placeholder string
	CharLimit   int
	Width       int
	KeyMap      InputKeyMaps
}

type InputKeyMaps struct {
	Accept key.Binding
	Back   key.Binding
}
