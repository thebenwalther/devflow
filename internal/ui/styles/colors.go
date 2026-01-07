// Package styles provides basic styling for DevFlow
package styles

import "github.com/charmbracelet/lipgloss"

// Basic color palette
var (
	Primary   = lipgloss.Color("#00FFC8") // Vibrant cyan
	Secondary = lipgloss.Color("#B478FF") // Soft purple
	Success   = lipgloss.Color("#50FA7B") // Fresh green
	Warning   = lipgloss.Color("#FFB86C") // Warm orange
	Error     = lipgloss.Color("#FF5555") // Clear red
	Info      = lipgloss.Color("#8BE9FD") // Light blue

	Base    = lipgloss.Color("#181825") // Deep dark base
	Surface = lipgloss.Color("#24243A") // Elevated surfaces

	TextPrimary   = lipgloss.Color("#CDD6F4") // Light primary text
	TextSecondary = lipgloss.Color("#A6ADC8") // Muted secondary text
	TextTertiary  = lipgloss.Color("#7F8498") // Subtle tertiary text

	BorderNormal = lipgloss.Color("#45475A") // Normal borders
	BorderFocus  = lipgloss.Color("#7D56F4") // Focused borders
)

// Component styles
var (
	Panel = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(BorderNormal).
		Background(Surface).
		Padding(1, 2).
		Margin(0, 1)

	Header = lipgloss.NewStyle().
		Bold(true).
		Foreground(Primary).
		Background(lipgloss.Color("#121212")).
		Padding(0, 2).
		MarginBottom(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Primary)

	ActiveTab = lipgloss.NewStyle().
			Foreground(Base).
			Background(Primary).
			Bold(true).
			Padding(0, 3).
			Margin(0, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Primary)

	InactiveTab = lipgloss.NewStyle().
			Foreground(TextSecondary).
			Background(Surface).
			Padding(0, 3).
			Margin(0, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(BorderNormal)

	StatusSuccess = lipgloss.NewStyle().
			Foreground(Success).
			Bold(true)

	StatusError = lipgloss.NewStyle().
			Foreground(Error).
			Bold(true)

	Content = lipgloss.NewStyle().
		Foreground(TextPrimary).
		Background(Surface).
		Padding(1).
		MarginBottom(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF6AC1"))

	Selected = lipgloss.NewStyle().
			Foreground(Base).
			Background(Primary).
			Bold(true)

	Help = lipgloss.NewStyle().
		Foreground(TextTertiary).
		Italic(true)

	Loading = lipgloss.NewStyle().
		Foreground(Info).
		Italic(true)
)
