package conv

import "fmt"

type Meter float64
type Centimeter float64
type Feet float64
type Kilometer float64

func (m Meter) String() string {
	return fmt.Sprintf("%g m.", m)
}

func (c Centimeter) String() string {
	return fmt.Sprintf("%g cm.", c)
}

func (f Feet) String() string {
	return fmt.Sprintf("%g ft.", f)
}

func (k Kilometer) String() string {
	return fmt.Sprintf("%g km.", k)
}
