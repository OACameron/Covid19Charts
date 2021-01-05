package main

import (
	"flag"
	"log"
	"time"
)

type DayData struct {
	Date      string `json:"date"`
	NewCases  int    `json:"newCases"`
	NewDeaths int    `json:"newDeaths"`
}

type AreaStatistics struct {
	Area string
	Days []DayData
}

func main() {

	log.Print("Initialising")

	//dates := []string{"2021-01-02","2021-01-03","2021-01-04","2021-01-05"}

	area := flag.String("area", "Stockport","Area to get data about")
	numDays := flag.Int("days", 5, "Number of previous days to collect")

	flag.Parse()

	stats := AreaStatistics{
		Area: *area,
	}

	for *numDays >= 0{
		thisDate := time.Now().AddDate(0, 0, *numDays * -1)
		thisDay := DayData{
			Date: thisDate.Format("2006-01-02"),
		}
		thisDay.getDailyDeaths(*area)
		thisDay.getDailyCases(*area)

		stats.Days = append(stats.Days, thisDay)

		*numDays--
	}

	stats.generateLineChart(stats.Area)
}
