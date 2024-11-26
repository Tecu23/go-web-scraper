// Package main is the entry point of the application
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := flag.String("url", "", "The URL of the webpage to scrape (required)")
	tags := flag.String(
		"tags",
		"h1,p",
		"Comma-separated list of HTML tags to scrape (default: 'h1,p')",
	)
	output := flag.String(
		"output",
		"output.json",
		"File name to save the scraped data (default: 'output.json')",
	)
	flag.Parse()

	if *url == "" {
		log.Fatal("Error: URL is required. Use --url flag to specify it.")
	}

	fmt.Printf("Scraping URL: %s\n", *url)
	fmt.Printf("Scraping tags: %s\n", *tags)
	fmt.Printf("Output will be saved to: %s\n", *output)

	html := fetchHTML(*url)

	fmt.Println(html[:200])
}

func fetchHTML(url string) string {
	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching URL %s: %v", url, err)
	}

	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	return string(body)
}
