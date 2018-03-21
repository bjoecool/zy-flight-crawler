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
	DateOut, Start, End time.Time
	Price               float64
	Unit                string
	FlightNumber        string
	Duration            float64
}

func (trip TripDetail) String() string {
	if trip.FlightNumber == "" {
		return fmt.Sprintf(color.Blue("DateOut")+": %v\n", trip.DateOut.Format("2006-01-02"))
	}

	trip.Duration = trip.End.Sub(trip.Start).Hours()
	return fmt.Sprintf(color.Blue("DateOut") + ": %v\n"+
		color.Blue("Time")+ ": %v - %v\n"+
		color.Blue("Price")+ ": %v %v\n"+
		color.Blue("Duration")+ ": %vh\n"+
		color.Blue("FlightNumber")+ ": %v\n",
		trip.DateOut.Format("2006-01-02"),
		trip.Start.Format("15:04:05"), trip.End.Format("15:04:05 MST"),
		trip.Price, trip.Unit,
		trip.Duration,
		trip.FlightNumber)
}
