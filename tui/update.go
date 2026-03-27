package tui

import (
	"strconv"
	"unit-converter-terminal-client/internal"

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

		case "esc": // reset only the current tab
			if m.ShowingResult[m.ActiveTab] {
				switch m.ActiveTab {
				case 0:
					LengthVal, LengthFromUnit, LengthToUnit = "", "", ""
				case 1:
					WeightVal, WeightFromUnit, WeightToUnit = "", "", ""
				case 2:
					TemperatureVal, TemperatureFromUnit, TemperatureToUnit = "", "", ""
				}

				m.ShowingResult[m.ActiveTab] = false
				m.Forms[m.ActiveTab] = m.Factories[m.ActiveTab]()
				return m, m.Forms[m.ActiveTab].Init()
			}
		}
	}

	newForm, cmd := m.Forms[m.ActiveTab].Update(msg)
	if f, ok := newForm.(*huh.Form); ok {
		m.Forms[m.ActiveTab] = f
		cmds = append(cmds, cmd)
	}

	// trigger result
	if !m.ShowingResult[m.ActiveTab] && m.Forms[m.ActiveTab].State == huh.StateCompleted {
		var valStr, funit, tunit string

		// Determine which variables to read based on the active tab
		switch m.ActiveTab {
		case 0:
			valStr = LengthVal
			funit = LengthFromUnit
			tunit = LengthToUnit
		case 1:
			valStr = WeightVal
			funit = WeightFromUnit
			tunit = WeightToUnit
		case 2:
			valStr = TemperatureVal
			funit = TemperatureFromUnit
			tunit = TemperatureToUnit
		}

		// Directly parse the correctly selected string variable
		val, _ := strconv.ParseFloat(valStr, 64)

		m.UnitConverter[m.ActiveTab] = internal.NewUnitConverter(funit, tunit, float32(val), float32(0))
		m.ShowingResult[m.ActiveTab] = true
	}

	return m, tea.Batch(cmds...)
}
