package service

import (
	"github.com/zhangyuyu/zy-flight-crawler/models"
	"fmt"
	"github.com/bclicn/color"
)

func DrawGraph(tripResult models.TripResult) {
	fmt.Printf(color.Green("Origin") + " : %v(%v)\n"+
		color.Green("Destination")+ " : %v(%v)\n\n",
		tripResult.Origin, tripResult.OriginName, tripResult.Destination, tripResult.DestinationName)

	details := tripResult.TripDetails
	for i := 0; i < len(details); i++ {
		fmt.Println(details[i])
	}
}
