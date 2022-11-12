package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/libgit2/git2go/v28"
)

var URL = "https://github.com/disploy/create-disploy-app"

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			project := m.table.SelectedRow()[1];

			copy(project)

			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func copy(name string) {
	println("Moving " + name);
	os.Rename(".disploy/assets/" + name, name);
	os.RemoveAll(".disploy");
}

func main() {
	git.Clone(URL, ".disploy", &git.CloneOptions{})

	files, err := ioutil.ReadDir(".disploy/assets")

	if err != nil {
		println(err)
	}

	columns := []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Name", Width: 24},
	}

	rows := []table.Row{}

	for i, f := range files {
		rows = append(rows, table.Row{fmt.Sprint(i), f.Name()})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		println("Error running program:", err)
	}
}
