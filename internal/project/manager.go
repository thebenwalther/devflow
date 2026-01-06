// Package project provides project discovery and management functionality
package project

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thebenwalther/devflow/internal/ui/styles"
)

// Project represents a discovered development project
type Project struct {
	Name     string
	Path     string
	Language string
	Type     string
	Git      bool
	Status   string
	Modified time.Time
	Files    int
}

// ProjectType represents different project types we can detect
type ProjectType int

const (
	UnknownType ProjectType = iota
	NodeJS
	Go
	Rust
	Python
	Cargo // Rust project
	Makefile
	GitRepo
)

// Model holds the state for project management
type Model struct {
	projects []Project
	cursor   int
	selected int
	loading  bool
	error    error
	width    int
	height   int
}

// New creates a new project management model
func New() *Model {
	return &Model{
		projects: []Project{},
		cursor:   0,
		selected: 0,
		loading:  true,
		width:    80,
		height:   24,
	}
}

// detectProjectType determines what type of project this is
func detectProjectType(path string) ProjectType {
	info, err := os.Stat(path)
	if err != nil {
		return UnknownType
	}

	if !info.IsDir() {
		return UnknownType
	}

	// Check for project indicators in priority order
	files, err := os.ReadDir(path)
	if err != nil {
		return UnknownType
	}

	fileNameMap := make(map[string]bool)
	for _, file := range files {
		fileNameMap[file.Name()] = true
	}

	// Node.js (package.json)
	if fileNameMap["package.json"] {
		return NodeJS
	}

	// Go (go.mod)
	if fileNameMap["go.mod"] {
		return Go
	}

	// Rust (Cargo.toml)
	if fileNameMap["Cargo.toml"] {
		return Cargo
	}

	// Python (pyproject.toml, requirements.txt, setup.py)
	if fileNameMap["pyproject.toml"] || fileNameMap["requirements.txt"] || fileNameMap["setup.py"] {
		return Python
	}

	// Makefile
	if fileNameMap["Makefile"] || fileNameMap["makefile"] {
		return Makefile
	}

	// Git repository (.git)
	if fileNameMap[".git"] {
		return GitRepo
	}

	return UnknownType
}

// getLanguageFromType returns display language for project type
func getLanguageFromType(projectType ProjectType) string {
	switch projectType {
	case NodeJS:
		return "Node.js"
	case Go:
		return "Go"
	case Rust, Cargo:
		return "Rust"
	case Python:
		return "Python"
	case Makefile:
		return "Make"
	case GitRepo:
		return "Git"
	default:
		return "Unknown"
	}
}

// getIconForType returns an icon for project type
func getIconForType(projectType ProjectType) string {
	switch projectType {
	case NodeJS:
		return "🟢"
	case Go:
		return "🐹"
	case Rust, Cargo:
		return "🦀"
	case Python:
		return "🐍"
	case Makefile:
		return "⚙"
	case GitRepo:
		return "📁"
	default:
		return "📁"
	}
}

// getProjectInfo extracts information about a project
func getProjectInfo(path string) Project {
	projectType := detectProjectType(path)
	name := filepath.Base(path)

	// Get modified time
	info, err := os.Stat(path)
	var modified time.Time
	if err == nil {
		modified = info.ModTime()
	}

	// Count files (simplified)
	fileCount := 0
	projectFiles, _ := os.ReadDir(path)
	for _, file := range projectFiles {
		if !file.IsDir() {
			fileCount++
		}
	}

	// Check git status (simplified)
	hasGit := false
	if _, err := os.Stat(filepath.Join(path, ".git")); err == nil {
		hasGit = true
	}

	status := "Clean"
	if hasGit {
		status = "Git Repo"
	} else if fileCount > 0 {
		status = "Active"
	} else {
		status = "Empty"
	}

	return Project{
		Name:     name,
		Path:     path,
		Language: getLanguageFromType(projectType),
		Type:     getLanguageFromType(projectType),
		Git:      hasGit,
		Status:   status,
		Modified: modified,
		Files:    fileCount,
	}
}

// DiscoverProjects scans common directories for projects
func DiscoverProjects() []Project {
	var allProjects []Project

	// Search paths to scan
	searchPaths := []string{
		".",
		"~/dev",
		"~/Projects",
		"~/code",
		"~/workspace",
	}

	for _, searchPath := range searchPaths {
		// Expand ~ to home directory
		expandedPath := os.ExpandEnv(searchPath)

		// Check if path exists
		if _, err := os.Stat(expandedPath); os.IsNotExist(err) {
			continue
		}

		// Walk through directory
		filepath.WalkDir(expandedPath, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}

			// Skip hidden directories and common non-project directories
			if strings.HasPrefix(d.Name(), ".") ||
				d.Name() == "node_modules" ||
				d.Name() == "target" ||
				d.Name() == "build" ||
				d.Name() == "dist" {
				return nil
			}

			// Only check directories
			if !d.IsDir() {
				return nil
			}

			// Don't go too deep
			depth := strings.Count(path, string(filepath.Separator))
			if depth > 3 {
				return nil
			}

			// Check if this looks like a project
			if detectProjectType(path) != UnknownType {
				project := getProjectInfo(path)
				allProjects = append(allProjects, project)
			}

			return nil
		})
	}

	return allProjects
}

// Update handles updates to the project model
func (m *Model) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return nil

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.projects)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selected = m.cursor
			return func() tea.Msg {
				return ProjectSelectedMsg{Project: m.projects[m.selected]}
			}
		case "r":
			// Refresh project list
			m.loading = true
			return func() tea.Msg {
				return ProjectsLoadedMsg{Projects: DiscoverProjects()}
			}
		}

	case ProjectsLoadedMsg:
		m.projects = msg.Projects
		m.loading = false
		if len(m.projects) == 0 {
			m.cursor = 0
			m.selected = 0
		} else {
			m.cursor = 0
			m.selected = 0
		}

	}

	return nil
}

// View renders the project management interface
func (m *Model) View() string {
	if m.error != nil {
		return styles.StatusError.Render("Error: " + m.error.Error())
	}

	// Header
	header := styles.Header.Render("📁 Projects")

	if m.loading {
		loading := styles.Loading.Render("Scanning for projects...")
		content := lipgloss.JoinVertical(
			lipgloss.Left,
			loading,
		)
		return lipgloss.JoinVertical(
			lipgloss.Top,
			header,
			styles.Panel.Render(content),
		)
	}

	if len(m.projects) == 0 {
		noProjects := lipgloss.NewStyle().
			Foreground(styles.TextSecondary).
			Italic(true).
			Render("No projects found.\n\nPress 'r' to scan again.")

		content := lipgloss.JoinVertical(
			lipgloss.Left,
			noProjects,
		)
		return lipgloss.JoinVertical(
			lipgloss.Top,
			header,
			styles.Panel.Render(content),
		)
	}

	// Project list
	var projectStrings []string
	for i, project := range m.projects {
		// Cursor
		cursor := "  "
		if i == m.cursor {
			cursor = "❯ "
		}

		// Selection highlight
		var nameStyle lipgloss.Style
		if i == m.selected {
			nameStyle = styles.Selected
		} else {
			nameStyle = lipgloss.NewStyle().Foreground(styles.TextPrimary)
		}

		// Project name
		name := nameStyle.Render(project.Name)

		// Project details
		detailsStyle := lipgloss.NewStyle().
			Foreground(styles.TextSecondary).
			Italic(true)

		details := detailsStyle.Render(fmt.Sprintf("  %s • %s • %s",
			getIconForType(detectProjectType(project.Path)),
			project.Language,
			project.Status))

		// Add details on same line if cursor
		projectStr := fmt.Sprintf("%s%s", cursor, name)
		if i == m.cursor {
			projectStr += " " + details
		} else {
			projectStr += "\n" + details
		}

		projectStrings = append(projectStrings, projectStr)
	}

	content := strings.Join(projectStrings, "\n")

	// Help text
	help := styles.Help.Render("↑↓: Navigate | Enter: Select | r: Refresh")

	fullContent := lipgloss.JoinVertical(
		lipgloss.Left,
		content,
		"",
		help,
	)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		styles.Panel.Render(fullContent),
	)
}

// Message types for project management
type ProjectsLoadedMsg struct {
	Projects []Project
}

type ProjectSelectedMsg struct {
	Project Project
}
