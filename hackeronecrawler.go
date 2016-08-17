package bbcrawler

import (
	"sync"
	"fmt"
	"github.com/labstack/gommon/log"
)

var (
	wg sync.WaitGroup
)

const (
	HACKER_ONE_SEARCH_URL = "https://hackerone.com/programs/search"
)

type HackerOneCrawler struct {
	sync.RWMutex
	fetcher Fetcher
	reader Reader
	store Storer
	pages map[int]*HackerOneResponse
}

func (h *HackerOneCrawler) hackerOneCrawl(url string, queryParams map[string]string) {
	response, err := h.hackerOneCrawlPage(url, queryParams, 1)
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
		hackerOneQuery := h.makeQuery(i)
		urlparams = append(urlparams, hackerOneQuery)

		wg.Add(1)
		go h.hackerOneCrawlPage(
			HACKER_ONE_SEARCH_URL,
			hackerOneQuery, i)
	}
	return
}

func (h *HackerOneCrawler) hackerOneCrawlPage(url string, queryParams map[string]string, page int) (*HackerOneResponse, error) {
	defer wg.Done()
	data, err := h.fetcher.Fetch(url, queryParams)
	if err != nil {
		return nil, err
	}
	jsonResponse, err := h.reader.Read(data)
	if err != nil {
		return nil, err
	}

	response := jsonResponse.(HackerOneResponse)

	h.Lock()
	h.pages[page] = &response
	h.Unlock()

	return &response, nil
}

func (h *HackerOneCrawler) makeQuery(pageNum int) map[string]string {
	hackerOneQuery := make(map[string]string)
	hackerOneQuery["query"] = "bounties:yes" //"bounties:yes ibb:false"
	hackerOneQuery["sort"] = "published_at:descending"
	hackerOneQuery["page"] = fmt.Sprintf("%d", pageNum)
	return hackerOneQuery
}

func (h *HackerOneCrawler) Crawl() {
	wg.Add(1)
	h.hackerOneCrawl(
		HACKER_ONE_SEARCH_URL,
		h.makeQuery(1))
	wg.Wait()

	h.RLock()
	defer h.RUnlock()

	for i := 1; i < len(h.pages); i++{
		response := h.pages[i]
		h.store.Store(*response)
	}

	newRecords := h.store.GetNewRecords().([]HackerOneRecord)
	if len(newRecords) > 0 {
		log.Println("New records: ", len(newRecords))
		log.Println("Content: ", newRecords)
	} else {
		log.Println("No new records found")
	}
}

var HackerOneCrawlerInstance = HackerOneCrawler{
	fetcher:*HackerOneFetcherInstance,
	reader:*HackerOneParserInstance,
	store:*HackerOneStoreInstance,
	pages:make(map[int]*HackerOneResponse),
}