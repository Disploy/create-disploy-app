package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var PackageManagers = []string{
	"npm",
	"yarn",
	"pnpm",
}

type PackageManagerOptionStruct struct {
	cursor int
	choice string
}

func (m PackageManagerOptionStruct) Init() tea.Cmd {
	return nil
}

func (m PackageManagerOptionStruct) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "enter":
			m.choice = PackageManagers[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(PackageManagers) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(PackageManagers) - 1
			}
		}
	}

	return m, nil
}

func (m PackageManagerOptionStruct) View() string {
	s := strings.Builder{}
	s.WriteString("Which package manager you would like to use (esc to quit)\n\n")

	for i := 0; i < len(PackageManagers); i++ {
		if m.cursor == i {
			s.WriteString("[•] ")
		} else {
			s.WriteString("[ ] ")
		}
		s.WriteString(PackageManagers[i])
		s.WriteString("\n")
	}

	return s.String()
}

func PackageManagerChoiceModel() tea.Model {
	p := tea.NewProgram(PackageManagerOptionStruct{})

	m, err := p.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return m
}
