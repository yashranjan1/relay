package messages

import "github.com/maniac-en/req/internal/backend/http"

type ItemAdded struct {
	Item string
}
type ItemEdited struct {
	Item   string
	ItemID int64
}

type DeleteItem struct {
	ItemID int64
}

type ChooseItem[T any] struct {
	Item   T
	Source string
}

type DeactivateView struct{}

type NavigateToView struct {
	ViewName string
	Target   string
	Data     interface{}
}

type Response struct {
	Data *http.Response
	Err  error
}

type RefreshItemsList struct{}

type ShowError struct {
	Message string
}
