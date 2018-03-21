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
	Duration     float64
}

func (trip TripDetail) String() string {
	trip.Duration = trip.End.Sub(trip.Start).Hours()
	return fmt.Sprintf(color.Blue("Time") + ": %v - %v\n"+
		color.Blue("Price")+ ": %v %v\n"+
		color.Blue("Duration")+ ": %vh\n"+
		color.Blue("FlightNumber")+ ": %v\n",
		trip.Start.Format("2006-01-02 15:04:05"), trip.End.Format("2006-01-02 15:04:05 MST"),
		trip.Price, trip.Unit,
		trip.Duration,
		trip.FlightNumber)
}
