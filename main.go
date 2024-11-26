// Package main is the entry point of the application
package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

	// Print Args
	fmt.Printf("Scraping URL: %s\n", *url)
	fmt.Printf("Scraping tags: %s\n", *tags)
	fmt.Printf("Output will be saved to: %s\n", *output)

	// Fetch HTML
	html := fetchHTML(*url)

	// Parse HTML
	results := parseHTML(html, *tags)

	// Save Output
	saveAsJSON(results, *output)

	fmt.Printf("Scraped data saved to %s\n", output)
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

func parseHTML(html, tags string) []string {
	var results []string

	// Load HTML into GoQuery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}

	// Extract specified tags
	tagList := strings.Split(tags, ",")
	for _, tag := range tagList {
		doc.Find(tag).Each(func(_ int, s *goquery.Selection) {
			results = append(results, s.Text())
		})
	}

	return results
}

func saveAsJSON(data []string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}
}

func saveAsCSV(data []string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, line := range data {
		if err := writer.Write([]string{line}); err != nil {
			log.Fatalf("Error writing CSV: %v", err)
		}
	}
}
