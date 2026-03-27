package internal

type UnitConverter struct {
	FromUnit string
	ToUnit   string
	Val      float32
	Ans      float32
}

func NewUnitConverter(funit, tunit string, v, a float32) *UnitConverter {
	return &UnitConverter{
		FromUnit: funit,
		ToUnit:   tunit,
		Val:      v,
		Ans:      a,
	}
}
