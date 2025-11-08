package views

import (
	"errors"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maniac-en/req/internal/backend/database"
	methodpicker "github.com/maniac-en/req/internal/tui/components/MethodPicker"
	"github.com/maniac-en/req/internal/tui/keybinds"
	"github.com/maniac-en/req/internal/tui/styles"
)

type RequestView struct {
	width        int
	height       int
	help         help.Model
	keys         *keybinds.ListKeyMap
	methodPicker methodpicker.MethodPicker[string]
	order        int
	endpoint     database.Endpoint
}

func (r RequestView) Init() tea.Cmd {
	return nil
}

func (r RequestView) Name() string {
	return "Request"
}

func (r RequestView) Help() []key.Binding {
	return []key.Binding{}
}

func (r RequestView) GetFooterSegment() string {
	return ""
}

func (r RequestView) Update(msg tea.Msg) (ViewInterface, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.height = msg.Height
		r.width = msg.Width
	}

	r.methodPicker, cmd = r.methodPicker.Update(msg)
	cmds = append(cmds, cmd)

	return r, tea.Batch(cmds...)
}

func (r RequestView) View() string {
	return styles.RequestLayout(10, 100)(r.methodPicker.View())
}

func (r RequestView) SetState(items ...any) error {
	if len(items) != 1 {
		return errors.New("Incorrect amount of fields supplied")
	}
	r.SetState(items[0])
	return nil
}

func (r RequestView) Order() int {
	return r.order
}

func (r RequestView) OnFocus() {

}

func (r RequestView) OnBlur() {

}

func methodItemMapper(items []string) []list.Item {
	listItems := make([]list.Item, len(items))
	for index, item := range items {
		listItems[index] = methodpicker.MethodOption{
			Name: item,
			Type: item,
		}
	}
	return listItems
}

func NewRequestView() *RequestView {
	methods := []string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
	}

	config := methodpicker.MethodPickerConfig[string]{
		Items:            methods,
		ItemMapper:       methodItemMapper,
		FilteringEnabled: false,
		ShowPagination:   false,
		ShowStatusBar:    false,
		Delegate:         list.DefaultDelegate{},
		ShowHelp:         false,
		ShowTitle:        false,
		Width:            30,
		Height:           1,
	}
	return &RequestView{
		methodPicker: methodpicker.NewMethodPicker[string](config),
	}
}
