package parser

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"remote-jobs-parser/internal/utils"
)

func ParseHH(out chan<- []string, client *http.Client, url string) {
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
		doc.Find(`[id="a11y-main-content"] [data-qa="serp-item__title-text"]`).Map(func(_ int, item *goquery.Selection) string {
			return item.Text()
		})

	log.Println("Количество вакансий на странице:", len(jobTitles))
	out <- jobTitles
}
