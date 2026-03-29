package tui

import (
	"fmt"
	"math"

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

		// formating
		var valstr, ansStr string
		val := m.UnitConverter[m.ActiveTab].Val
		ans := m.UnitConverter[m.ActiveTab].Ans

		valDiff := math.Abs(val - math.Round(val))
		if valDiff == 0 {
			valstr = fmt.Sprintf("%.f", val)
		} else if valDiff < 0.001 {
			valstr = fmt.Sprintf("%g", val)
		} else {
			valstr = fmt.Sprintf("%.2f", val)
		}

		ansDiff := math.Abs(ans - math.Round(ans))
		if ansDiff == 0 {
			ansStr = fmt.Sprintf("%.f", ans)
		} else if ansDiff < 0.001 {
			ansStr = fmt.Sprintf("%g", ans)
		} else {
			ansStr = fmt.Sprintf("%.2f", ans)
		}

		// 1 kg = 1000g
		return fmt.Sprintf(
			"%s %s = %s %s\n\nPress 'esc' to reset",
			valstr,
			m.UnitConverter[m.ActiveTab].FromUnit,
			ansStr,
			m.UnitConverter[m.ActiveTab].ToUnit,
		)

	} else {
		return m.Forms[m.ActiveTab].View()
	}
}
