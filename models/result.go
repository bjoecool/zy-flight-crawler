package models

import (
	"time"
	"fmt"
)

type TripResult struct {
	Origin, OriginName           string
	Destination, DestinationName string
	TripDetails                  []TripDetail
}

type TripDetail struct {
	DateDetail   DateDetail
	Price        float32
	Unit         string
	FlightNumber string
	Duration     time.Duration
}

type DateDetail struct {
	Start, End time.Time
}

func (trip TripDetail) String() string {
	trip.Duration = trip.DateDetail.End.Sub(trip.DateDetail.Start)
	return fmt.Sprintf("Time: %v - %v\n"+
		"Price: %v %v\n"+
		"Duration: %v\n"+
		"FlightNumber: %v\n",
		trip.DateDetail.Start, trip.DateDetail.End,
		trip.Price, trip.Unit,
		trip.Duration,
		trip.FlightNumber)
}
