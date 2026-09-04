package urlinput

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/yashranjan1/relay/internal/backend/endpoints"
	componenttypes "github.com/yashranjan1/relay/internal/tui/components/ComponentTypes"
)

type UrlInput struct {
	text        textinput.Model
	styleGenner func(bool) lipgloss.Style
	width       int
	height      int
	active      bool
}

func (u *UrlInput) Init() tea.Cmd {
	return nil
}

func (u *UrlInput) SetWidth(width int) {
	u.text.SetWidth(width)
}

func (u *UrlInput) GetWidth() int {
	return u.text.Width()
}

func (u *UrlInput) Help() []key.Binding {
	return []key.Binding{}
}

func (u *UrlInput) SetState(ep endpoints.EndpointEntity) {
	u.text.SetValue(ep.Url)
}

func (u *UrlInput) UpdateState(ep endpoints.EndpointEntity) endpoints.EndpointEntity {
	ep.Url = u.text.Value()
	return ep
}

func (u *UrlInput) Update(msg tea.Msg) (componenttypes.ReqViewComponent, tea.Cmd) {
	var cmd tea.Cmd

	u.text, cmd = u.text.Update(msg)

	return u, cmd
}

func (u *UrlInput) View() string {
	return u.styleGenner(u.active).Render(u.text.View())
}

func (u *UrlInput) OnFocus() {
	u.text.Focus()
	u.active = true
}

func (u *UrlInput) OnBlur() {
	u.text.Blur()
	u.active = false
}

func NewUrlInput(config UrlInputConfig) *UrlInput {
	return &UrlInput{
		text:        textinput.New(),
		width:       config.Width,
		height:      config.Height,
		styleGenner: config.StyleGenner,
		active:      false,
	}
}
