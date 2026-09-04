package componenttypes

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/yashranjan1/relay/internal/backend/endpoints"
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
