package keybinds

import (
	"charm.land/bubbles/v2/key"
)

type Keymaps struct {
	InsertItem           key.Binding
	DeleteItem           key.Binding
	EditItem             key.Binding
	Choose               key.Binding
	Remove               key.Binding
	Back                 key.Binding
	Up                   key.Binding
	Down                 key.Binding
	NextPage             key.Binding
	SendRequest          key.Binding
	Next                 key.Binding
	Prev                 key.Binding
	Save                 key.Binding
	PrevPage             key.Binding
	Filter               key.Binding
	ClearFilter          key.Binding
	CancelWhileFiltering key.Binding
	AcceptWhileFiltering key.Binding
	Quit                 key.Binding
}

var Keys = Keymaps{
	Back: key.NewBinding(
		key.WithKeys("ctrl+b"),
		key.WithHelp("ctrl+b", "back"),
	),
	SendRequest: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r", "send request"),
	),
	Save: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "save"),
	),
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	Prev: key.NewBinding(
		key.WithKeys("ctrl+h"),
		key.WithHelp("ctrl+h", "prev field"),
	),
	Next: key.NewBinding(
		key.WithKeys("ctrl+l"),
		key.WithHelp("ctrl+l", "next field"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	PrevPage: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "prev page"),
	),
	NextPage: key.NewBinding(
		key.WithKeys("right", "l", "pgdown"),
		key.WithHelp("→/l", "next page"),
	),
	Filter: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "filter"),
	),
	ClearFilter: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "clear filter"),
	),
	CancelWhileFiltering: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "cancel"),
	),
	AcceptWhileFiltering: key.NewBinding(
		key.WithKeys("enter", "tab", "shift+tab", "ctrl+k", "up", "ctrl+j", "down"),
		key.WithHelp("enter", "apply filter"),
	),
	InsertItem: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add"),
	),
	EditItem: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit"),
	),
	Choose: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Choose"),
	),
	Remove: key.NewBinding(
		key.WithKeys("x", "backspace"),
		key.WithHelp("x", "delete"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
