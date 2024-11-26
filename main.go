// Package main is the entry point of the application
package main

import (
	"flag"
	"fmt"
	"log"
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
}
