package cabinbag

import "github.com/fatih/color"

type CabinBag struct {
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Z    int    `json:"z"`
}

func (c CabinBag) FitsIn(limit *CabinBag) bool {
	return c.X <= limit.X && c.Y <= limit.Y && c.Z <= limit.Z
}

func CheckBag(airline Airline, bag CabinBag) {
	limits := [2]*CabinBag{
		airline.SmallBagLimit,
		airline.LargeBagLimit,
	}

	fitFound := false

	for _, limit := range limits {
		if limit == nil {
			continue
		}

		if bag.FitsIn(limit) {
			color.Green("  %s fits within %s %s limit", bag.Name, limit.Name, airline.Name)
			fitFound = true
		} else {
			color.Yellow("  %s does not fit within %s %s limit", bag.Name, limit.Name, airline.Name)
		}
	}

	if !fitFound {
		color.Red("  %s does not fit within %s limits", bag.Name, airline.Name)
	}
}
