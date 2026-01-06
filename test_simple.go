// Simple test to check if basic Bubble Tea works
package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	content string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC || msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00FFC8")).
		Background(lipgloss.Color("#24243A")).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder())

	return style.Render("🎉 Welcome to DevFlow!\n\nPress 'q' to quit")
}

func main() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
