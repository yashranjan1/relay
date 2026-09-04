package toast

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type ToastFeed struct {
	width  int
	height int
	toasts []Toast
}

func (t *ToastFeed) Init() tea.Cmd {
	return nil
}

func (t *ToastFeed) AddToast(toastType ToastType, msg string) {
	t.toasts = append(t.toasts, Toast{
		Type:    toastType,
		Message: msg,
	})
}

func (t *ToastFeed) Update(msg tea.Msg) (*ToastFeed, tea.Cmd) {
	var cmd tea.Cmd

	return t, cmd
}

func (t *ToastFeed) View() string {
	s := ""
	for _, toast := range t.toasts {
		s = lipgloss.JoinVertical(lipgloss.Bottom, toast.Render())
	}
	return s
}

func (t *ToastFeed) OnFocus() {
}

func (t *ToastFeed) OnBlur() {
}

func NewToastFeed() *ToastFeed {
	return &ToastFeed{
		toasts: []Toast{
			{
				Type:    Error,
				Message: "Your balls are gone",
			},
		},
	}
}
