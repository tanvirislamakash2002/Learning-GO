package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Url    string
	Status string
	Err    error
}

func checkWebsiteUrl(url string, ch chan Result) {
	res, err := http.Get(url)
	if err != nil {

		ch <- Result{
			Url:    url,
			Status: "is down",
			Err:    err,
		}
		return
	}

	ch <- Result{
		Url:    url,
		Status: "is up and running",
		Err:    nil,
	}
	defer res.Body.Close()
}

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://wrong-url.com",
	}
	ch := make(chan Result)

	start := time.Now()

	for _, url := range urls {
		go checkWebsiteUrl(url, ch)
	}

	for range urls {
		result := <-ch

		if result.Err != nil {
			fmt.Println(result.Url, result.Status, "Error:", result.Err)
			continue
		}

		fmt.Println(result)
	}

	fmt.Println("time taken", time.Since(start))
}
