package structs

import (
	"container/list"
	"sync"
	"time"
)

type CacheItem struct {
	Key        string
	Value      interface{}
	Expiration *time.Time
}

type LRUCache struct {
	Capacity int
	Items    map[string]*list.Element
	order    []string
	List     *list.List
	Mutex    sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity: capacity,
		Items:    make(map[string]*list.Element),
		List:     list.New(),
	}
}
