package componenttypes

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maniac-en/req/internal/backend/endpoints"
)

type ReqViewComponent interface {
	Init() tea.Cmd
	Update(tea.Msg) (ReqViewComponent, tea.Cmd)
	View() string
	SetWidth(int)
	UpdateState(endpoints.EndpointEntity) endpoints.EndpointEntity
	SetState(endpoints.EndpointEntity)
	GetWidth() int
	Help() []key.Binding
	OnFocus()
	OnBlur()
}
