package toast

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/yashranjan1/relay/internal/tui/messages"
)

type ToastFeed struct {
	width   int
	height  int
	toasts  map[int]Toast
	counter int
}

func (t *ToastFeed) Init() tea.Cmd {
	return nil
}

func (t *ToastFeed) AddToast(toastType messages.ToastType, msg string) tea.Cmd {
	t.counter++

	id := t.counter

	t.toasts[id] = Toast{
		Type:    toastType,
		Message: msg,
	}

	return func() tea.Msg {
		time.Sleep(5 * time.Second)
		return messages.RemoveToast{
			ID: id,
		}
	}
}

func (t *ToastFeed) RemoveToast(id int) {
	delete(t.toasts, id)
}

func (t *ToastFeed) Update(msg tea.Msg) (*ToastFeed, tea.Cmd) {
	var cmd tea.Cmd

	return t, cmd
}

func (t *ToastFeed) View() string {
	s := ""
	for _, toast := range t.toasts {
		s = lipgloss.JoinVertical(lipgloss.Bottom, s, toast.Render())
	}
	return s
}

func (t *ToastFeed) GetWidth() int {
	return TOAST_WIDTH
}

func (t *ToastFeed) GetHeight() int {
	return TOAST_HEIGHT * len(t.toasts)
}

func (t *ToastFeed) OnFocus() {
}

func (t *ToastFeed) OnBlur() {
}

func NewToastFeed() *ToastFeed {
	return &ToastFeed{
		counter: 0,
		toasts:  map[int]Toast{},
	}
}
