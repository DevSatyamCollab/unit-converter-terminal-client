package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.ContentWidth = min(defaultWidth, msg.Width)
		m.Style = DefaultStyle(m.ContentWidth, msg.Height)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "ctrl+l", "right":
			m.ActiveTab = min(m.ActiveTab+1, len(m.Tabs)-1)
			cmds = append(cmds, m.Forms[m.ActiveTab].Init())

		case "ctrl+h", "left":
			m.ActiveTab = max(m.ActiveTab-1, 0)
			cmds = append(cmds, m.Forms[m.ActiveTab].Init())

		case "esc":
			if m.ShowingResult {
				m.ShowingResult = false
				m.Forms[m.ResultTab] = m.Factories[m.ResultTab]()
				return m, m.Forms[m.ResultTab].Init()
			}
		}
	}

	newForm, cmd := m.Forms[m.ActiveTab].Update(msg)
	if f, ok := newForm.(*huh.Form); ok {
		m.Forms[m.ActiveTab] = f
		cmds = append(cmds, cmd)
	}

	// trigger result
	if !m.ShowingResult && m.Forms[m.ActiveTab].State == huh.StateCompleted {
		m.ResultTab = m.ActiveTab
		m.ShowingResult = true
	}

	return m, tea.Batch(cmds...)
}
