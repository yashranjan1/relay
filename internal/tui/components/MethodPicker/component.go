package methodpicker

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	componenttypes "github.com/maniac-en/req/internal/tui/components/ComponentTypes"
	"github.com/maniac-en/req/internal/tui/keybinds"
)

type MethodPicker[T any] struct {
	list           list.Model
	onSelectAction tea.Msg
	delegateGenner func(bool) list.ItemDelegate
	keys           *keybinds.ListKeyMap
	width          int
	height         int
	itemMapper     func([]T) []list.Item
}

type MethodOption struct {
	Name string
	Type string
}

func (m MethodOption) Title() string       { return m.Name }
func (m MethodOption) Description() string { return m.Name }
func (m MethodOption) Value() string       { return m.Type }
func (m MethodOption) FilterValue() string { return m.Name }

func (m *MethodPicker[T]) Init() tea.Cmd {
	return nil
}

func (m *MethodPicker[T]) SetSize(width, height int) {
}

func (m *MethodPicker[T]) Update(msg tea.Msg) (componenttypes.ComponentInterface, tea.Cmd) {
	var cmd tea.Cmd

	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m *MethodPicker[T]) View() string {
	return m.list.View()
}

func (m *MethodPicker[T]) OnFocus() {
	m.list.SetDelegate(m.delegateGenner(true))
}

func (m *MethodPicker[T]) OnBlur() {
	m.list.SetDelegate(m.delegateGenner(false))
}

func (m *MethodPicker[T]) GetSelected() MethodOption {
	return m.list.SelectedItem().(MethodOption)
}

func (m *MethodPicker[T]) Help() []key.Binding {
	return []key.Binding{
		m.keys.CursorDown,
		m.keys.CursorUp,
	}
}

func NewMethodPicker[T any](config MethodPickerConfig[T]) *MethodPicker[T] {
	method := &MethodPicker[T]{}

	picker := list.New(
		config.ItemMapper(config.Items),
		config.Delegate(false),
		config.Width,
		config.Height,
	)

	picker.SetFilteringEnabled(config.FilteringEnabled)
	picker.SetShowStatusBar(config.ShowStatusBar)
	picker.SetShowPagination(config.ShowPagination)
	picker.SetShowHelp(config.ShowHelp)
	picker.SetShowTitle(config.ShowTitle)

	method.list = picker
	method.keys = config.KeyMap
	method.delegateGenner = config.Delegate

	return method
}
