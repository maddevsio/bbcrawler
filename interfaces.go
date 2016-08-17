package bbcrawler

type Fetcher interface {
	Fetch(url string, params map[string]string) ([]byte, error)
}

type Reader interface {
	Read(data []byte) (interface{}, error)
}

type Storer interface {
	Store(data interface{}) error
	GetNewRecords() interface{}
}