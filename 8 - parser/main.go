package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var websiteUrl string = "http://httpbin.org/get"

func main() {
	res, err := http.Get(websiteUrl)
	if err != nil {
		log.Fatalf("Error query: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Request code not equal 200")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error body: %v", err)
	}

	fmt.Println(string(body))
}
