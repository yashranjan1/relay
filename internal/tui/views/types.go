package views

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	optionsProvider "github.com/maniac-en/req/internal/tui/components/OptionsProvider"
)

type ViewInterface interface {
	Init() tea.Cmd
	Name() string
	Help() []key.Binding
	GetFooterSegment() string
	Update(tea.Msg) (ViewInterface, tea.Cmd)
	View() string
	Order() int
	SetState(...any) error
	OnFocus() tea.Cmd
	OnBlur()
}

type EndpointData struct {
	EndpointID int64
	Collection optionsProvider.Option
}
