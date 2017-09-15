package main

import (
	"bysykkelCharts/bysykkel"
	"fmt"
	"github.com/wcharczuk/go-chart" //exposes "chart"
	"log"
	"net/http"
	"os"
	"strconv"
)

func drawChart(res http.ResponseWriter, req *http.Request) {

	day := req.URL.Path[1:]
	trips := bysykkel.GetTripsConfig("data/tripsJune2017.json")
	points := bysykkel.GetPoints(trips, day)

	var values []chart.Value
	for i := 1; i < 25; i++ {
		value := chart.Value{
			Value: points[i],
			Label: strconv.Itoa(i),
		}
		values = append(values, value)
	}

	sbc := chart.BarChart{
		Title: "Number of bikes on Schous plass on June " + day,
		TitleStyle: chart.Style{
			Show: true,
		},
		Height:   600,
		BarWidth: 60,
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Bars: values,
	}

	res.Header().Set("Content-Type", "image/png")
	err := sbc.Render(chart.PNG, res)
	if err != nil {
		fmt.Printf("Error rendering chart: %v\n", err)
	}
}

func port() string {
	if len(os.Getenv("PORT")) > 0 {
		return os.Getenv("PORT")
	}
	return "8080"
}

func main() {

	listenPort := fmt.Sprintf(":%s", port())
	fmt.Printf("Listening on %s\n", listenPort)
	http.HandleFunc("/", drawChart)
	log.Fatal(http.ListenAndServe(listenPort, nil))

}
