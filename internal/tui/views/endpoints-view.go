package views

import (
	"context"
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maniac-en/req/internal/backend/database"
	"github.com/maniac-en/req/internal/backend/endpoints"
	optionsProvider "github.com/maniac-en/req/internal/tui/components/OptionsProvider"
	"github.com/maniac-en/req/internal/tui/keybinds"
	"github.com/maniac-en/req/internal/tui/messages"
)

type epFocused string

const (
	listView    = "listview"
	requestView = "requestview"
)

type EndpointsView struct {
	height      int
	focused     epFocused
	collection  optionsProvider.Option
	width       int
	order       int
	list        optionsProvider.OptionsProvider[endpoints.EndpointEntity, database.Endpoint]
	requestView ViewInterface
	manager     *endpoints.EndpointsManager
}

func (e *EndpointsView) Init() tea.Cmd {
	return nil
}

func (e *EndpointsView) Name() string {
	return "Endpoints"
}

func (e *EndpointsView) Help() []key.Binding {
	if e.focused == listView {
		return e.list.Help()
	} else {
		return e.requestView.Help()
	}
}

func (e *EndpointsView) GetFooterSegment() string {
	return fmt.Sprintf("%s/%s", e.collection.Name, e.list.GetSelected().Name)
}

func (e *EndpointsView) Update(msg tea.Msg) (ViewInterface, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		e.height = msg.Height
		e.width = msg.Width
		e.list, cmd = e.list.Update(msg)
		_, w := e.list.GetSize()
		e.requestView, cmd = e.requestView.Update(tea.WindowSizeMsg{
			Height: msg.Height,
			Width:  msg.Width - w,
		})
		cmds = append(cmds, cmd)
	case messages.ItemAdded:
		ep, err := e.manager.CreateEndpoint(context.Background(), endpoints.EndpointData{
			CollectionID: e.collection.ID,
			Name:         msg.Item,
			Method:       "GET",
		})
		if err != nil {
			//TODO: handle this
		}
		e.requestView.SetState(ep)
	case messages.RefreshItemsList:
		e.list.RefreshItems()
	case messages.ItemEdited:
		e.manager.UpdateEndpointName(context.Background(), msg.ItemID, msg.Item)
	case messages.DeleteItem:
		e.manager.Delete(context.Background(), msg.ItemID)
		e.list.RefreshItems()
	case messages.ChooseItem[optionsProvider.Option]:
		e.list.OnBlur()
		e.requestView.OnFocus()
		e.focused = requestView
	case messages.NavigateToView:
		if msg.Target != Endpoints {
			break
		}
		e.requestView.OnBlur()
		e.focused = listView
		e.list.OnFocus()

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keybinds.Keys.Back):
			if e.focused == listView {
				return e, func() tea.Msg {
					return messages.NavigateToView{
						ViewName: Collections,
						Target:   MainModel,
					}
				}
			}
		}
	}

	switch e.focused {
	case listView:
		e.list, cmd = e.list.Update(msg)
	case requestView:
		e.requestView, cmd = e.requestView.Update(msg)
	}

	cmds = append(cmds, cmd)

	return e, tea.Batch(cmds...)
}

func (e *EndpointsView) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, e.list.View(), e.requestView.View())
}

func (e *EndpointsView) OnFocus() {

}

func (e *EndpointsView) SetState(items ...any) error {
	// TODO: Something over here should also set the state for the request view
	if len(items) == 1 {
		if collection, ok := items[0].(optionsProvider.Option); ok {
			e.collection = collection
			epListFunc := func(ctx context.Context) ([]endpoints.EndpointEntity, error) {
				return e.manager.ListByCollection(ctx, collection.ID)
			}
			e.list.SetGetItemsFunc(epListFunc)
			return nil
		}
	}
	return errors.New("Invalid inputs, this function takes 1 input of type optionsProvider.Options")
}

func (e *EndpointsView) OnBlur() {

}

func (e *EndpointsView) Order() int {
	return e.order
}

func itemMapperEp(items []endpoints.EndpointEntity) []list.Item {
	opts := make([]list.Item, len(items))
	for i, item := range items {
		newOpt := optionsProvider.Option{
			Name:    item.GetName(),
			Subtext: item.Method,
			ID:      item.GetID(),
		}
		opts[i] = newOpt
	}
	return opts
}

func NewEndpointsView(epManager *endpoints.EndpointsManager, order int) *EndpointsView {
	view := &EndpointsView{
		order: order,
		collection: optionsProvider.Option{
			Name:    "",
			Subtext: "",
			ID:      0,
		},
		manager: epManager,
	}

	keybinds := keybinds.NewListKeyMap()
	config := defaultListConfig[endpoints.EndpointEntity, database.Endpoint](keybinds, createEndpointsDelegate)

	epUpdateFunc := epManager.UpdateEndpoint
	view.requestView = NewRequestView(epUpdateFunc)

	config.OnChangeAction = func(id int64) {
		epDetails, err := epManager.Read(context.Background(), id)
		if err != nil {
			// TODO: what do we do here? prob throw a crash ig?
		}
		view.requestView.SetState(epDetails)
	}

	epListFunc := func(ctx context.Context) ([]endpoints.EndpointEntity, error) {
		return epManager.ListByCollection(ctx, view.collection.ID)
	}

	config.GetItemsFunc = epListFunc
	config.ItemMapper = itemMapperEp
	config.AdditionalKeymaps = keybinds
	config.Source = "endpoints"
	config.Placeholder = "Add a new endpoint..."

	view.list = optionsProvider.NewOptionsProvider(config)
	view.focused = listView

	return view
}
