package bysykkel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

// TripsConfig contains all the trips in a month
type TripsConfig struct {
	Trips []Trip `json:"trips"`
}

// Trip contains the trip information
type Trip struct {
	StartStationID int    `json:"start_station_id"`
	StartTime      string `json:"start_time"`
	EndStationID   int    `json:"end_station_id"`
	EndTime        string `json:"end_time"`
}

// GetTripsConfig gets all the trips in a month from a JSON file
func GetTripsConfig(filepath string) TripsConfig {

	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	var trips TripsConfig
	err = json.Unmarshal(raw, &trips)
	if err != nil {
		panic(err)
	}

	return trips

}

// GetPoints gets all the points for the barchart
func GetPoints(trips TripsConfig, dayString string) map[int]float64 {

	points := make(map[int]float64)
	for i := 1; i < 25; i++ {
		points[i] = 0
	}
	for _, trip := range trips.Trips {

		t, err := time.Parse("2006-01-02 15:04:05 -0700", trip.StartTime)
		if err != nil {
			fmt.Println("parse error", err.Error())
		}

		day, err := strconv.Atoi(dayString)
		if err != nil {
			panic(err)
		}
		if t.Day() == day && trip.StartStationID == 229 {
			points[t.Hour()]++
		}
	}

	return points

}
