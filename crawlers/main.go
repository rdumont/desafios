package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("One argument required")
	}

	scraper := NewScraper()

	subs := strings.Split(os.Args[1], ";")
	for _, s := range subs {
		threads, err := scraper.Scrape(s)
		if err != nil {
			log.Fatal(err)
		}

		printThreads(threads)
	}
}

func printThreads(threads []Thread) {
	if len(threads) == 0 {
		return
	}

	for _, t := range threads {
		fmt.Printf("[%v] %v - %v\nPermalink: %v\nComments:  %v\n\n", t.Score, t.Subreddit, t.Title, t.Permalink, "https://www.reddit.com"+t.CommentsURL)
	}
}
