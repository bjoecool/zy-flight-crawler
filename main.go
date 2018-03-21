package main

import (
	"github.com/zhangyuyu/zy-flight-crawler/service"
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Today: %v\n", time.Now().Format(time.RFC1123))

	// 瑞安航空公司 Ryanair
	dateOut := "2018-05-14"
	flexDaysOut := 6
	url := fmt.Sprintf("https://desktopapps.ryanair.com/v4/zh-cn/availability"+
		"?Origin=%s"+
		"&Destination=%s"+
		"&DateOut=%s"+
		"&FlexDaysOut=%d"+
		"&INF=%d&ADT=%d&CHD=%d&TEEN=%d"+
		"&IncludeConnectingFlights=%t&RoundTrip=%t"+
		"&ToUs=AGREED&exists=false&promoCode=",
		"HAM", "BCN", dateOut, flexDaysOut, 0, 2, 0, 0, true, false)

	service.Crawl(url, flexDaysOut, &service.TripFetcher{})
}
