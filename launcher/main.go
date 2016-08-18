package main

import (
	"github.com/cyberlight/bbcrawler"
	"time"
)

func main() {
	go bbcrawler.HackerOneCrawlerInstance.Crawl()
	done := false
	for {
		select {
		case <-bbcrawler.HackerOneCrawlerInstance.Done:
			done = true
		case <-time.After(1*time.Minute):
			if done {
				done = false
				go bbcrawler.HackerOneCrawlerInstance.Crawl()
			}
		}
	}
}
