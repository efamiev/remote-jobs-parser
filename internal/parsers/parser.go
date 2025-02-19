package parsers

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"strconv"
	"hh-scraper/internal/utils"

	"github.com/PuerkitoBio/goquery"
)

const itemsOnPage = 20

func getPageCount(client *http.Client, url string) int {
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

	pageCount := doc.Find(`[class^="magritte-number-pages-wrapper"] li`).Eq(-2).Text()

	fmt.Println("Количество страниц:", pageCount)

	count, err := strconv.Atoi(pageCount)
	if err != nil {
		log.Fatal(err)
	}
	
	return count
}

func Start(page int) []string {
	url := "https://hh.ru/search/vacancy?text=%22go%22&salary=&professional_role=96&items_on_page=" + strconv.Itoa(itemsOnPage)

	client := &http.Client{}
	pageCount := getPageCount(client, url)

	if pageCount < page {
		fmt.Println("The page does not exist")

		return []string{}
	}
	
	results := make(chan []string, pageCount)
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		defer wg.Done()

		Parse(results, client, url + "&page=" + strconv.Itoa(page))
	}()

	wg.Wait()
	close(results)
	
	flatResult := make([]string, itemsOnPage)

	for x := range results {
		flatResult = append(flatResult, x...)
	}

	return flatResult
}

