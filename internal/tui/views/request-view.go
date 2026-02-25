package views

import (
	"errors"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maniac-en/req/internal/backend/database"
	componenttypes "github.com/maniac-en/req/internal/tui/components/ComponentTypes"
	methodpicker "github.com/maniac-en/req/internal/tui/components/MethodPicker"
	urlinput "github.com/maniac-en/req/internal/tui/components/UrlInput"
	"github.com/maniac-en/req/internal/tui/keybinds"
	"github.com/maniac-en/req/internal/tui/messages"
	"github.com/maniac-en/req/internal/tui/styles"
)

type reqFocused string

const (
	methodPicker = "method"
	urlInput     = "input"
)

var componentList = []reqFocused{
	methodPicker,
	urlInput,
}

type RequestView struct {
	width      int
	focused    reqFocused
	components map[reqFocused]componenttypes.ComponentInterface
	index      int
	height     int
	help       help.Model
	keys       *keybinds.ListKeyMap
	order      int
	endpoint   database.Endpoint
}

func (r *RequestView) Init() tea.Cmd {
	return nil
}

func (r *RequestView) Name() string {
	return "Request"
}

func (r *RequestView) Help() []key.Binding {
	binds := r.components[r.focused].Help()
	return append(binds, keybinds.Keys.Next)
}

func (r *RequestView) GetFooterSegment() string {
	return ""
}

func (r *RequestView) Update(msg tea.Msg) (ViewInterface, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.height = msg.Height
		r.width = msg.Width
		r.components[urlInput].SetSize(r.width-20, r.height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keybinds.Keys.Back):
			return r, func() tea.Msg {
				return messages.NavigateToView{
					ViewName: Endpoints,
					Target:   Endpoints,
				}
			}
		case key.Matches(msg, keybinds.Keys.Next):
			r.index = (r.index + 1) % len(componentList)
			r.components[r.focused].OnBlur()
			r.focused = componentList[r.index]
			r.components[r.focused].OnFocus()
			return r, nil
		}
	}

	r.components[r.focused], cmd = r.components[r.focused].Update(msg)
	cmds = append(cmds, cmd)

	return r, tea.Batch(cmds...)
}

func (r *RequestView) View() string {
	views := []string{}
	for _, val := range componentList {
		views = append(views, r.components[val].View())
	}
	view := lipgloss.JoinHorizontal(
		lipgloss.Left,
		views...,
	)
	return styles.RequestLayout(r.height, r.width)(view)
}

func (r *RequestView) SetState(items ...any) error {
	if len(items) != 1 {
		return errors.New("Incorrect amount of fields supplied")
	}
	// r.SetState(items[0])
	return nil
}

func (r *RequestView) Order() int {
	return r.order
}

func (r *RequestView) OnFocus() {
	r.components[r.focused].OnFocus()
}

func (r *RequestView) OnBlur() {
	r.components[r.focused].OnBlur()
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
	mpConfig := createMethodPickerConfig()

	uiConfig := createURLInputConfig()

	return &RequestView{
		components: map[reqFocused]componenttypes.ComponentInterface{
			methodPicker: methodpicker.NewMethodPicker(mpConfig),
			urlInput:     urlinput.NewUrlInput(uiConfig),
		},
		focused: methodPicker,
	}
}
