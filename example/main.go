package main

import (
	"github.com/cyberlight/bbcrawler"
	"time"
)

func main() {
	config := &bbcrawler.HackerOneCrawlerConfig{
		SearchUrl:     "URL",
		PathToLocalDb: "DB",
		FireBaseUrl:   "https://project_name.firebaseio.com",
		FireBaseToken: "TOKEN",
	}
	crawler := bbcrawler.NewHackerOneCrowler(config)

	go crawler.Crawl()
	done := false
	for {
		select {
		case <-crawler.Done:
			crawler.ClearNewRecords()
			done = true
		case <-time.After(1 * time.Minute):
			if done {
				done = false
				go crawler.Crawl()
			}
		}
	}
}
