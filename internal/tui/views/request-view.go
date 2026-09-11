package views

import (
	"context"
	"errors"
	"fmt"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/yashranjan1/relay/internal/backend/endpoints"
	"github.com/yashranjan1/relay/internal/backend/http"
	"github.com/yashranjan1/relay/internal/log"
	componenttypes "github.com/yashranjan1/relay/internal/tui/components/ComponentTypes"
	methodpicker "github.com/yashranjan1/relay/internal/tui/components/MethodPicker"
	optionsProvider "github.com/yashranjan1/relay/internal/tui/components/OptionsProvider"
	urlinput "github.com/yashranjan1/relay/internal/tui/components/UrlInput"
	viewport "github.com/yashranjan1/relay/internal/tui/components/ViewPort"
	"github.com/yashranjan1/relay/internal/tui/keybinds"
	"github.com/yashranjan1/relay/internal/tui/messages"
	"github.com/yashranjan1/relay/internal/tui/styles"
)

type reqFocused string

const (
	methodPicker = "method"
	urlInput     = "input"
	responseView = "responseView"
)

var componentList = []reqFocused{
	methodPicker,
	urlInput,
	responseView,
}

type RequestView struct {
	width      int
	focused    reqFocused
	components map[reqFocused]componenttypes.FocusableComponent
	index      int
	epManager  *endpoints.EndpointsManager
	height     int
	loading    bool
	help       help.Model
	keys       *keybinds.ListKeyMap
	client     *http.HTTPManager
	order      int
	update     func(context.Context, int64, endpoints.EndpointData) (endpoints.EndpointEntity, error)
	endpoint   endpoints.EndpointEntity
	collection optionsProvider.Option
}

func (r *RequestView) Init() tea.Cmd {
	return nil
}

func (r *RequestView) Name() string {
	return "Request"
}

func (r *RequestView) Help() []key.Binding {
	var reqViewBinds []key.Binding

	binds := r.components[r.focused].Help()
	reqViewBinds = []key.Binding{
		keybinds.Keys.Prev,
		keybinds.Keys.Next,
		keybinds.Keys.Save,
		keybinds.Keys.SendRequest,
	}
	reqViewBinds = append(binds, reqViewBinds...)

	return reqViewBinds
}

func (r *RequestView) GetFooterSegment() string {
	return fmt.Sprintf("%s/%s", r.collection.Name, r.endpoint.Name)
}

func (r *RequestView) Update(msg tea.Msg) (ViewInterface, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.height = msg.Height
		r.width = msg.Width
		w := r.components[methodPicker].GetWidth()
		r.components[urlInput].SetWidth(r.width - w)

		topMenu := 5
		r.components[responseView], cmd = r.components[responseView].Update(tea.WindowSizeMsg{
			Height: msg.Height - topMenu,
			Width:  r.width - topMenu,
		})
	case messages.Response:
		if msg.Err != nil {
			log.Error(msg.Err.Error())
			return r, func() tea.Msg {
				return messages.AddToast{
					Message: "Failed to get response",
				}
			}
		}
		if settable, ok := r.components[responseView].(componenttypes.ResponseSettable); ok {
			settable.SetState(msg.Data)
		}
		r.shiftFocusTo(responseView)
		r.loading = false
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, keybinds.Keys.Back):
			if settable, ok := r.components[responseView].(componenttypes.ResponseSettable); ok {
				settable.EraseState()
			}
			return r, func() tea.Msg {
				return messages.NavigateToView{
					ViewName: Endpoints,
					Target:   MainModel,
				}
			}
		case key.Matches(msg, keybinds.Keys.Next) || key.Matches(msg, keybinds.Keys.Prev) || key.Matches(msg, keybinds.Keys.Over) || key.Matches(msg, keybinds.Keys.Under):
			r.shift(msg)
			return r, func() tea.Msg {
				return messages.RefreshItemsList{}
			}
		case key.Matches(msg, keybinds.Keys.SendRequest):
			r.Save()
			request := &http.Request{
				Method: r.endpoint.Method,
				URL:    r.endpoint.Url,
			}
			r.loading = true
			sendMsg := func() tea.Msg {
				res, err := r.client.ExecuteRequest(request)
				if err != nil {
					return messages.Response{
						Err: err,
					}
				}

				data := messages.Response{
					Data: res,
				}

				return data
			}

			loaderMsg := func() tea.Msg {
				return messages.StartLoader{
					Message: " Sending",
				}
			}

			return r, tea.Batch(loaderMsg, sendMsg)

		case key.Matches(msg, keybinds.Keys.Save):
			r.Save()
		}
	}

	r.components[r.focused], cmd = r.components[r.focused].Update(msg)

	cmds = append(cmds, cmd)

	return r, tea.Batch(cmds...)
}

func (r *RequestView) Save() {
	for _, val := range r.components {
		if bindable, ok := val.(componenttypes.EndpointBindable); ok {
			r.endpoint = bindable.UpdateState(r.endpoint)
		}
	}
	r.update(
		context.Background(),
		r.endpoint.GetID(),
		endpoints.EndpointData{
			Name:   r.endpoint.Name,
			Method: r.endpoint.Method,
			URL:    r.endpoint.Url,
		},
	)
}

func (r *RequestView) shiftFocusTo(pane reqFocused) {
	r.components[r.focused].OnBlur()
	r.focused = pane
	r.components[r.focused].OnFocus()
}

func (r *RequestView) shift(msg tea.KeyPressMsg) {
	switch {
	case key.Matches(msg, keybinds.Keys.Next):
		r.index = (r.index + 1) % len(componentList)
	case key.Matches(msg, keybinds.Keys.Prev):
		r.index = (r.index - 1 + len(componentList)) % len(componentList)
	// INFO: bit stupid rn but will be cool when we have the request view
	case key.Matches(msg, keybinds.Keys.Under):
		r.index = (r.index + 2) % len(componentList)
	case key.Matches(msg, keybinds.Keys.Over):
		r.index = (r.index - 2 + len(componentList)) % len(componentList)
	default:
	}
	r.shiftFocusTo(componentList[r.index])
	r.Save()
}

func (r *RequestView) View() string {
	if r.endpoint.Name == "" {
		return styles.RequestLayout(r.height, r.width)("No Endpoint selected")
	}
	topMenu := lipgloss.JoinHorizontal(
		lipgloss.Left,
		r.components[methodPicker].View(),
		r.components[urlInput].View(),
	)
	botMenu := lipgloss.JoinHorizontal(
		lipgloss.Left,
		r.components[responseView].View(),
	)
	view := lipgloss.JoinVertical(
		lipgloss.Left,
		topMenu,
		botMenu,
	)
	return styles.RequestLayout(r.height, r.width)(view)
}

func (r *RequestView) SetState(items ...any) error {
	if len(items) == 1 {
		if data, ok := items[0].(EndpointData); ok {
			ep, err := r.epManager.Read(context.Background(), data.EndpointID)
			r.collection = data.Collection
			if err != nil {
				log.Error(err.Error())
				return err
			}
			r.endpoint = ep
			for _, val := range componentList {
				if bindable, ok := r.components[val].(componenttypes.EndpointBindable); ok {
					bindable.SetState(ep)
				}
			}
			if settable, ok := r.components[responseView].(componenttypes.ResponseSettable); ok {
				settable.EraseState()
			}
			return nil
		}
	}
	return errors.New("Invalid inputs, this function takes 1 input of type endpoints.Entity")
}

func (r *RequestView) Order() int {
	return r.order
}

func (r *RequestView) OnFocus() tea.Cmd {
	r.components[r.focused].OnFocus()
	return nil
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

type RequestViewConfig struct {
	EpManager *endpoints.EndpointsManager
	Update    func(context.Context, int64, endpoints.EndpointData) (endpoints.EndpointEntity, error)
	Order     int
	Client    *http.HTTPManager
}

func NewRequestView(cfg RequestViewConfig) *RequestView {
	mpConfig := createMethodPickerConfig()

	uiConfig := createURLInputConfig()

	return &RequestView{
		components: map[reqFocused]componenttypes.FocusableComponent{
			methodPicker: methodpicker.NewMethodPicker(mpConfig),
			urlInput:     urlinput.NewUrlInput(uiConfig),
			responseView: viewport.NewViewport(),
		},

		epManager: cfg.EpManager,
		focused:   methodPicker,
		order:     cfg.Order,
		update:    cfg.Update,
		client:    cfg.Client,
	}
}
