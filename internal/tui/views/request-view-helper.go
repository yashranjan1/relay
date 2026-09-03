package views

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	methodpicker "github.com/yashranjan1/relay/internal/tui/components/MethodPicker"
	urlinput "github.com/yashranjan1/relay/internal/tui/components/UrlInput"
	"github.com/yashranjan1/relay/internal/tui/keybinds"
	"github.com/yashranjan1/relay/internal/tui/styles"
)

func createRequestDelegate(active bool) list.ItemDelegate {
	d := requestDelegate{}

	if active {
		d.SelectedTitleStyle = styles.ActiveRequestItem
	} else {
		d.SelectedTitleStyle = styles.InactiveRequestItem
	}
	d.UnselectedTitleStyle = styles.InactiveRequestItem

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

func createMethodPickerConfig() methodpicker.MethodPickerConfig[string] {
	methods := []string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
	}

	keys := keybinds.NewListKeyMap()

	return methodpicker.MethodPickerConfig[string]{
		Items:            methods,
		ItemMapper:       methodItemMapper,
		FilteringEnabled: false,
		ShowPagination:   false,
		ShowStatusBar:    false,
		KeyMap:           keys,
		Delegate:         createRequestDelegate,
		ShowHelp:         false,
		ShowTitle:        false,
		Width:            30,
		Height:           1,
	}
}

func createURLInputConfig() urlinput.UrlInputConfig {
	return urlinput.UrlInputConfig{
		Width:       50,
		Height:      1,
		StyleGenner: styles.UrlInputStyle,
	}
}
