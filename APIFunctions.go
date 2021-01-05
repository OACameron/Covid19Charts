package main

import (
	"encoding/json"
	"net/http"
)

const APIPath string = "https://api.coronavirus.data.gov.uk/v1/data?"

func (day *DayData) getDailyCases(area string) {

	filters := "areaType=utla;areaName=" + area + ";date=" + day.Date

	type responseStructure struct {
		Data []struct {
			Date     string `json:"date"`
			NewCases int    `json:"newCases"`
		} `json:"data"`
	}

	inboundStructure := "{" +
		"\"date\":\"date\"," +
		"\"newCases\":\"newCasesByPublishDate\"" +
		"}"

	endpoint := APIPath + "filters=" + filters + "&" + "structure=" + inboundStructure
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	var thisResponse responseStructure

	err = json.NewDecoder(resp.Body).Decode(&thisResponse)

	if err != nil {
		panic(err.Error())
	}

	day.NewCases = thisResponse.Data[0].NewCases

}

func (day *DayData) getDailyDeaths(area string) {
	filters := "areaType=utla;areaName=" + area + ";date=" + day.Date

	type responseStructure struct {
		Data []struct {
			Date     string `json:"date"`
			NewDeaths int    `json:"newDeaths"`
		} `json:"data"`
	}

	inboundStructure := "{" +
		"\"date\":\"date\"," +
		"\"newDeaths\":\"newDeathsByPublishDate\"" +
		"}"

	endpoint := APIPath + "filters=" + filters + "&" + "structure=" + inboundStructure
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	var thisResponse responseStructure

	err = json.NewDecoder(resp.Body).Decode(&thisResponse)

	if err != nil {
		panic(err.Error())
	}

	day.NewDeaths = thisResponse.Data[0].NewDeaths
}
