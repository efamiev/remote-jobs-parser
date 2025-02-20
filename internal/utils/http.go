package utils

import (
	"log"
	"net/http"
)

func Request(url string) *http.Request {
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
