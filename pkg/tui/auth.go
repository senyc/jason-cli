package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AuthModel struct {
	Input   textinput.Model
	Err     error
	focused int
}

func InitiaAuthModel() AuthModel {
	var input textinput.Model = textinput.Model{}
	input = textinput.New()
	input.Placeholder = "Api key"
	input.Focus()
	input.CharLimit = 20
	input.Width = 30
	input.Prompt = ""

	return AuthModel{
		Input:   input,
		Err:     nil,
		focused: 0,
	}
}

func (m AuthModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	// We handle errors just like any other message
	case errMsg:
		m.Err = msg
		return m, nil

	}
	m.Input, cmd = m.Input.Update(msg)
	return m, cmd
}

func (m AuthModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AuthModel) View() string {
	return fmt.Sprintf(` %s
 %s
 %s  
`,
		inputStyle.Width(30).Render("Api key"),
		m.Input.View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}
