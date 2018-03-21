package service

import (
	"fmt"
	"github.com/zhangyuyu/zy-trip-spider/models"
	"net/http"
	_ "io/ioutil"
	_ "encoding/json"
	"io/ioutil"
	"encoding/json"
)

func Crawl(url string, days int, fetcher Fetcher) {
	response, _ := fetcher.Fetch(url)

	tripResponse := Analysis(response)
	tripResult := Transfer(tripResponse, days)

	//bytes, _ := json.Marshal(tripResult)
	//fmt.Printf(string(bytes))
	fmt.Printf("Origin: %v(%v)\n"+"Destination: %v(%v)\n\n",
		tripResult.Origin, tripResult.OriginName, tripResult.Destination, tripResult.DestinationName)

	details := tripResult.TripDetails
	for i := 0; i < days; i++ {
		fmt.Println(details[i])
	}
}

func Transfer(tripResponse models.TripResponse, days int) models.TripResult {
	tripDetails := make([]models.TripDetail, days)
	unit := tripResponse.Currency

	dates := tripResponse.Trips[0].Dates
	for i := 0; i < days; i++ {
		day := dates[i]
		tripDetails[i] = models.TripDetail{
			Price:        day.Flights[0].RegularFare.Fares[0].Amount,
			Unit:         unit,
			FlightNumber: day.Flights[0].FlightNumber,
			DateDetail: models.DateDetail{
				Start: day.Flights[0].TimeUTC[0],
				End:   day.Flights[0].TimeUTC[1],
			},
		}
	}

	tripResult := models.TripResult{
		Origin:          tripResponse.Trips[0].Origin,
		OriginName:      tripResponse.Trips[0].OriginName,
		Destination:     tripResponse.Trips[0].Destination,
		DestinationName: tripResponse.Trips[0].DestinationName,
		TripDetails:     tripDetails,
	}

	return tripResult
}

func Analysis(response *http.Response) models.TripResponse {
	body, _ := ioutil.ReadAll(response.Body)
	var tripResponse models.TripResponse
	json.Unmarshal(body, &tripResponse)
	return tripResponse
}
