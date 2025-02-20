package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"hh-scraper/internal/utils"
)

func main() {
	url := "https://hh.ru/search/vacancy?text=%22go%22&salary=&professional_role=96&items_on_page=20"
	client := &http.Client{}

	req := utils.Request(url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	file, err := os.Create("./test/hh-page.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Fatal(err)
	}
}
