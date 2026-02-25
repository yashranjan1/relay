package methodpicker

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maniac-en/req/internal/tui/keybinds"
)

type MethodPickerConfig[T any] struct {
	Items            []T
	ItemMapper       func([]T) []list.Item
	FilteringEnabled bool

	Delegate          func(bool) list.ItemDelegate
	KeyMap            list.KeyMap
	AdditionalKeymaps *keybinds.ListKeyMap

	OnSelectAction tea.Msg

	ShowPagination bool
	ShowStatusBar  bool
	ShowHelp       bool
	ShowTitle      bool
	Width, Height  int
}

func config() {}
