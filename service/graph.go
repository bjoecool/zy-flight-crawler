package service

import (
	"bytes"
	"fmt"
	"github.com/bclicn/color"
	"github.com/wcharczuk/go-chart"
	"github.com/zhangyuyu/zy-flight-crawler/models"
	"io/ioutil"
	"time"
)

func DrawGraph(tripResult models.TripResult) {
	fmt.Printf(color.Green("Today") + ": %v\n"+
		color.Green("Origin")+ " : %v\n"+
		color.Green("Destination")+ " : %v\n\n",
		time.Now().Format("2006-01-02 Mon"),
		tripResult.Origin,
		tripResult.Destination)

	details := tripResult.TripDetails
	length := len(details)

	prices := make([]float64, length)
	dates := make([]time.Time, length)

	for i := 0; i < length; i++ {
		fmt.Println(details[i])
		prices[i] = details[i].Price
		dates[i] = details[i].Start
	}

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:   20,
				Left:  20,
				Right: 50,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
			Name: "Price",
		},
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
			Name: "Time",
		},
		Series: []chart.Series{
			chart.TimeSeries{
				YValues: prices,
				XValues: dates,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	graph.Render(chart.PNG, buffer)
	ioutil.WriteFile("chart.png", buffer.Bytes(), 0644)
}
