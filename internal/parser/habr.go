package parser

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"remote-jobs-parser/internal/utils"
)

func ParseHarb(out chan<- []string, client *http.Client, url string) {
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

	fmt.Println("Количество вакансий на странице:", len(jobTitles))
	out <- jobTitles
}
