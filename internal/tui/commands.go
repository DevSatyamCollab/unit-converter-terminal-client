package tui

import (
	"unit-converter-terminal-client/internal/api"

	tea "github.com/charmbracelet/bubbletea"
)

func DoConversionCmd(url string, uc *api.UnitConverter) tea.Cmd {
	return func() tea.Msg {
		result, err := api.ConversionRequest(url, uc)
		if err != nil {
			return MsgConvertError{Err: err}
		}

		return MsgConvertSuccess(*result)
	}
}
