package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range websites {
		checkLink(link)
	}
}

func checkLink(link string) {
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(link, "status: ", resp.Status)
}
