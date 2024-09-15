package main

import (
	"encoding/json"
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

	for _, airline := range airlines.Airlines {
		color.Cyan(airline.Name)

		for _, bag := range ownedBags.Bags {
			cabinbag.CheckBag(airline, bag)
		}
	}
}
