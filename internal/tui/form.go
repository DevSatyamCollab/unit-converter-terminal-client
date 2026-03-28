package tui

import (
	"errors"
	"strconv"

	"github.com/charmbracelet/huh"
)

var (
	LengthVal, LengthFromUnit, LengthToUnit                string
	WeightVal, WeightFromUnit, WeightToUnit                string
	TemperatureVal, TemperatureFromUnit, TemperatureToUnit string
)

func newLengthForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the length to convert").
				CharLimit(10).
				Value(&LengthVal).
				Validate(func(s string) error {
					return isFloat(s)
				}),

			huh.NewSelect[string]().
				Title("Unit to Convert from").
				Options(
					huh.NewOptions("mm", "cm", "m", "km", "in", "ft", "yd", "mil")...,
				).
				Value(&LengthFromUnit),

			huh.NewSelect[string]().
				Title("Unit to Convert to").
				Options(
					huh.NewOptions("mm", "cm", "m", "km", "in", "ft", "yd", "mil")...,
				).
				Value(&LengthToUnit),
		),
	)
}

func newWeightForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the weight to convert").
				CharLimit(10).
				Value(&WeightVal).
				Validate(func(s string) error {
					return isFloat(s)
				}),

			huh.NewSelect[string]().
				Title("Unit to Convert from").
				Options(
					huh.NewOptions("mg", "g", "kg", "oz", "lb")...,
				).
				Value(&WeightFromUnit),

			huh.NewSelect[string]().
				Title("Unit to Convert to").
				Options(
					huh.NewOptions("mg", "g", "kg", "oz", "lb")...,
				).
				Value(&WeightToUnit),
		),
	)
}

func newTemperatureForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the temperature to convert").
				CharLimit(10).
				Value(&TemperatureVal).
				Validate(func(s string) error {
					return isFloat(s)
				}),

			huh.NewSelect[string]().
				Title("Unit to Convert from").
				Options(
					huh.NewOptions("c", "k", "f")...,
				).
				Value(&TemperatureFromUnit),

			huh.NewSelect[string]().
				Title("Unit to Convert to").
				Options(
					huh.NewOptions("c", "k", "f")...,
				).
				Value(&TemperatureToUnit),
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
