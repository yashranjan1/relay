package toast

import (
	"github.com/yashranjan1/relay/internal/tui/messages"
	"github.com/yashranjan1/relay/internal/tui/styles"
)

const TOAST_WIDTH = 30
const TOAST_HEIGHT = 2

type Toast struct {
	Message string
	Type    messages.ToastType
}

func (t Toast) Render() string {
	switch t.Type {
	case messages.Error:
		return styles.ErrorToast(t.Message, TOAST_WIDTH, TOAST_HEIGHT)
	case messages.Warning:
		return styles.WarnToast(t.Message, TOAST_WIDTH, TOAST_HEIGHT)
	case messages.Info:
		return styles.InfoToast(t.Message, TOAST_WIDTH, TOAST_HEIGHT)
	default:
		return styles.InfoToast(t.Message, TOAST_WIDTH, TOAST_HEIGHT)
	}
}
