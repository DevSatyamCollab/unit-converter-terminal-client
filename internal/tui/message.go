package tui

import "unit-converter-terminal-client/internal/api"

type MsgConvertSuccess api.UnitConverter
type MsgConvertError struct{ Err error }
