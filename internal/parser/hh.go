package parser

import (
	"log"
	"net/http"
	"strings"

	"remote-jobs-parser/internal/utils"

	"github.com/PuerkitoBio/goquery"
)

func ParseHH(out chan<- []VacancyData, client *http.Client, url string) {
	req := utils.Request(url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	vacancies := goquery.Map(doc.Find("#a11y-main-content [class^='vacancy-card--']"), func(i int, s *goquery.Selection) VacancyData {
		vacancy := VacancyData{Service: "hh"}

		vacancy.Company = s.Find(`[class^='company-name-badges-container'] [data-qa='vacancy-serp__vacancy-employer-text']`).Text()
		vacancy.Title = s.Find(".bloko-header-section-2").Text()
		vacancy.Link = s.Find(".bloko-header-section-2 a").AttrOr("href", "")
		vacancy.Id = strings.Split(strings.Split(vacancy.Link, "/")[4], "?")[0]

		return vacancy
	})

	log.Println("Count of vacancies on the HH:", len(vacancies))

	out <- vacancies
}
