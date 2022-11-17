package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func InputModel() tea.Model {
	p := tea.NewProgram(initialModel())

	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	return m
}

type (
	errMsg error
)

type InputStruct struct {
	textInput textinput.Model
	err       error
}

func initialModel() InputStruct {
	ti := textinput.New()
	ti.Placeholder = "disploy"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return InputStruct{
		textInput: ti,
		err:       nil,
	}
}

func (m InputStruct) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputStruct) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m InputStruct) View() string {
	return fmt.Sprintf(
		"What is the name of your project? (esc to quit)\n%s\n",
		m.textInput.View(),
	) + "\n"
}
