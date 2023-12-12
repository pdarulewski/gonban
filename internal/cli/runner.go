package cli

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Run(m *Board) {
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
