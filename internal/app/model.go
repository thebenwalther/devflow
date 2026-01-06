// Package app contains the main application model and logic
package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thebenwalther/devflow/internal/project"
	"github.com/thebenwalther/devflow/internal/ui/styles"
)

// FocusArea represents which area of the UI is currently focused
type FocusArea int

const (
	ProjectsFocus FocusArea = iota
	GitFocus
	BuildFocus
	TasksFocus
)

// Model represents the main application state
type Model struct {
	// Navigation state
	currentTab string
	focused    FocusArea

	// Application state
	dimensions tea.WindowSizeMsg
	quitting   bool
	err        error

	// Content for each tab
	projectsContent string
	gitContent      string
	buildContent    string
	tasksContent    string

	// Frame counter for animations
	frameCount int

	// Project management integration
	projectManager *project.Model
}

// New creates a new application model with initial state
func New() *Model {
	return &Model{
		currentTab: "projects",
		focused:    ProjectsFocus,

		projectsContent: "📁 Projects\n\nReady to discover your development projects!",
		gitContent:      "🔀 Git\n\nComing soon: git status and operations",
		buildContent:    "🔨 Build\n\nComing soon: build monitoring",
		tasksContent:    "📋 Tasks\n\nComing soon: task management",

		frameCount: 0,

		// Initialize project manager
		projectManager: project.New(),
	}
}

// Init initializes the application
func (m *Model) Init() tea.Cmd {
	// Start project discovery
	return func() tea.Msg {
		projects := project.DiscoverProjects()
		return project.ProjectsLoadedMsg{Projects: projects}
	}()
}

// Update handles application updates and messages
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.dimensions = msg
		return m, nil

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC || msg.String() == "q" {
			m.quitting = true
			return m, tea.Quit
		}

		// Handle tab switching
		switch msg.String() {
		case "tab", "l":
			m.switchTab(1)
		case "shift+tab", "h":
			m.switchTab(-1)
		case "1":
			m.currentTab = "projects"
			m.focused = ProjectsFocus
		case "2":
			m.currentTab = "git"
			m.focused = GitFocus
		case "3":
			m.currentTab = "build"
			m.focused = BuildFocus
		case "4":
			m.currentTab = "tasks"
			m.focused = TasksFocus
		}

	case project.ProjectsLoadedMsg:
		m.projectManager.Update(msg)

	case tea.QuitMsg:
		return m, nil
	}

	m.frameCount++
	return m, nil
}

// View renders the application interface
func (m *Model) View() string {
	if m.err != nil {
		return styles.StatusError.Render("Error: " + m.err.Error())
	}

	// Render header
	header := styles.Header.Render(" DevFlow v0.1 ")

	// Render tabs
	tabs := m.renderTabs()

	// Render content based on current tab
	content := m.renderContent()

	// Combine everything
	return lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		tabs,
		content,
	)
}

// switchTab changes the current tab by the given offset
func (m *Model) switchTab(offset int) {
	tabs := []string{"projects", "git", "build", "tasks"}
	currentIdx := 0

	// Find current tab index
	for i, tab := range tabs {
		if tab == m.currentTab {
			currentIdx = i
			break
		}
	}

	// Calculate new index with wraparound
	newIdx := (currentIdx + offset + len(tabs)) % len(tabs)
	m.currentTab = tabs[newIdx]

	// Update focus area based on tab
	switch m.currentTab {
	case "projects":
		m.focused = ProjectsFocus
	case "git":
		m.focused = GitFocus
	case "build":
		m.focused = BuildFocus
	case "tasks":
		m.focused = TasksFocus
	}
}

// renderTabs creates the tab navigation bar
func (m *Model) renderTabs() string {
	tabs := []string{"[1] Projects", "[2] Git", "[3] Build", "[4] Tasks"}
	var tabStyles []lipgloss.Style

	for i, _ := range tabs {
		if (m.currentTab == "projects" && i == 0) ||
			(m.currentTab == "git" && i == 1) ||
			(m.currentTab == "build" && i == 2) ||
			(m.currentTab == "tasks" && i == 3) {
			tabStyles = append(tabStyles, styles.ActiveTab)
		} else {
			tabStyles = append(tabStyles, styles.InactiveTab)
		}
	}

	var renderedTabs []string
	for i, tab := range tabs {
		renderedTabs = append(renderedTabs, tabStyles[i].Render(tab))
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
}

// renderContent returns the content for the current tab
func (m *Model) renderContent() string {
	content := ""

	switch m.currentTab {
	case "projects":
		content = m.projectManager.View()
	case "git":
		content = styles.Content.Render(m.gitContent)
	case "build":
		content = styles.Content.Render(m.buildContent)
	case "tasks":
		content = styles.Content.Render(m.tasksContent)
	default:
		content = styles.Content.Render("Unknown tab")
	}

	// Add help hint at the bottom
	help := styles.Help.Render("Tab: Switch | 1-4: Select Tab | ↑↓: Navigate | Enter: Select | r: Refresh")
	contentWithHelp := lipgloss.JoinVertical(
		lipgloss.Top,
		content,
		"",
		help,
	)

	return styles.Panel.Render(contentWithHelp)
}
