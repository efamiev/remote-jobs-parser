package parsers

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"math/rand"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const itemsOnPage = 20

func request(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.AddCookie(
		&http.Cookie{
			Name:  "device_magritte_breakpoint",
			Value: "xxl",
		})

	return req
}

func parsePage(out chan <- []string, client *http.Client, url string) {
	req := request(url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	jobTitles := doc.Find(`[id="a11y-main-content"] [data-qa="serp-item__title-text"]`).Map(func(index int, item *goquery.Selection) string {
		return item.Text()
	})
	
	fmt.Println("Количество вакансий на странице:", len(jobTitles))
	out <- jobTitles
}

func getPageCount(client *http.Client, url string) int {
	req := request(url)

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

func progressBar() {
	start := 0

	for {
		time.Sleep(time.Second * 1)

		start = start + 1
		fmt.Println("Время обработки в секундах:", start)
	}
}

func Parse() {
	url := "https://hh.ru/search/vacancy?text=%22go%22&salary=&professional_role=96&ored_clusters=true&items_on_page=20&enable_snippets=true&hhtmFrom=vacancy_search_list&hhtmFromLabel=vacancy_search_line"

	client := &http.Client{}
	pageCount := getPageCount(client, url + "&page=0")

	results := make(chan []string, pageCount)
	var wg sync.WaitGroup
	
	go func() {
		for x := range results {
			fmt.Println("Название вакансий:", x)
		}
	}()

	for i := 0; i < pageCount; i++ {
		wg.Add(1)
		delay := time.Duration(rand.Intn(301) + 300) * time.Millisecond

		go func(n int) {
			parsePage(results, client, url + "&page=" + strconv.Itoa(n))

			wg.Done()
		}(i)

		time.Sleep(delay)
	}

	wg.Wait()
	close(results)
}

