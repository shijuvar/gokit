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
		"https://www.xyzwitter.com/",
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
	g := new(errgroup.Group)
	responses := make([]response, len(urls))

	for i, url := range urls {
		// Launch a goroutine to fetch the URL.
		i := i
		url := url // https://golang.org/doc/faq#closures_and_goroutines
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
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return responses, nil
}
