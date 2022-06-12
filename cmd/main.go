package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	ps "github.com/cynicade/systemgo/pkg/process_selector"
)

type model struct {
	choices         []string
	cursor          int
	unitName        string
	unitDescription string
}

func initialModel(unit ps.Unit) model {
	return model{
		choices:         []string{"Restart", "Stop", "Disable"},
		unitName:        unit.Name,
		unitDescription: unit.Description,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			s := m.choices[m.cursor]
			fmt.Printf("%v\n", s)
		}
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("%s:\n%s\n\n", m.unitName, m.unitDescription)

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel(ps.UnitSelector()))

	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
