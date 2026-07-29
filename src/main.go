package main

import (
	"fmt"
	"parallel_scraper/src/parallel_scraper"
	"time"
)

func main() {
	url := "https://google.com"
	count := 10 // Let's test with 10 requests

	fmt.Printf("Testing with %d requests to %s...\n\n", count, url)

	// ==========================================
	// 1. NON-PARALLEL (SEQUENTIAL) BENCHMARK
	// ==========================================
	seqStart := time.Now()
	var seqResults []parallel_scraper.Result

	for i := 0; i < count; i++ {
		// We wait for each request to finish before starting the next one
		resultCh := parallel_scraper.CheckUrl(url)
		result := <-resultCh
		seqResults = append(seqResults, result)
	}
	seqDuration := time.Since(seqStart)

	fmt.Printf("Sequential: Checked %d sites in %v\n", len(seqResults), seqDuration)

	// ==========================================
	// 2. PARALLEL BENCHMARK
	// ==========================================
	parStart := time.Now()
	var parResults []parallel_scraper.Result
	var channels []chan parallel_scraper.Result

	// Step A: Launch all requests concurrently (non-blocking)
	for i := 0; i < count; i++ {
		ch := parallel_scraper.CheckUrl(url)
		channels = append(channels, ch)
	}

	// Step B: Wait for and collect ALL results
	for _, ch := range channels {
		result := <-ch
		parResults = append(parResults, result)
	}
	parDuration := time.Since(parStart)

	fmt.Printf("Parallel:   Checked %d sites in %v\n", len(parResults), parDuration)

	// ==========================================
	// SPEEDUP CALCULATIONS
	// ==========================================
	speedup := float64(seqDuration) / float64(parDuration)
	fmt.Printf("\nParallel version is %.2fx faster!\n", speedup)
}
