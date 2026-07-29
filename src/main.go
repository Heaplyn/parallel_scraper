package main

import (
	"fmt"
	"parallel_scraper/src/parallel_scraper"
	"time"
)

func main() {
	var results []parallel_scraper.Result
	var threads []chan parallel_scraper.Result

	times := 40
	start := time.Now()

	for i := 0; i < times; i++ {
		threads = append(threads, parallel_scraper.CheckUrl("https://google.com"))
		if len(results) < i-20 || i > times-20 {
			results = append(results, <-threads[len(results)+1])
		}
		if i%10 == 0 {
			time.Sleep(time.Millisecond * 500)
			fmt.Println("i is ", i)
		}

	}
	fmt.Println(results)
	fmt.Println(time.Since(start))
}
