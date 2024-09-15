package cabinbag

type Airline struct {
	Name          string    `json:"name"`
	SmallBagLimit *CabinBag `json:"small_bag_limit"`
	LargeBagLimit *CabinBag `json:"large_bag_limit"`
}
