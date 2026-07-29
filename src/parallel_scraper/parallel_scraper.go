package parallel_scraper

import (
	"net/http"
	"time"
)

// Result represents the outcome of checking a URL
type Result struct {
	URL        string
	StatusCode int
	Duration   time.Duration
	Err        error
}

func check_url(url string, channel chan Result) {
	start := time.Now()
	// 1. Make the request
	res, err := http.Get(url)
	if err != nil {
		// If it fails, send the error
		channel <- Result{URL: url, Err: err, Duration: time.Since(start)}
		return
	}
	defer res.Body.Close()
	// 2. If it succeeds, send the status code
	channel <- Result{
		URL:        url,
		StatusCode: res.StatusCode,
		Duration:   time.Since(start),
		Err:        nil,
	}
}
func CheckUrl(url string) chan Result {
	newResult := make(chan Result, 1)
	go check_url(url, newResult)
	return newResult
}
