package views

import (
	"context"
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maniac-en/req/internal/backend/endpoints"
	"github.com/maniac-en/req/internal/backend/http"
	"github.com/maniac-en/req/internal/log"
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
	width        int
	focused      reqFocused
	components   map[reqFocused]componenttypes.ReqViewComponent
	index        int
	height       int
	loading      bool
	help         help.Model
	keys         *keybinds.ListKeyMap
	spinner      spinner.Model
	responsePage bool
	client       *http.HTTPManager
	order        int
	res          *http.Response
	update       func(context.Context, int64, endpoints.EndpointData) (endpoints.EndpointEntity, error)
	endpoint     endpoints.EndpointEntity
}

func (r *RequestView) Init() tea.Cmd {
	return nil
}

func (r *RequestView) Name() string {
	return "Request"
}

func (r *RequestView) Help() []key.Binding {
	binds := r.components[r.focused].Help()
	reqViewBinds := []key.Binding{
		keybinds.Keys.Prev,
		keybinds.Keys.Next,
		keybinds.Keys.Save,
		keybinds.Keys.SendRequest,
	}
	return append(binds, reqViewBinds...)
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
		w := r.components[urlInput].GetWidth()
		r.components[urlInput].SetWidth(r.width - (w + 30))
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keybinds.Keys.Back):
			if r.responsePage {
				r.responsePage = false
				return r, nil
			} else {
				return r, func() tea.Msg {
					return messages.NavigateToView{
						ViewName: Endpoints,
						Target:   Endpoints,
					}
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
			res, err := r.client.ExecuteRequest(request)
			r.loading = false
			r.res = res
			if err != nil {
				log.Warn("error occurred while trying to send a request", err)
			}
			return r, r.spinner.Tick
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
	}

	cmds = append(cmds, cmd)

	if r.loading {
		log.Info("working")
		r.spinner, cmd = r.spinner.Update(msg)
		cmds = append(cmds, cmd)
	}

	return r, tea.Batch(cmds...)
}

func (r *RequestView) shift(next bool) {
	if next {
		r.index = (r.index + 1) % len(componentList)
	} else {
		r.index = (r.index - 1) % len(componentList)
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
		if r.res == nil {
			return styles.RequestLayout(r.height, r.width)(lipgloss.Place(r.width, r.height-1, lipgloss.Center, lipgloss.Center, "No responses so far"))
		}
		if r.loading {
			loadingString := fmt.Sprintf("%s Loading...", r.spinner.View())
			return styles.RequestLayout(r.height, r.width)(lipgloss.Place(r.width, r.height-1, lipgloss.Center, lipgloss.Center, loadingString))
		}
		// TODO: use a viewport here PLEASE => https://github.com/charmbracelet/bubbles?tab=readme-ov-file#viewport
		responseDetails := []string{
			fmt.Sprintf("Status Code: %d", r.res.StatusCode),
			fmt.Sprintf("Body: %s", r.res.Body),
		}
		for key, val := range r.res.Headers {
			for _, item := range val {
				responseDetails = append(responseDetails, fmt.Sprintf("%s: %s", key, item))
			}
		}
		view := styles.ResponseStyle(r.height, r.width)(lipgloss.JoinVertical(lipgloss.Top, responseDetails...))
		return view
	} else {
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
		if ep, ok := items[0].(endpoints.EndpointEntity); ok {
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

func NewRequestView(update func(context.Context, int64, endpoints.EndpointData) (endpoints.EndpointEntity, error)) *RequestView {
	mpConfig := createMethodPickerConfig()

	uiConfig := createURLInputConfig()
	s := spinner.New()
	s.Spinner = spinner.Dot

	return &RequestView{
		components: map[reqFocused]componenttypes.ReqViewComponent{
			methodPicker: methodpicker.NewMethodPicker(mpConfig),
			urlInput:     urlinput.NewUrlInput(uiConfig),
		},
		focused:      methodPicker,
		update:       update,
		client:       http.NewHTTPManager(),
		responsePage: false,
		spinner:      s,
	}
}
