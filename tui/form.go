package tui

import (
	"errors"
	"strconv"

	"github.com/charmbracelet/huh"
)

func newLengthForm() *huh.Form {
	var val, fromUnit, toUnit string

	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the length to convert").
				CharLimit(10).
				Value(&val).
				Validate(func(s string) error {
					return isFloat(s)
				}),

			huh.NewSelect[string]().
				Title("Unit to Convert from").
				Options(
					huh.NewOptions("mm", "cm", "m", "km", "in", "ft", "yd", "mil")...,
				).
				Value(&fromUnit),

			huh.NewSelect[string]().
				Title("Unit to Convert to").
				Options(
					huh.NewOptions("mm", "cm", "m", "km", "in", "ft", "yd", "mil")...,
				).
				Value(&toUnit),
		),
	)
}

func newWeightForm() *huh.Form {
	var val, fromUnit, toUnit string

	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the weight to convert").
				CharLimit(10).
				Value(&val).
				Validate(func(s string) error {
					return isFloat(s)
				}),

			huh.NewSelect[string]().
				Title("Unit to Convert from").
				Options(
					huh.NewOptions("mg", "g", "kg", "oz", "lb")...,
				).
				Value(&fromUnit),

			huh.NewSelect[string]().
				Title("Unit to Convert to").
				Options(
					huh.NewOptions("mg", "g", "kg", "oz", "lb")...,
				).
				Value(&toUnit),
		),
	)
}

func newTemperatureForm() *huh.Form {
	var val, fromUnit, toUnit string

	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the temperature to convert").
				CharLimit(10).
				Value(&val).
				Validate(func(s string) error {
					return isFloat(s)
				}),

			huh.NewSelect[string]().
				Title("Unit to Convert from").
				Options(
					huh.NewOptions("c", "k", "f")...,
				).
				Value(&fromUnit),

			huh.NewSelect[string]().
				Title("Unit to Convert to").
				Options(
					huh.NewOptions("c", "k", "f")...,
				).
				Value(&toUnit),
		),
	)
}

var formFactories = []func() *huh.Form{
	newLengthForm,
	newWeightForm,
	newTemperatureForm,
}

func isFloat(s string) error {
	if _, err := strconv.ParseFloat(s, 32); err != nil {
		return errors.New("invalid Input")
	}
	return nil
}
