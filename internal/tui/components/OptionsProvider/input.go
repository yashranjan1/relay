package optionsProvider

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/yashranjan1/relay/internal/tui/messages"
	"github.com/yashranjan1/relay/internal/tui/styles"
)

type OptionsInput struct {
	input  textinput.Model
	height int
	width  int
	editId int64
	keys   InputKeyMaps
}

func NewOptionsInput(config *InputConfig) OptionsInput {
	input := textinput.New()
	input.CharLimit = config.CharLimit
	input.Placeholder = config.Placeholder
	input.SetWidth(config.Width)
	input.Prompt = config.Prompt

	inputStyles := input.Styles()
	inputStyles.Focused.Text = styles.InputStyle
	inputStyles.Blurred.Text = styles.InputStyle
	input.SetStyles(inputStyles)

	return OptionsInput{
		input:  input,
		editId: -1,
		keys:   config.KeyMap,
	}
}

func (i OptionsInput) Init() tea.Cmd {
	return nil
}

func (i OptionsInput) Update(msg tea.Msg) (OptionsInput, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, i.keys.Accept):
			itemName := i.input.Value()
			i.input.SetValue("")
			if i.editId == -1 {
				return i, func() tea.Msg { return messages.ItemAdded{Item: itemName} }
			}
			return i, func() tea.Msg { return messages.ItemEdited{Item: itemName, ItemID: i.editId} }
		case key.Matches(msg, i.keys.Back):
			return i, func() tea.Msg { return messages.DeactivateView{} }
		}
	}

	i.input, cmd = i.input.Update(msg)
	cmds = append(cmds, cmd)

	return i, tea.Batch(cmds...)
}

func (i OptionsInput) View() string {
	return styles.InputStyle.Render(i.input.View())
}

func (i OptionsInput) Help() []key.Binding {
	return []key.Binding{
		i.keys.Accept,
		i.keys.Back,
	}
}

func (i *OptionsInput) SetInput(text string) {
	i.input.SetValue(text)
}

func (i *OptionsInput) OnFocus(id ...int64) {
	if len(id) > 0 {
		i.editId = id[0]
	}
	i.input.Focus()
}

func (i *OptionsInput) OnBlur() {
	i.editId = -1
	i.input.Blur()
}
