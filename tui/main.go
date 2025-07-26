package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(InitialModel())
	if err := p.Start(); err != nil {
		log.Fatalf("Error starting program: %v", err)
	}
}

func Start() error {
	p := tea.NewProgram(InitialModel())
	return p.Start()
}
