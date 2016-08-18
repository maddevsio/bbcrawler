package bbcrawler

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"sync"
	"time"
)

var (
	PATH_TO_DB = "hacker_one.db"
)

type HackerOneStore struct {
	PathToDb   string
	newRecords []HackerOneRecord
	sync.RWMutex
}

func (h *HackerOneStore) Store(data interface{}) error {
	h.Lock()
	defer h.Unlock()
	db, err := bolt.Open(h.PathToDb, 0600, &bolt.Options{Timeout: 5 * time.Second})
	defer db.Close()

	if err != nil {
		return err
	}

	if response, ok := data.(HackerOneResponse); ok {
		for _, v := range response.Results {
			fmt.Print(".")
			jsonStr, _ := json.Marshal(v)
			db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte("All"))
				if err != nil {
					return fmt.Errorf("create All bucket: %s", err)
				}
				bn, err := tx.CreateBucketIfNotExists([]byte("New"))
				if err != nil {
					return fmt.Errorf("create New bucket: %s", err)
				}
				if b.Get([]byte(v.Handle)) == nil {
					fmt.Print("+")
					h.newRecords = append(h.newRecords, v)
					err = b.Put([]byte(v.Handle), jsonStr)
					if err != nil {
						return err
					}
					return bn.Put([]byte(v.Handle), jsonStr)
				}
				return nil
			})
		}
	}
	return nil
}

func (h HackerOneStore) GetNewRecords() interface{} {
	h.RLock()
	defer h.RUnlock()
	return h.newRecords
}

var HackerOneStoreInstance = &HackerOneStore{PathToDb: PATH_TO_DB, newRecords: make([]HackerOneRecord, 0)}
