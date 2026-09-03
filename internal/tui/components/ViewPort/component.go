package viewport

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yashranjan1/relay/internal/backend/http"
	"github.com/yashranjan1/relay/internal/tui/keybinds"
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

	fmt.Fprintf(&str, "Status code: %d\n", state.StatusCode)
	fmt.Fprintf(&str, "ContentType header: %s\n", state.Headers["Content-Type"])
	fmt.Fprintf(&str, "Body: %s\n", state.Body)

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
			v.width = msg.Width
			v.height = msg.Height
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
	str := fmt.Sprintf(" %d%% ", int(v.viewport.ScrollPercent()*100))

	padding := v.height - v.viewport.Height - 7

	return lipgloss.JoinVertical(lipgloss.Left,
		v.viewport.View(),
		lipgloss.NewStyle().
			Width(v.width).
			PaddingTop(padding).
			Align(lipgloss.Right).
			Render(str),
	)
}

func (v *Viewport) OnFocus() {
}

func (v *Viewport) OnBlur() {
}

func (v *Viewport) Help() []key.Binding {
	// FIX: this
	return []key.Binding{
		keybinds.Keys.Down,
		keybinds.Keys.Up,
	}
}

func NewViewport() Viewport {
	return Viewport{
		viewport: viewport.New(1, 1),
		ready:    false,
		state:    "No response received yet :(",
	}
}
