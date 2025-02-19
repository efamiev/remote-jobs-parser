package main

import (
	"fmt"
	"hh-scraper/internal/parsers"
	"os"
	"strconv"
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

	fmt.Println(parsers.Start(page - 1))
}

