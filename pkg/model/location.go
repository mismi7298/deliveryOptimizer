package model

type Restaurant struct {
	Location Location
	Name     string
	Id       string
	PrepTime float64
}

type Customer struct {
	Location   Location
	CustomerId string
}

type Location struct {
	Id        string
	Latitude  float64
	Longitude float64
}
