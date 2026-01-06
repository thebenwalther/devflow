// Package styles provides a beautiful design system for DevFlow
package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// Color palette inspired by OpenCode's modern aesthetic
var (
	// Primary brand colors
	Primary   = lipgloss.Color("#00FFC8") // Vibrant cyan
	Secondary = lipgloss.Color("#B478FF") // Soft purple
	Accent    = lipgloss.Color("#FF6AC1") // Pink accent

	// Semantic colors
	Success = lipgloss.Color("#50FA7B") // Fresh green
	Warning = lipgloss.Color("#FFB86C") // Warm orange
	Error   = lipgloss.Color("#FF5555") // Clear red
	Info    = lipgloss.Color("#8BE9FD") // Light blue

	// Background system
	Base    = lipgloss.Color("#181825") // Deep dark base
	Surface = lipgloss.Color("#24243A") // Elevated surfaces
	Overlay = lipgloss.Color("#313244") // Overlay elements

	// Text hierarchy
	TextPrimary   = lipgloss.Color("#CDD6F4") // Light primary text
	TextSecondary = lipgloss.Color("#A6ADC8") // Muted secondary text
	TextTertiary  = lipgloss.Color("#7F8498") // Subtle tertiary text

	// Interactive states
	Focus  = lipgloss.Color("#96CDFB") // Focused elements
	Active = lipgloss.Color("#CBA6F7") // Active/selected elements
	Hover  = lipgloss.Color("#F9E2AF") // Hover state

	// Border colors
	BorderNormal = lipgloss.Color("#45475A") // Normal borders
	BorderFocus  = lipgloss.Color("#7D56F4") // Focused borders
)

// Component styles
var (
	// Base panel with rounded corners
	Panel = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(BorderNormal).
		Background(Surface).
		Padding(1, 2).
		Margin(0, 1)

	// Active/focused panel
	ActivePanel = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(BorderFocus).
			Background(lipgloss.Color("#1E1E2E")).
			BorderForeground(BorderFocus).
			Background(lipgloss.Color("#1E1E2E"))

	// Header styling
	Header = lipgloss.NewStyle().
		Bold(true).
		Foreground(Primary).
		Background(lipgloss.Color("#121212")).
		Padding(0, 2).
		MarginBottom(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Primary)

	// Tab system
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

	// Status indicators
	StatusSuccess = lipgloss.NewStyle().
			Foreground(Success).
			Bold(true)

	StatusWarning = lipgloss.NewStyle().
			Foreground(Warning).
			Bold(true)

	StatusError = lipgloss.NewStyle().
			Foreground(Error).
			Bold(true)

	StatusInfo = lipgloss.NewStyle().
			Foreground(Info).
			Bold(true)

	// Content styling
	Content = lipgloss.NewStyle().
		Foreground(TextPrimary).
		Background(Surface).
		Padding(1).
		MarginBottom(1)

	// Cursor and selection
	Cursor = lipgloss.NewStyle().
		Foreground(Primary).
		Bold(true)

	Selected = lipgloss.NewStyle().
			Foreground(Base).
			Background(Primary).
			Bold(true)

	// Help and hint text
	Help = lipgloss.NewStyle().
		Foreground(TextTertiary).
		Italic(true)

	// Icons and symbols
	Icon = lipgloss.NewStyle().
		Foreground(Primary).
		Bold(true)

	// Loading and spinner
	Loading = lipgloss.NewStyle().
		Foreground(Info).
		Italic(true)
)

// GetStatusColor returns appropriate color for status
func GetStatusColor(status string) lipgloss.Color {
	switch status {
	case "success", "completed", "clean":
		return Success
	case "warning", "pending", "modified":
		return Warning
	case "error", "failed", "conflict":
		return Error
	case "info", "building", "progress":
		return Info
	default:
		return TextSecondary
	}
}

// GetStatusIcon returns appropriate icon for status
func GetStatusIcon(status string) string {
	switch status {
	case "success", "completed":
		return "✓"
	case "warning", "pending":
		return "⚠"
	case "error", "failed":
		return "✗"
	case "info", "building":
		return "ℹ"
	case "clean":
		return "✓"
	case "modified":
		return "●"
	case "added":
		return "+"
	case "deleted":
		return "-"
	default:
		return "?"
	}
}
