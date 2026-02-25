package views

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maniac-en/req/internal/tui/components/OptionsProvider"
	"github.com/maniac-en/req/internal/tui/styles"
)

func createEndpointsDelegate(active bool) list.ItemDelegate {
	d := epItemDelegate{}

	if active {
		d.SelectedTitleStyle = styles.SelectedActiveListStyle
	} else {
		d.SelectedTitleStyle = styles.SelectedInactiveListStyle
	}
	d.UnselectedTitleStyle = styles.UnselectedListStyle

	return d
}

type epItemDelegate struct {
	SelectedTitleStyle   lipgloss.Style
	UnselectedTitleStyle lipgloss.Style
}

func (e epItemDelegate) Height() int                             { return 1 }
func (e epItemDelegate) Spacing() int                            { return 1 }
func (e epItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (e epItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(optionsProvider.Option)

	if !ok {
		return
	}

	str := fmt.Sprintf("%s %s", item.Subtext, item.Name)

	fn := e.UnselectedTitleStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return e.SelectedTitleStyle.Render(strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
