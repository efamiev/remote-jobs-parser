package parser

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"remote-jobs-parser/internal/utils"
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
		log.Println("Error converting the pageCount to a string", err)

		return 0
	}

	return count
}

func Start(url string, page int) []string {
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

		ParseHH(results, client, url+"&page="+strconv.Itoa(page))
	}()

	wg.Wait()
	close(results)

	flatResult := make([]string, 0, 20)

	for x := range results {
		flatResult = append(flatResult, x...)
	}

	return flatResult
}
