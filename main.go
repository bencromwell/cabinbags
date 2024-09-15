package main

import (
	"github.com/bencromwell/cabinbags/cabinbag"
	"github.com/fatih/color"
)

func AvailableBags() []cabinbag.CabinBag {
	return []cabinbag.CabinBag{
		{Name: "Small", X: 40, Y: 25, Z: 20},
		{Name: "Large", X: 50, Y: 33, Z: 20},
		{Name: "Test", X: 60, Y: 40, Z: 20},
	}
}

func Airlines() []cabinbag.Airline {
	return []cabinbag.Airline{
		{
			Name:          "Ryanair",
			SmallBagLimit: &cabinbag.CabinBag{Name: "Small", X: 40, Y: 25, Z: 20},
			LargeBagLimit: &cabinbag.CabinBag{Name: "Large", X: 55, Y: 40, Z: 20},
		},
		{
			Name:          "EasyJet",
			SmallBagLimit: &cabinbag.CabinBag{Name: "Small", X: 45, Y: 36, Z: 20},
			LargeBagLimit: &cabinbag.CabinBag{Name: "Large", X: 56, Y: 45, Z: 25},
		},
		{
			Name:          "Jet2",
			LargeBagLimit: &cabinbag.CabinBag{Name: "Large", X: 56, Y: 45, Z: 25},
		},
	}
}

func main() {
	airlines := Airlines()
	ownedBags := AvailableBags()

	for _, airline := range airlines {
		color.Cyan(airline.Name)

		for _, bag := range ownedBags {
			cabinbag.CheckBag(airline, bag)
		}
	}
}
