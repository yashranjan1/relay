package viewport

import (
	"fmt"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/yashranjan1/relay/internal/backend/http"
	componenttypes "github.com/yashranjan1/relay/internal/tui/components/ComponentTypes"
	"github.com/yashranjan1/relay/internal/tui/keybinds"
	"github.com/yashranjan1/relay/internal/tui/styles"
)

type Viewport struct {
	viewport viewport.Model
	width    int
	height   int
	ready    bool
	focused  bool
	status   int
	headers  map[string][]string
	body     string
}

func (v *Viewport) Init() tea.Cmd {
	return nil
}

func (v *Viewport) EraseState() {
	v.status = 0
	v.body = ""
	v.headers = nil
}

func (v *Viewport) SetWidth(width int) {
	v.width = width
}

func (v *Viewport) GetWidth() int {
	return v.width
}

func (v *Viewport) SetState(state *http.Response) {
	v.body = state.Body
	v.status = state.StatusCode
	v.headers = state.Headers
	v.viewport.SetContent(v.body)
}

func (v *Viewport) Update(msg tea.Msg) (componenttypes.FocusableComponent, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:

		v.width = msg.Width

		footerHeight := 1
		vpHeight := max(1, msg.Height-footerHeight)
		v.height = vpHeight

		if !v.ready {
			v.viewport = viewport.New(viewport.WithWidth(msg.Width), viewport.WithHeight(vpHeight))
			v.ready = true
		} else {
			v.viewport.SetWidth(msg.Width)
			v.viewport.SetHeight(vpHeight)
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

	hasResponse := !(v.status == 0 && v.body == "" && v.headers == nil)

	contentHeight := v.height
	if hasResponse {
		contentHeight = max(1, v.height-1)
	}

	v.viewport.SetHeight(contentHeight)
	contentView := styles.ResponseContentStyle(contentHeight, v.width)(v.viewport.View())

	if !hasResponse {
		return styles.ResponseStyle(v.focused)(contentView)
	}

	scrollPercent := fmt.Sprintf(" %d%% ", int(v.viewport.ScrollPercent()*100))
	status := fmt.Sprintf("  Status: %s", styles.ResponseStatus(v.status))
	statusWidth := lipgloss.Width(status)
	scrollPercentView := styles.ResponseFooterStyle(v.width - statusWidth)(scrollPercent)

	footer := lipgloss.JoinHorizontal(
		lipgloss.Right,
		status,
		scrollPercentView,
	)

	view := lipgloss.JoinVertical(lipgloss.Left,
		contentView,
		footer,
	)

	return styles.ResponseStyle(v.focused)(view)
}

func (v *Viewport) OnFocus() {
	v.focused = true
}

func (v *Viewport) OnBlur() {
	v.focused = false
}

func (v *Viewport) Help() []key.Binding {
	return []key.Binding{
		keybinds.Keys.Down,
		keybinds.Keys.Up,
		keybinds.Keys.PageDown,
		keybinds.Keys.PageUp,
	}
}

func NewViewport() *Viewport {
	return &Viewport{
		viewport: viewport.New(viewport.WithWidth(1), viewport.WithHeight(1)),
		ready:    false,
		body:     "No response received yet :(",
		focused:  false,
	}
}
