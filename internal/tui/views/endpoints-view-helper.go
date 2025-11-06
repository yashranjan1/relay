package views

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maniac-en/req/internal/tui/components/OptionsProvider"
)

func createEndpointsDelegate() list.ItemDelegate {
	d := epItemDelegate{}

	// d.Styles.SelectedTitle = styles.SelectedListStyle
	// d.Styles.SelectedDesc = styles.SelectedListStyle

	return d
}

type epItemDelegate struct{}

func (e epItemDelegate) Height() int                             { return 1 }
func (e epItemDelegate) Spacing() int                            { return 1 }
func (e epItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (e epItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(optionsProvider.Option)

	if !ok {
		return
	}

	str := fmt.Sprintf("%s %s", item.Subtext, item.Name)

	fn := lipgloss.NewStyle().Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return lipgloss.NewStyle().Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
