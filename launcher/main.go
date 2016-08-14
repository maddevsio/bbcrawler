package main

import (
	"encoding/json"
	"fmt"
	"github.com/cyberlight/bbcrawler"
	"runtime"
	"sync"
)

var (
	wg sync.WaitGroup
)

func HackerOneCrawl(url string, queryParams map[string]string, fetcher bbcrawler.Fetcher, reader bbcrawler.Reader) {
	response, err := HackerOneCrawlPage(url, queryParams, fetcher, reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	pages := response.Total / response.Limit
	if response.Total%response.Limit > 0 {
		pages += 1
	}

	var urlparams []map[string]string
	for i := 2; i <= pages; i++ {
		hackerOneQuery := make(map[string]string)
		hackerOneQuery["query"] = "bounties:yes" //"bounties:yes ibb:false"
		hackerOneQuery["sort"] = "published_at:descending"
		hackerOneQuery["page"] = fmt.Sprintf("%d", i)
		urlparams = append(urlparams, hackerOneQuery)

		wg.Add(1)
		go HackerOneCrawlPage(
			"https://hackerone.com/programs/search",
			hackerOneQuery,
			bbcrawler.HackerOneFetcher,
			bbcrawler.HackerOneParser)
	}
	return
}

func HackerOneCrawlPage(url string, queryParams map[string]string, fetcher bbcrawler.Fetcher, reader bbcrawler.Reader) (*bbcrawler.HackerOneResponse, error) {
	defer wg.Done()
	data, err := fetcher.Fetch(url, queryParams)
	if err != nil {
		return nil, err
	}
	jsonResponse, err := reader.Read(data)
	if err != nil {
		return nil, err
	}

	response := jsonResponse.(bbcrawler.HackerOneResponse)

	var NewLineString string = "\n"
	if runtime.GOOS == "windows" {
		NewLineString = "\r\n"
	}

	info, _ := json.MarshalIndent(response, NewLineString, "    ")
	fmt.Print(string(info))
	return &response, nil
}

func main() {
	hackerOneQuery := make(map[string]string, 0)
	hackerOneQuery["query"] = "bounties:yes" //"bounties:yes ibb:false"
	hackerOneQuery["sort"] = "published_at:descending"
	hackerOneQuery["page"] = "1"

	wg.Add(1)
	HackerOneCrawl(
		"https://hackerone.com/programs/search",
		hackerOneQuery,
		bbcrawler.HackerOneFetcher,
		bbcrawler.HackerOneParser)
	wg.Wait()
}
