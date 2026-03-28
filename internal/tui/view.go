package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	header := m.Style.header.Render("Unit Converter")
	footer := m.Style.footer.Render("ctrl+l/right (nextTab) . ctrl+h/left (previousTab) . esc (reset)")

	var renderedTabs []string
	for i, t := range m.Tabs {
		style := m.Style.inactiveTab
		if i == m.ActiveTab {
			style = m.Style.activeTab
		}
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	tabs := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	usedHeight := lipgloss.Height(header) + lipgloss.Height(tabs) + lipgloss.Height(footer)
	windowFrameH := m.Style.window.GetVerticalFrameSize()
	contentHeight := m.Style.container.GetHeight() - usedHeight - windowFrameH - 2
	window := m.Style.window.Height(contentHeight)
	content := m.getContent()

	ui := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		tabs,
		window.Render(content),
		footer,
	)

	return lipgloss.Place(
		m.Width,
		m.Height,
		lipgloss.Center,
		lipgloss.Center,
		m.Style.container.Render(ui),
	)
}

func (m *model) getContent() string {
	if m.ShowingResult[m.ActiveTab] {

		if m.Err != nil {
			return m.Style.error.Render(
				fmt.Sprintf("❌ Error: %v\n\n  Press 'esc' to try again.", m.Err),
			)
		}

		// 1 kg = 1000g
		return fmt.Sprintf(
			"%.f %s = %.1f %s",
			m.UnitConverter[m.ActiveTab].Val,
			m.UnitConverter[m.ActiveTab].FromUnit,
			m.UnitConverter[m.ActiveTab].Ans,
			m.UnitConverter[m.ActiveTab].ToUnit,
		)

	} else {
		return m.Forms[m.ActiveTab].View()
	}
}
