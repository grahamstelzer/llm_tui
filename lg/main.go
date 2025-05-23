package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input string
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

		case "enter":
			// Reset input on enter
			m.input = ""
		default:
			// Append typed key
			m.input += msg.String()
		}
	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf(
		"Type something! (Press 'q' or ctrl+c to quit)\n\nInput: %s",
		m.input,
	)
}

func main() {
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}




// package main

// import (
// 	"fmt"

// 	"github.com/charmbracelet/lipgloss"
// )

// func main() {
// 	headerStyle := lipgloss.NewStyle().
// 		Bold(true).
// 		Foreground(lipgloss.Color("204")).
// 		Align(lipgloss.Center).
// 		Border(lipgloss.ThickBorder()).
// 		Padding(1, 2)

// 	bodyStyle := lipgloss.NewStyle().
// 		Foreground(lipgloss.Color("252")).
// 		Padding(1, 2)

// 	footerStyle := lipgloss.NewStyle().
// 		Foreground(lipgloss.Color("241")).
// 		Italic(true).
// 		Align(lipgloss.Center).
// 		Padding(1, 2)

// 	header := headerStyle.Render("🧠 RamaLama Terminal UI")
// 	body := bodyStyle.Render("Use this tool to interact with your language model.\n\nComing soon: model selection, prompt input, and responses!")
// 	footer := footerStyle.Render("Press CTRL+C to exit")

// 	ui := lipgloss.JoinVertical(lipgloss.Top, header, body, footer)
// 	fmt.Println(ui)
// }
