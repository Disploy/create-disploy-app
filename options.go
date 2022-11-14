package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)


var Choices = []string{}

type OptionStruct struct {
	cursor int
	choice string
}

func (m OptionStruct) Init() tea.Cmd {
	return nil
}

func (m OptionStruct) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = Choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(Choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(Choices) - 1
			}
		}
	}

	return m, nil
}

func (m OptionStruct) View() string {
	s := strings.Builder{}
	s.WriteString("Which template you would like to use (esc to quit)\n\n")

	for i := 0; i < len(Choices); i++ {
		if m.cursor == i {
			s.WriteString("[•] ")
		} else {
			s.WriteString("[ ] ")
		}
		s.WriteString(Choices[i])
		s.WriteString("\n")
	}

	return s.String()
}

func ChoiceModel() tea.Model {
	p := tea.NewProgram(OptionStruct{})

	m, err := p.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return m
}