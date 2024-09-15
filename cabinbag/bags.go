package cabinbag

import (
	"fmt"
)

type (
	CabinBag struct {
		Name string `json:"name"`
		X    int    `json:"x"`
		Y    int    `json:"y"`
		Z    int    `json:"z"`
	}

	LimitResult struct {
		Limit *CabinBag
		Fits  bool
	}

	BagAirlineResult struct {
		Airline      *Airline
		LimitResults []LimitResult
	}

	BagCheckResult struct {
		Bag               *CabinBag
		BagAirlineResults []BagAirlineResult
	}
)

func (c CabinBag) FitsIn(limit *CabinBag) bool {
	return c.X <= limit.X && c.Y <= limit.Y && c.Z <= limit.Z
}

func (c CabinBag) Ref() string {
	return fmt.Sprintf("%s (%dx%dx%d)", c.Name, c.X, c.Y, c.Z)
}

func CheckBag(airline Airline, bag CabinBag, bagCheckResult *BagCheckResult) {
	limits := [2]*CabinBag{
		airline.SmallBagLimit,
		airline.LargeBagLimit,
	}

	result := BagAirlineResult{
		Airline: &airline,
	}

	for _, limit := range limits {
		if limit == nil {
			continue
		}

		result.LimitResults = append(result.LimitResults, LimitResult{
			Limit: limit,
			Fits:  bag.FitsIn(limit),
		})
	}

	bagCheckResult.BagAirlineResults = append(bagCheckResult.BagAirlineResults, result)
}
