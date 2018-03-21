package main

import (
	"github.com/zhangyuyu/zy-flight-crawler/service"
	"github.com/zhangyuyu/zy-flight-crawler/models"
	"fmt"
	"time"
)

func main() {
	// 瑞安航空公司 Ryanair
	var tripDetails []models.TripDetail
	startDateOut := "2018-05-14"
	origin := "HAM"
	destination := "BCN"
	weeks := 4

	for i := 0; i < weeks; i++ {
		url := fmt.Sprintf("https://desktopapps.ryanair.com/v4/zh-cn/availability"+
			"?Origin=%s"+
			"&Destination=%s"+
			"&DateOut=%s"+
			"&FlexDaysOut=%d"+
			"&INF=%d&ADT=%d&CHD=%d&TEEN=%d"+
			"&IncludeConnectingFlights=%t&RoundTrip=%t"+
			"&ToUs=AGREED&exists=false&promoCode=",
			origin, destination, startDateOut, 6, 0, 1, 0, 0, true, false)

		result := service.Crawl(url, &service.TripFetcher{})

		tripDetails = append(tripDetails, result.TripDetails...)
		startDateOut = nextStartDate(startDateOut)
	}

	tripResult := models.TripResult{
		TripDetails: tripDetails,
		Origin:      origin,
		Destination: destination,
	}

	service.DrawGraph(tripResult)
}
func nextStartDate(startDateOut string) string {
	parse, _ := time.Parse("2006-01-02", startDateOut)
	nextDate := parse.AddDate(0, 0, 7)
	return nextDate.Format("2006-01-02")
}
