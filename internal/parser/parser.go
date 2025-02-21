package parser

import (
	"log"
	"net/http"
)

// const itemsOnPage = 20

func Start(url string, page int) []string {
	client := &http.Client{}

	hh := make(chan []string)
	habr := make(chan []string)
	
	go ParseHH(hh, client, "https://hh.ru/search/vacancy?text=%22go%22&salary=&professional_role=96&items_on_page=20")
	go ParseHabr(habr, client, "https://career.habr.com/vacancies?page=1&type=all")

	results := []string{}

	for {
		select {
		case jobsList, ok := <-hh:
			if !ok {
				hh = nil
				break
			}
			results = append(results, jobsList...)

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
