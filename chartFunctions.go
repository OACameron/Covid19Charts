package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
)

func (data *AreaStatistics) generateLineChart(output string){

	outputFileName := output + ".html"

	lineChart := charts.NewLine()

	lineChart.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Daily Cases & Deaths",
			Subtitle: "Data for " + data.Area + ". Provided by gov.uk",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show: true,
		}),
	)

	var xAxisString []string

	casesLine := make([]opts.LineData, 0)
	deathsLine := make([]opts.LineData, 0)

	for _, day := range data.Days{
		xAxisString = append(xAxisString, day.Date)
		casesLine = append(casesLine, opts.LineData{
			Value: day.NewCases,
		})
		deathsLine = append(deathsLine, opts.LineData{
			Value: day.NewDeaths,
		})
	}

	lineChart.SetXAxis(xAxisString).AddSeries("New Cases", casesLine).AddSeries("New Deaths", deathsLine)

	f, _ := os.Create(outputFileName)
	lineChart.Render(f)

	return

}
