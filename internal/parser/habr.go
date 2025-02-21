package parser

import (
	"log"
	"net/http"

	"remote-jobs-parser/internal/utils"

	"github.com/PuerkitoBio/goquery"
)

func ParseHabr(out chan<- []string, client *http.Client, url string) {
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

	jobTitles :=
		doc.Find(`.vacancy-card .vacancy-card__title`).Map(func(_ int, item *goquery.Selection) string {
			return item.Text()
		})

	log.Println("Количество вакансий на странице habr:", len(jobTitles))
	out <- jobTitles
}
