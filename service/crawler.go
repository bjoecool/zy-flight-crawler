package service

import (
	"encoding/json"
	"github.com/zhangyuyu/zy-flight-crawler/models"
	"io/ioutil"
	"net/http"
	"time"
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
		dateOut, _ := time.Parse("2006-01-02T15:04:05.000", day.DateOut)
		if len(day.Flights) > 0 && day.Flights[0].FaresLeft != 0 {
			tripDetails[i] = models.TripDetail{
				DateOut:      dateOut,
				Price:        day.Flights[0].RegularFare.Fares[0].Amount,
				Unit:         unit,
				FlightNumber: day.Flights[0].FlightNumber,
				Start:        day.Flights[0].TimeUTC[0],
				End:          day.Flights[0].TimeUTC[1],
			}
		} else {
			tripDetails[i] = models.TripDetail{
				DateOut: dateOut,
			}
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
