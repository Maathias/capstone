package api

import (
	"fmt"
	"time"

	"github.com/maathias/capstone/db"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metric_locations = promauto.NewCounter(prometheus.CounterOpts{
		Name: "capstone_log_location",
		Help: "The total number of saved location points",
	})
)

// save location to db
func LogLocation(long float64, lat float64, uname string) {
	timestamp := time.Now().Unix()
	userHash := fmt.Sprint("user:", uname)
	timeHash := fmt.Sprint("timestamps:", uname)

	db.SortedAdd(timeHash, timestamp, fmt.Sprint(timestamp))
	db.GeoAdd(userHash, long, lat, fmt.Sprint(timestamp))

	metric_locations.Inc()
}

// get location points in range
func GetLocations(uname string, start int64, end int64) []string {
	timeHash := fmt.Sprint("timestamps:", uname)

	timestamps := db.SortedRange(timeHash, start, end)

	return timestamps
}

// get distance between two points
func GetDistance(uname string, stampA string, stampB string) float64 {
	userHash := fmt.Sprint("user:", uname)

	distance := db.GeoDist(userHash, stampA, stampB)

	return distance
}

// get distance between n points
func GetDistanceTotal(uname string, stamps []string) float64 {
	userHash := fmt.Sprint("user:", uname)

	// loop over
	distances := make([]float64, len(stamps)-1)
	for i := 0; i < len(stamps)-1; i++ {
		distances[i] = db.GeoDist(userHash, stamps[i], stamps[i+1])
	}

	// sum
	var sum float64
	for _, d := range distances {
		sum += d
	}

	return sum
}

func GetDinstanceInTimeRange(uname string, start int64, end int64) float64 {
	timestamps := GetLocations(uname, start, end)

	return GetDistanceTotal(uname, timestamps)
}

func Run() {
	fmt.Println("Running api service")

	// TODO: gRPC
}
