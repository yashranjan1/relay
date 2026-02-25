package urlinput

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	componenttypes "github.com/maniac-en/req/internal/tui/components/ComponentTypes"
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

func (u *UrlInput) SetSize(width, height int) {
	u.text.PromptStyle.Height(height)
	u.text.PromptStyle.Width(width)
}

func (u *UrlInput) Help() []key.Binding {
	return []key.Binding{}
}

func (u *UrlInput) Update(msg tea.Msg) (componenttypes.ComponentInterface, tea.Cmd) {
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
