package main

import (
	"log"
	"os"
	"unit-converter-terminal-client/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if _, err := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen()).Run(); err != nil {
		log.Fatalf("Error: %v\n", err)
		os.Exit(1)
	}
}
