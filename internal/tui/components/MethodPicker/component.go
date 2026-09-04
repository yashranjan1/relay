package methodpicker

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/yashranjan1/relay/internal/backend/endpoints"
	componenttypes "github.com/yashranjan1/relay/internal/tui/components/ComponentTypes"
	"github.com/yashranjan1/relay/internal/tui/keybinds"
)

type MethodPicker[T any] struct {
	list           list.Model
	onSelectAction tea.Msg
	delegateGenner func(bool) list.ItemDelegate
	keys           *keybinds.ListKeyMap
	state          endpoints.EndpointEntity
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

func (m *MethodPicker[T]) SetWidth(width int) {
	m.list.SetWidth(width)
}

func (m *MethodPicker[T]) GetWidth() int {
	return m.list.Width()
}

func (m *MethodPicker[T]) SetState(ep endpoints.EndpointEntity) {
	if index := m.findIndex(ep.Method); index >= 0 {
		m.list.Select(index)
		m.state = ep
	}
}

func (m *MethodPicker[T]) UpdateState(ep endpoints.EndpointEntity) endpoints.EndpointEntity {
	ep.Method = m.list.SelectedItem().(MethodOption).Name
	return ep
}

func (m *MethodPicker[T]) findIndex(method string) int {
	items := m.list.Items()
	for index, item := range items {
		if option, ok := item.(MethodOption); ok && option.Name == method {
			return index
		}
	}
	return -1
}

func (m *MethodPicker[T]) Update(msg tea.Msg) (componenttypes.ReqViewComponent, tea.Cmd) {
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
