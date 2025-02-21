package main

import (
	"fmt"
	"os"
	"strconv"
	
	"remote-jobs-parser/internal/parsers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [arguments]")
		return
	}

	page, err := strconv.Atoi(os.Args[1:][0])
	if err != nil {
		fmt.Println("Pass page number")
		return
	}

	url := "https://hh.ru/search/vacancy?text=%22go%22&salary=&professional_role=96&items_on_page=20"

	fmt.Println(parsers.Start(url, page-1))
}
