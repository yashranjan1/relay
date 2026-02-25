package componenttypes

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type ComponentInterface interface {
	Init() tea.Cmd
	Update(tea.Msg) (ComponentInterface, tea.Cmd)
	View() string
	SetSize(int, int)
	Help() []key.Binding
	OnFocus()
	OnBlur()
}
