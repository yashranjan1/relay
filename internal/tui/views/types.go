package views

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	optionsProvider "github.com/yashranjan1/relay/internal/tui/components/OptionsProvider"
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
