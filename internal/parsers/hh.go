package parsers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"hh-scraper/internal/utils"
)

func Parse(out chan<- []string, client *http.Client, url string) {
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

	jobTitles := doc.Find(`[id="a11y-main-content"] [data-qa="serp-item__title-text"]`).Map(func(_ int, item *goquery.Selection) string {
		return item.Text()
	})
	
	fmt.Println("Количество вакансий на странице:", len(jobTitles))
	out <- jobTitles
}

