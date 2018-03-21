package models

import "time"

type TripResponse struct {
	Currency string
	Trips [] Trip
}

type Trip struct {
	Origin, OriginName, Destination, DestinationName string
	Dates [] Date
}

type Date struct {
	DateOut string
	Flights [] Flight
}

type Flight struct {
	Duration string
	RegularFare RegularFare
	TimeUTC [] time.Time
	FlightNumber string
	FaresLeft int16
}

type RegularFare struct {
	Fares []Fare
}

type Fare struct {
	Type string
	Count int16
	Amount float64
	PublishedFare float64
	HasDiscount bool
	HasPromoDiscount bool
	DiscountInPercent int16
}