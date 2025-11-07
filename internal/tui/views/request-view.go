package views

import (
	"errors"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maniac-en/req/internal/backend/database"
	methodpicker "github.com/maniac-en/req/internal/tui/components/MethodPicker"
	"github.com/maniac-en/req/internal/tui/keybinds"
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

	return r, cmd
}

func (r RequestView) View() string {
	return r.methodPicker.View()
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

func NewRequestView() *RequestView {
	return &RequestView{
		methodPicker: methodpicker.NewMethodPicker[string](),
	}
}
