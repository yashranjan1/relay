package app

import (
	"sort"
	"strings"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/yashranjan1/relay/internal/log"
	toast "github.com/yashranjan1/relay/internal/tui/components/Toast"
	"github.com/yashranjan1/relay/internal/tui/keybinds"
	"github.com/yashranjan1/relay/internal/tui/messages"
	"github.com/yashranjan1/relay/internal/tui/styles"
	"github.com/yashranjan1/relay/internal/tui/views"
)

type ViewName string

const (
	Collections ViewName = "collections"
	Endpoints   ViewName = "endpoints"
	Request     ViewName = "request"
)

type Heading struct {
	name  string
	order int
}

type AppModel struct {
	ctx         *Context
	width       int
	height      int
	Views       map[ViewName]views.ViewInterface
	toast       *toast.ToastFeed
	focusedView ViewName
	keys        []key.Binding
	help        help.Model
	errorMsg    string
}

func (a AppModel) Init() tea.Cmd {
	return nil
}

func (a AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case messages.DeleteItem:
		a.Views[Collections], cmd = a.Views[Collections].Update(messages.RefreshItemsList{})
		cmds = append(cmds, cmd)
	case messages.ItemAdded:
		a.Views[Collections], cmd = a.Views[Collections].Update(messages.RefreshItemsList{})
		cmds = append(cmds, cmd)
	case messages.RemoveToast:
		a.toast.RemoveToast(msg.ID)
	case messages.AddToast:
		cmd = a.toast.AddToast(msg.Type, msg.Message)
		cmds = append(cmds, cmd)
	case tea.WindowSizeMsg:
		a.height = msg.Height
		a.width = msg.Width
		for key := range a.Views {
			a.Views[key], cmd = a.Views[key].Update(tea.WindowSizeMsg{Height: a.AvailableHeight(), Width: msg.Width})
			cmds = append(cmds, cmd)
		}
		return a, tea.Batch(cmds...)
	case messages.NavigateToView:
		if msg.Data != nil {
			err := a.Views[ViewName(msg.ViewName)].SetState(msg.Data)
			if err != nil {
				log.Error("failed to set view state during navigation", "target_view", msg.ViewName, "error", err)
				return a, nil
			}
		} else if msg.Target != views.MainModel {
			break
		}

		a.Views[a.focusedView].OnBlur()

		a.focusedView = ViewName(msg.ViewName)
		cmd = a.Views[a.focusedView].OnFocus()
		cmds = append(cmds, cmd)

		return a, tea.Batch(cmds...)

	case messages.ShowError:
		log.Error("user operation failed", "error", msg.Message)
		a.errorMsg = msg.Message
		return a, nil
	case tea.KeyPressMsg:
		a.errorMsg = ""
		switch {
		case key.Matches(msg, keybinds.Keys.Quit):
			return a, tea.Quit
		}
	}

	a.Views[a.focusedView], cmd = a.Views[a.focusedView].Update(msg)
	cmds = append(cmds, cmd)

	return a, tea.Batch(cmds...)
}

func (a AppModel) View() tea.View {
	footer := a.Footer()
	header := a.Header()
	view := a.Views[a.focusedView].View()
	help := a.Help()

	appView := lipgloss.NewLayer(lipgloss.JoinVertical(lipgloss.Top, header, view, help, footer))

	offset := 1

	toasts := lipgloss.NewLayer(a.toast.View()).
		X(a.width - a.toast.GetWidth() - offset).
		Y(a.height - a.toast.GetHeight() - offset).
		Z(1)

	composite := lipgloss.NewCompositor(appView, toasts)

	v := tea.NewView(composite.Render())
	v.AltScreen = true
	return v

}

func (a AppModel) Help() string {
	viewHelp := a.Views[a.focusedView].Help()

	var appHelp []key.Binding
	appHelp = append(appHelp, a.keys...)

	if a.focusedView == Endpoints {
		appHelp = append(appHelp, keybinds.Keys.Back)
	}

	allHelp := append(viewHelp, appHelp...)
	helpStruct := keybinds.Help{
		Keys: allHelp,
	}
	return styles.HelpStyle.Render(a.help.View(helpStruct))
}

func (a *AppModel) AvailableHeight() int {
	footer := a.Footer()
	header := a.Header()
	help := a.Help()
	return a.height - lipgloss.Height(header) - lipgloss.Height(footer) - lipgloss.Height(help)
}

func (a AppModel) Header() string {
	var b strings.Builder

	// INFO: this might be a bit messy, could be a nice idea to look into OrderedMaps maybe?
	views := []Heading{}
	for key := range a.Views {
		views = append(views, Heading{
			name:  a.Views[key].Name(),
			order: a.Views[key].Order(),
		})
	}
	sort.Slice(views, func(i, j int) bool {
		return views[i].order < views[j].order
	})

	for _, item := range views {
		if item.name == a.Views[a.focusedView].Name() {
			b.WriteString(styles.TabHeadingActive.Render(item.name))
		} else {
			b.WriteString(styles.TabHeadingInactive.Render(item.name))
		}
	}

	b.WriteString(styles.TabHeadingInactive.Render(""))

	return b.String()
}

func (a AppModel) Footer() string {
	name := styles.ApplyGradientToFooter("REQ")
	footerText := styles.FooterSegmentStyle.Render(a.Views[a.focusedView].GetFooterSegment())
	version := styles.FooterVersionStyle.Width(a.width - lipgloss.Width(name) - lipgloss.Width(footerText)).Render(a.ctx.Version)
	return lipgloss.JoinHorizontal(lipgloss.Left, name, footerText, version)
}

func NewAppModel(ctx *Context) AppModel {
	appKeybinds := []key.Binding{
		keybinds.Keys.Quit,
	}

	toasts := toast.NewToastFeed()

	model := AppModel{
		focusedView: Collections,
		ctx:         ctx,
		help:        help.New(),
		keys:        appKeybinds,
		toast:       toasts,
	}

	epUpdateFunc := model.ctx.Endpoints.UpdateEndpoint

	reqCfg := views.RequestViewConfig{
		EpManager: model.ctx.Endpoints,
		Update:    epUpdateFunc,
		Order:     3,
		Client:    model.ctx.HTTP,
	}

	model.Views = map[ViewName]views.ViewInterface{
		Collections: views.NewCollectionsView(model.ctx.Collections, model.ctx.Endpoints, 1),
		Endpoints:   views.NewEndpointsView(model.ctx.Endpoints, 2),
		Request:     views.NewRequestView(reqCfg),
	}
	return model
}
