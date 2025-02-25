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

	hh := make(chan []string)
	habr := make(chan []VacancyData)

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

	for {
		select {
		// case jobsList, ok := <-hh:
		// 	if !ok {
		// 		hh = nil
		// 		break
		// 	}
		// 	results = append(results, jobsList...)

		case jobsList, ok := <-habr:
			if !ok {
				habr = nil
				break
			}
			results = append(results, jobsList...)
		}

		if hh == nil && habr == nil {
			log.Println("Final jobs count", len(results))
			return results
		}
	}
}
