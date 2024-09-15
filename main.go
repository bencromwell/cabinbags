package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bencromwell/cabinbags/cabinbag"
	"github.com/fatih/color"
)

type (
	Airlines struct {
		Airlines []cabinbag.Airline `json:"airlines"`
	}

	OwnedBags struct {
		Bags []cabinbag.CabinBag `json:"owned_bags"`
	}
)

func loadFile(filename string, v interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var airlines Airlines
	var ownedBags OwnedBags
	var results []cabinbag.BagCheckResult = make([]cabinbag.BagCheckResult, 0)

	err := loadFile("airlines.json", &airlines)
	if err != nil {
		color.Red("Failed to load airlines.json: %s", err)
		os.Exit(1)
	}

	err = loadFile("bags.json", &ownedBags)
	if err != nil {
		color.Red("Failed to load bags.json: %s", err)
		os.Exit(1)
	}

	for _, bag := range ownedBags.Bags {
		bagCheckResult := cabinbag.BagCheckResult{}

		for _, airline := range airlines.Airlines {
			cabinbag.CheckBag(airline, bag, &bagCheckResult)
		}

		bagResult := cabinbag.BagCheckResult{
			Bag:               &bag,
			BagAirlineResults: bagCheckResult.BagAirlineResults,
		}

		results = append(results, bagResult)
	}

	for _, result := range results {
		color.Yellow(result.Bag.Ref())

		for _, bagAirlineResult := range result.BagAirlineResults {
			color.Cyan("  %s", bagAirlineResult.Airline.Name)

			for _, limit := range bagAirlineResult.LimitResults {
				if limit.Limit == nil {
					continue
				}

				if limit.Fits {
					color.Green("    - %s: fits", limit.Limit.Ref())
				} else {
					color.Red("    - %s: does not fit", limit.Limit.Ref())
				}
			}
		}

		fmt.Println()
	}
}
