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
	today := time.Now()
	startDateOut := today.AddDate(0, 0, 1)
	origin := "HAM"
	destination := "BCN"
	weeks := 4

	for i := 0; i < weeks; i++ {
		url := fmt.Sprintf("https://desktopapps.ryanair.com/v4/zh-cn/availability"+
			"?Origin=%s"+
			"&Destination=%s"+
			"&DateOut=%s"+
			"&FlexDaysOut=%d"+
			"&ADT=%d&CHD=%d&INF=%d&TEEN=%d"+
			"&IncludeConnectingFlights=%t&RoundTrip=%t"+
			"&ToUs=AGREED&exists=false",
			origin, destination, startDateOut.Format("2006-01-02"), 6, 1, 0, 0, 0, true, false)

		result := service.Crawl(url, &service.TripFetcher{})

		tripDetails = append(tripDetails, result.TripDetails...)
		startDateOut = startDateOut.AddDate(0, 0, 7)
	}

	tripResult := models.TripResult{
		TripDetails: tripDetails,
		Origin:      origin,
		Destination: destination,
	}

	service.DrawGraph(tripResult)
}
