package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var defaultWidth = 55

type Styles struct {
	container   lipgloss.Style
	header      lipgloss.Style
	footer      lipgloss.Style
	activeTab   lipgloss.Style
	inactiveTab lipgloss.Style
	window      lipgloss.Style
	error       lipgloss.Style
}

func DefaultStyle(width, height int) *Styles {
	highlightColor := lipgloss.Color("#7D56F4")
	containerHeight, containerWidth := height-2, width-4

	s := new(Styles)

	s.container = lipgloss.NewStyle().
		Width(containerWidth).
		Height(containerHeight).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("205"))

	s.header = lipgloss.NewStyle().
		Foreground(lipgloss.Color("87")).
		Bold(true).
		MarginLeft(1).MarginTop(1)

	s.footer = lipgloss.NewStyle().
		Foreground(lipgloss.Color("106")).
		MarginTop(1).PaddingLeft(1)

	s.activeTab = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(highlightColor).
		Margin(1).
		UnsetMarginBottom()

	s.inactiveTab = s.activeTab.UnsetBorderBottom()

	s.window = lipgloss.NewStyle().
		Border(lipgloss.HiddenBorder()).
		Height(height - 33).
		Width(width - 8)

	s.error = s.window.
		Foreground(lipgloss.Color("196"))

	return s
}
