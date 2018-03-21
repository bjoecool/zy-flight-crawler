package service

import (
	"encoding/json"
	"github.com/zhangyuyu/zy-flight-crawler/models"
	"io/ioutil"
	"net/http"
)

func Crawl(url string, fetcher Fetcher) models.TripResult {
	httpResponse, _ := fetcher.Fetch(url)

	tripResponse := ToTripResponse(httpResponse)
	tripResult := ToTripResult(tripResponse)
	//bytes, _ := json.Marshal(tripResult)
	//fmt.Printf(string(bytes))

	return tripResult
}

func ToTripResult(tripResponse models.TripResponse) models.TripResult {
	unit := tripResponse.Currency
	dates := tripResponse.Trips[0].Dates
	days := len(dates)

	tripDetails := make([]models.TripDetail, days)
	for i := 0; i < days; i++ {
		day := dates[i]
		tripDetails[i] = models.TripDetail{
			Price:        day.Flights[0].RegularFare.Fares[0].Amount,
			Unit:         unit,
			FlightNumber: day.Flights[0].FlightNumber,
			Start:        day.Flights[0].TimeUTC[0],
			End:          day.Flights[0].TimeUTC[1],
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

func ToTripResponse(response *http.Response) models.TripResponse {
	body, _ := ioutil.ReadAll(response.Body)
	var tripResponse models.TripResponse
	json.Unmarshal(body, &tripResponse)
	return tripResponse
}
