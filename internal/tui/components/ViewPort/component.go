package viewport

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maniac-en/req/internal/backend/http"
)

type Viewport struct {
	viewport viewport.Model
	width    int
	height   int
	ready    bool
	state    string
}

func (v *Viewport) Init() tea.Cmd {
	return nil
}

func (v *Viewport) SetState(state *http.Response) {
	var str strings.Builder

	str.WriteString(fmt.Sprintf("Status code: %d\n", state.StatusCode))
	str.WriteString(fmt.Sprintf("ContentType header: %s\n", state.Headers["Content-Type"]))
	str.WriteString(fmt.Sprintf("Body: %s\n", state.Body))

	v.state = str.String()
	v.viewport.SetContent(v.state)
}

func (v Viewport) Update(msg tea.Msg) (Viewport, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		if !v.ready {
			v.viewport = viewport.New(msg.Width, msg.Height-20)
			v.ready = true
		} else {
			v.viewport.Width = msg.Width
			v.viewport.Height = msg.Height - 10
		}
	}

	v.viewport, cmd = v.viewport.Update(msg)

	cmds = append(cmds, cmd)

	return v, tea.Batch(cmds...)
}

func (v *Viewport) View() string {
	if !v.ready {
		return "Waiting.."
	}
	// FIX: this mess
	percentage := int(v.viewport.ScrollPercent() * 100)
	return lipgloss.JoinVertical(lipgloss.Top, v.viewport.View(), fmt.Sprintf("%d%%", percentage))
}

func (v *Viewport) OnFocus() {
}

func (v *Viewport) OnBlur() {
}

func (v *Viewport) Help() []key.Binding {
	// FIX: this
	return []key.Binding{}
}

func NewViewport() Viewport {
	return Viewport{
		viewport: viewport.New(1, 1),
		ready:    false,
		state:    "No response received yet :(",
	}
}
