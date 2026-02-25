package views

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	methodpicker "github.com/maniac-en/req/internal/tui/components/MethodPicker"
	"github.com/maniac-en/req/internal/tui/styles"
)

func createRequestDelegate(active bool) list.ItemDelegate {
	d := requestDelegate{}

	if active {
		d.SelectedTitleStyle = styles.SelectedActiveListStyle
	} else {
		d.SelectedTitleStyle = styles.SelectedInactiveListStyle
	}
	d.UnselectedTitleStyle = styles.UnselectedListStyle

	return d
}

type requestDelegate struct {
	SelectedTitleStyle   lipgloss.Style
	UnselectedTitleStyle lipgloss.Style
}

func (r requestDelegate) Height() int                             { return 1 }
func (r requestDelegate) Spacing() int                            { return 1 }
func (r requestDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (r requestDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(methodpicker.MethodOption)

	if !ok {
		return
	}

	str := fmt.Sprintf("< %s >", item.Name)

	fn := r.UnselectedTitleStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return r.SelectedTitleStyle.Render(strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
