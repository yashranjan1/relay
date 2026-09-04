package toast

import "github.com/yashranjan1/relay/internal/tui/styles"

type ToastType int

const (
	Error ToastType = iota
	Info
	Warning
)

type Toast struct {
	Message string
	Type    ToastType
}

func (t Toast) Render() string {
	switch t.Type {
	case Error:
		return styles.ErrorToast(t.Message)
	case Warning:
		return styles.WarnToast(t.Message)
	case Info:
		return styles.InfoToast(t.Message)
	default:
		return styles.InfoToast(t.Message)
	}
}
