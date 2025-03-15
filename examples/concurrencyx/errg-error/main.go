package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type response struct {
	url    string
	status int
}

func main() {
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"https://shijuvar.medium.com",
	}
	responses, err := getWebResponse(urls)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, v := range responses {
		fmt.Printf("URL:%s HTTP Stautus Code:%d\n", v.url, v.status)
	}
}

func getWebResponse(urls []string) ([]response, error) {
	// A Group is a collection of goroutines working on subtasks that are part of the same overall task.
	g := new(errgroup.Group)
	responses := make([]response, len(urls))

	for i, url := range urls {
		// Launch a goroutine to fetch the URL.
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				responseData := response{
					url:    url,
					status: resp.StatusCode,
				}
				responses[i] = responseData
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	// Wait blocks until all function calls from the Go method have returned,
	// then returns the first non-nil error (if any) from them.
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return responses, nil
}
