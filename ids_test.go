package lib

import (
	"errors"
	"log"
	"sync"
	"testing"
)

type Ids struct {
	Lock sync.Mutex
	Ids  map[string]bool
}

func (ids *Ids) Exist(id string) error {
	ids.Lock.Lock()
	defer ids.Lock.Unlock()
	if ids.Ids == nil {
		ids.Ids = make(map[string]bool)
	}

	_, ok := ids.Ids[id]
	if !ok {
		ids.Ids[id] = true
	} else {
		return errors.New("duplicated!")
	}
	return nil
}

func TestUniqueIds(t *testing.T) {
	var ids Ids
	var wg sync.WaitGroup
	for i := 0; i <= 200; i++ {
		wg.Add(1)
		go func(ids *Ids, wg *sync.WaitGroup) {
			for j := 0; j <= 100000; j++ {
				id := NewShortId("a")
				err := ids.Exist(id)
				if err != nil {
					log.Println(id)
					log.Println(err)
				}
			}
			wg.Done()
		}(&ids, &wg)
	}
	wg.Wait()
	log.Println(len(ids.Ids))
}
