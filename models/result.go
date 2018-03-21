package models

import (
	"time"
	"fmt"
	"github.com/bclicn/color"
)

type TripResult struct {
	Origin, OriginName           string
	Destination, DestinationName string
	TripDetails                  []TripDetail
}

type TripDetail struct {
	Start, End   time.Time
	Price        float64
	Unit         string
	FlightNumber string
	Duration     time.Duration
}


func (trip TripDetail) String() string {
	trip.Duration = trip.End.Sub(trip.Start)
	return fmt.Sprintf(color.Blue("Time") + ": %v - %v\n"+
		color.Blue("Price")+ ": %v %v\n"+
		color.Blue("Duration")+ ": %v\n"+
		color.Blue("FlightNumber")+ ": %v\n",
		trip.Start, trip.End,
		trip.Price, trip.Unit,
		trip.Duration,
		trip.FlightNumber)
}
