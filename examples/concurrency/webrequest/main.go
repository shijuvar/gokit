package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type webResponse struct {
	requestUrl, responseText string
	status                   int
	error
}

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func main() {
	urls := []string{
		"https://golang.org",
		"https://www.google.com",
		"https://twitter.com",
	}

	wg.Add(2)
	webResponses := make(chan webResponse, len(urls))
	go getWebResponses(urls, webResponses)
	go printWebResponses(webResponses)
	wg.Wait()
}

func getWebResponses(urls []string, webResponses chan<- webResponse) {
	defer wg.Done()
	defer close(webResponses)
	var wgUrlReq sync.WaitGroup
	wgUrlReq.Add(len(urls))
	for _, v := range urls {
		go func(url string) {
			defer wgUrlReq.Done()
			responseVal := webResponse{}
			responseVal.requestUrl = url
			response, err := http.Get(url)
			if err != nil {
				responseVal.error = err
				responseVal.status = http.StatusInternalServerError
			} else {
				defer response.Body.Close()
				body, err := ioutil.ReadAll(response.Body)
				if err == nil {
					responseVal.responseText = string(body)
				}
				responseVal.error = err
				responseVal.status = response.StatusCode
			}
			webResponses <- responseVal
		}(v)
	}
	wgUrlReq.Wait()
}

func printWebResponses(webResponses <-chan webResponse) {
	defer wg.Done()
	for v := range webResponses {
		fmt.Println(v.requestUrl)
		if v.error != nil {
			fmt.Println("Error:", v.error)
		}
		fmt.Println("HTTP Status Code:", v.status)
		fmt.Println("Response Text:\n", v.responseText)

	}
}
