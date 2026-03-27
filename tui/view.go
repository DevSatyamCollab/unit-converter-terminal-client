package tui

import "github.com/charmbracelet/lipgloss"

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

	var content string
	if m.ShowingResult && m.ActiveTab == m.ResultTab {
		content = "Result"
	} else {
		content = m.Forms[m.ActiveTab].View()
	}

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
