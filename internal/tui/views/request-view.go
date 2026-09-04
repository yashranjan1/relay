package views

import (
	"context"
	"errors"
	"fmt"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/yashranjan1/relay/internal/backend/endpoints"
	"github.com/yashranjan1/relay/internal/backend/http"
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
)

var componentList = []reqFocused{
	methodPicker,
	urlInput,
}

type RequestView struct {
	width        int
	focused      reqFocused
	viewport     viewport.Viewport
	components   map[reqFocused]componenttypes.ReqViewComponent
	index        int
	epManager    *endpoints.EndpointsManager
	height       int
	loading      bool
	help         help.Model
	keys         *keybinds.ListKeyMap
	spinner      spinner.Model
	responsePage bool
	client       *http.HTTPManager
	order        int
	update       func(context.Context, int64, endpoints.EndpointData) (endpoints.EndpointEntity, error)
	endpoint     endpoints.EndpointEntity
	collection   optionsProvider.Option
}

func (r *RequestView) Init() tea.Cmd {
	return nil
}

func (r *RequestView) Name() string {
	return "Request"
}

func (r *RequestView) Help() []key.Binding {
	var reqViewBinds []key.Binding
	if r.responsePage {
		reqViewBinds = r.viewport.Help()
	} else {
		binds := r.components[r.focused].Help()
		reqViewBinds = []key.Binding{
			keybinds.Keys.Prev,
			keybinds.Keys.Next,
			keybinds.Keys.Save,
			keybinds.Keys.SendRequest,
		}
		reqViewBinds = append(binds, reqViewBinds...)
	}
	// FIX: should append responsePage binds when necessary
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
		w := r.components[urlInput].GetWidth()
		r.components[urlInput].SetWidth(r.width - (w + 30))
		r.viewport, cmd = r.viewport.Update(tea.WindowSizeMsg{
			Height: msg.Height,
			Width:  r.width - 4,
		})
	case messages.Response:
		if msg.Err != nil {
			// TODO: do something here idek
		}
		r.viewport.SetState(msg.Data)
		r.loading = false
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, keybinds.Keys.Back):
			if r.responsePage {
				r.responsePage = false
				return r, nil
			}
			return r, func() tea.Msg {
				return messages.NavigateToView{
					ViewName: Endpoints,
					Target:   MainModel,
				}
			}
		case key.Matches(msg, keybinds.Keys.Next):
			r.shift(true)
			return r, func() tea.Msg {
				return messages.RefreshItemsList{}
			}
		case key.Matches(msg, keybinds.Keys.Prev):
			r.shift(false)
			return r, func() tea.Msg {
				return messages.RefreshItemsList{}
			}
		case key.Matches(msg, keybinds.Keys.SendRequest):
			request := &http.Request{
				Method: r.endpoint.Method,
				URL:    r.endpoint.Url,
			}
			r.responsePage = true
			r.loading = true
			sendMsg := func() tea.Msg {
				res, err := r.client.ExecuteRequest(request)
				if err != nil {
					return messages.Response{
						Err: err,
					}
				}
				return messages.Response{
					Data: res,
				}
			}
			return r, tea.Batch(r.spinner.Tick, sendMsg)
		case key.Matches(msg, keybinds.Keys.Save):
			for _, val := range r.components {
				r.endpoint = val.UpdateState(r.endpoint)
			}
			r.update(
				context.Background(),
				r.endpoint.GetID(),
				endpoints.EndpointData{
					Name:   r.endpoint.Name,
					Method: r.endpoint.Method,
					URL:    r.endpoint.Url,
				})
		}
	}

	if !r.responsePage {
		r.components[r.focused], cmd = r.components[r.focused].Update(msg)
	} else {
		r.viewport, cmd = r.viewport.Update(msg)
	}

	cmds = append(cmds, cmd)

	if r.loading {
		r.spinner, cmd = r.spinner.Update(msg)
		cmds = append(cmds, cmd)
	}

	return r, tea.Batch(cmds...)
}

func (r *RequestView) shift(next bool) {
	if next {
		r.index = (r.index + 1) % len(componentList)
	} else {
		r.index = (r.index - 1 + len(componentList)) % len(componentList)
	}
	r.endpoint = r.components[r.focused].UpdateState(r.endpoint)
	r.components[r.focused].OnBlur()
	r.focused = componentList[r.index]
	r.components[r.focused].OnFocus()
	r.update(
		context.Background(),
		r.endpoint.GetID(),
		endpoints.EndpointData{
			Name:   r.endpoint.Name,
			Method: r.endpoint.Method,
			URL:    r.endpoint.Url,
		})
}

func (r *RequestView) View() string {
	if r.responsePage {
		if r.loading {
			loadingString := fmt.Sprintf("%s Loading...", r.spinner.View())
			return styles.RequestLayout(r.height, r.width)(lipgloss.Place(r.width, r.height-1, lipgloss.Center, lipgloss.Center, loadingString))
		}
		view := styles.ResponseStyle(r.height, r.width)(r.viewport.View())
		return view
	} else {
		if r.endpoint.Name == "" {
			return styles.RequestLayout(r.height, r.width)("No Endpoint selected")
		}
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
}

func (r *RequestView) SetState(items ...any) error {
	if len(items) == 1 {
		if data, ok := items[0].(EndpointData); ok {
			ep, err := r.epManager.Read(context.Background(), data.EndpointID)
			r.collection = data.Collection
			if err != nil {
				// FIX: do something over here
			}
			r.endpoint = ep
			for _, val := range componentList {
				r.components[val].SetState(ep)
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
	s := spinner.New()
	s.Spinner = spinner.Dot

	return &RequestView{
		components: map[reqFocused]componenttypes.ReqViewComponent{
			methodPicker: methodpicker.NewMethodPicker(mpConfig),
			urlInput:     urlinput.NewUrlInput(uiConfig),
		},

		epManager:    cfg.EpManager,
		focused:      methodPicker,
		viewport:     viewport.NewViewport(),
		order:        cfg.Order,
		update:       cfg.Update,
		client:       cfg.Client,
		responsePage: false,
		spinner:      s,
	}
}
