package methodpicker

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maniac-en/req/internal/tui/keybinds"
)

type MethodPicker[T any] struct {
	list           list.Model
	onSelectAction tea.Msg
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
func (m MethodOption) FilterValue() string { return "" }

func (m MethodPicker[T]) Init() tea.Cmd {
	return nil
}

func (m MethodPicker[T]) Update(msg tea.Msg) (MethodPicker[T], tea.Cmd) {
	return m, nil
}

func (m MethodPicker[T]) View() string {
	return "hello"
}

func (m *MethodPicker[T]) OnFocus() {
}

func (m MethodPicker[T]) OnBlur() {

}

func (m MethodPicker[T]) GetSelected() MethodOption {
	return m.list.SelectedItem().(MethodOption)
}

func (m *MethodPicker[T]) Help() []key.Binding {
	return []key.Binding{
		m.keys.NextPage,
		m.keys.PrevPage,
		m.keys.CursorUp,
		m.keys.CursorDown,
		m.keys.Accept,
	}
}

func NewMethodPicker[T any]() MethodPicker[T] {

	return MethodPicker[T]{}
}
