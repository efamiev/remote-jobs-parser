package test

import (
	"io"
	"log"
	"net/http"
	"os"

	"remote-jobs-parser/internal/utils"
)

func ReadHTML(name string) string {
	file, err := os.Open("/Users/famiev/projects/remote-jobs-parser/test/" + name)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed reading data from file: %s", err)
	}

	content := string(data)
	return content
}

func SaveHHPage() {
	url := "https://hh.ru/search/vacancy?text=%22go%22&salary=&professional_role=96&items_on_page=20"
	client := &http.Client{}

	req := utils.Request(url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	file, err := os.Create("test/hh-page.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveHabrPage() {
	url := "https://career.habr.com/vacancies?q=go&sort=date&type=all"
	client := &http.Client{}

	req := utils.Request(url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	file, err := os.Create("test/habr-page.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Fatal(err)
	}
}
