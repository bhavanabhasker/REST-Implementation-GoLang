package main

type Location struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Coordinates struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	}
}

type Locations []Location
