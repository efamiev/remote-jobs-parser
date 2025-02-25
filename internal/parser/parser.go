package parser

import (
	"log"
	"net/http"
)

type ParserParams struct {
	Service string
	Url     string
}

type VacancyData struct {
	Id      string
	Company string
	Title   string
	Link    string
	Service string
}

func Start(params []ParserParams) []VacancyData {
	client := &http.Client{}

	hh := make(chan []VacancyData, 1)
	habr := make(chan []VacancyData, 1)

	for _, el := range params {
		switch el.Service {
		case "hh":
			go ParseHH(hh, client, el.Url)
		case "habr":
			go ParseHabr(habr, client, el.Url)
		default:
			log.Fatal("Unexpected Service name", el)
		}
	}

	results := []VacancyData{}

	for range make([]int, len(params)) {
		select {
		case jobsList := <-habr:
			results = append(results, jobsList...)

		case jobsList := <-hh:
			results = append(results, jobsList...)
		}
	}

	return results
}
