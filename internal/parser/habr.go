package parser

import (
	"log"
	"net/http"
	"strings"

	"remote-jobs-parser/internal/utils"

	"github.com/PuerkitoBio/goquery"
)

func ParseHabr(out chan<- []VacancyData, client *http.Client, url string) {
	defer close(out)

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

	vacancies := goquery.Map(doc.Find(".vacancy-card"), func(i int, s *goquery.Selection) VacancyData {
		vacancy := VacancyData{Service: "habr"}

		vacancy.Company = s.Find(".vacancy-card__company-title").Text()
		vacancy.Title = s.Find(".vacancy-card__title").Text()
		vacancy.Link = s.Find(".vacancy-card__title a").AttrOr("href", "")
		vacancy.Id = strings.Split(vacancy.Link, "/")[2]

		return vacancy
	})
	
	log.Println("Count of vacancies on the Habr:", len(vacancies))
	out <- vacancies
}
