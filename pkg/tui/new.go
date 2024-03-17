package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

const (
	Title = iota
	Date
	Priority
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type NewModel struct {
	Inputs  []textinput.Model
	focused int
	err     error
}


func dateValidator(s string) error {
	// The 3 character should be a slash (/)
	// The rest should be numbers
	e := strings.ReplaceAll(s, "/", "")
	_, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		return fmt.Errorf("EXP is invalid")
	}

	if (len(s) >= 3 &&  s[2] != '/' ) {
		return fmt.Errorf("date is invalid")
	}
	if  len(s) >= 6 && s[5] != '/' {
		return fmt.Errorf("date is invalid")
	}

	return nil
}

func priorityValidator(s string) error {
	priority, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	} else if  priority <= 0 || priority > 5 {
		return fmt.Errorf("EXP is invalid")
	}
	return nil
}

func InitialNewModel() NewModel {
	var inputs []textinput.Model = make([]textinput.Model, 3)
	inputs[Title] = textinput.New()
	inputs[Title].Placeholder = "New title"
	inputs[Title].Focus()
	inputs[Title].CharLimit = 20
	inputs[Title].Width = 30
	inputs[Title].Prompt = ""

	inputs[Date] = textinput.New()
	inputs[Date].Placeholder = "MM/DD/YYYY "
	inputs[Date].CharLimit = 10
	inputs[Date].Width = 10
	inputs[Date].Prompt = ""
	inputs[Date].Validate = dateValidator

	inputs[Priority] = textinput.New()
	inputs[Priority].Placeholder = "priority"
	inputs[Priority].CharLimit = 1
	inputs[Priority].Width = 7
	inputs[Priority].Prompt = ""
	inputs[Priority].Validate = priorityValidator

	return NewModel{
		Inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m NewModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m NewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.Inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.Inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.Inputs {
			m.Inputs[i].Blur()
		}
		m.Inputs[m.focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.Inputs {
		m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m NewModel) View() string {
	return fmt.Sprintf(` %s
 %s

 %s  %s
 %s  %s

 %s
`,
		inputStyle.Width(30).Render("Task title"),
		m.Inputs[Title].View(),
		inputStyle.Width(9).Render("Due date"),
		inputStyle.Width(8).Render("Priority"),
		m.Inputs[Date].View(),
		m.Inputs[Priority].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

// nextInput focuses the next input field
func (m *NewModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.Inputs)
}

// prevInput focuses the previous input field
func (m *NewModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.Inputs) - 1
	}
}
