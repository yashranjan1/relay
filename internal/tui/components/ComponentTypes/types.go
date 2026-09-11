package componenttypes

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/yashranjan1/relay/internal/backend/endpoints"
	"github.com/yashranjan1/relay/internal/backend/http"
)

type FocusableComponent interface {
	Init() tea.Cmd
	Update(tea.Msg) (FocusableComponent, tea.Cmd)
	View() string
	SetWidth(int)
	GetWidth() int
	Help() []key.Binding
	OnFocus()
	OnBlur()
}

type EndpointBindable interface {
	FocusableComponent
	SetState(endpoints.EndpointEntity)
	UpdateState(endpoints.EndpointEntity) endpoints.EndpointEntity
}

type ResponseSettable interface {
	SetState(*http.Response)
	EraseState()
}
