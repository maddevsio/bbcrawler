package bbcrawler

import (
	"github.com/syndtr/goleveldb/leveldb"
	"encoding/json"
	"fmt"
	"sync"
)

const (
	PATH_TO_DB = "./hacker_one_db"
)

var (
	WrongTypeOfDataError = func (msg string) error {return fmt.Errorf("WrongTypeError: %s", msg)}
)

type HackerOneStore struct {
	PathToDb string
	newRecords []HackerOneRecord
	sync.RWMutex
}

func (h HackerOneStore) Store(data interface{}) error {
	db, err := leveldb.OpenFile(h.PathToDb, nil)
	defer db.Close()

	if err != nil {
		return err
	}

	if response, ok := data.(HackerOneResponse); ok {
		for _, v := range response.Results {
			fmt.Print(".")
			jsonStr, _ := json.Marshal(v)
			if found, err := h.existsRecord(db, v); !found {
				h.Lock()
				h.newRecords = append(h.newRecords, v)
				h.Unlock()
				fmt.Print("+")
				err = db.Put([]byte("key"), jsonStr, nil)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (h HackerOneStore) GetNewRecords() interface{} {
	h.RLock()
	defer h.RUnlock()
	fmt.Println("Count new: ", len(h.newRecords))
	return h.newRecords
}

func (h HackerOneStore) existsRecord(db *leveldb.DB, data interface{}) (bool, error) {
	if rec, ok := data.(HackerOneRecord); ok {
		r, e := db.Has([]byte(rec.Handle), nil)
		return r, e
	}
	return false, WrongTypeOfDataError("Data should be HackerOneRecord type")
}

var HackerOneStoreInstance = &HackerOneStore{PathToDb: PATH_TO_DB, newRecords:make([]HackerOneRecord,0)}