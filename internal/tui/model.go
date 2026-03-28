package tui

import (
	"unit-converter-terminal-client/internal/api"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

const numTabs = 3

type model struct {
	Tabs          []string
	Forms         []*huh.Form
	Factories     []func() *huh.Form
	Style         *Styles
	UnitConverter []*api.UnitConverter
	Width         int
	Height        int
	ActiveTab     int
	ContentWidth  int
	ShowingResult []bool
	Err           error
}

func (m model) Init() tea.Cmd {
	return m.Forms[m.ActiveTab].Init()
}

func InitialModel() model {
	tabs := []string{"length", "weight", "temperature"}
	forms := make([]*huh.Form, len(formFactories))

	for i, f := range formFactories {
		forms[i] = f()
	}

	return model{
		Tabs:          tabs,
		Forms:         forms,
		Factories:     formFactories,
		Style:         DefaultStyle(0, 0),
		ActiveTab:     0,
		ShowingResult: make([]bool, numTabs),
		UnitConverter: make([]*api.UnitConverter, numTabs),
	}
}
