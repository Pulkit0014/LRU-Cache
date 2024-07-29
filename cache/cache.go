package cache

import (
	"fmt"
	"lru-cache/structs"
	"lru-cache/websocket"
	"time"
)

var lruCache *structs.LRUCache

func InitCache(capacity int) {
	lruCache = structs.NewLRUCache(capacity)
}

func Get(key string) (interface{}, string, bool) {
	lruCache.Mutex.Lock()
	defer lruCache.Mutex.Unlock()

	if element, found := lruCache.Items[key]; found {
		lruCache.List.MoveToFront(element)
		item := element.Value.(*structs.CacheItem)
		if item.Expiration == nil || item.Expiration.After(time.Now()) {
			return item.Value, "", true
		}
		lruCache.List.Remove(element)
		delete(lruCache.Items, key)
		return nil, "Key has expired", false
	}
	return nil, "Key not found", false
}

// Inside cache.go
func Set(key string, value interface{}, duration time.Duration) {
	lruCache.Mutex.Lock()
	defer lruCache.Mutex.Unlock()

	if element, found := lruCache.Items[key]; found {
		lruCache.List.MoveToFront(element)
		element.Value.(*structs.CacheItem).Value = value
		if duration > 0 {
			expiration := time.Now().Add(duration)
			element.Value.(*structs.CacheItem).Expiration = &expiration
		}
	} else {
		item := &structs.CacheItem{Key: key, Value: value}
		if duration > 0 {
			expiration := time.Now().Add(duration)
			item.Expiration = &expiration
		}
		element := lruCache.List.PushFront(item)
		lruCache.Items[key] = element

		if lruCache.List.Len() > lruCache.Capacity {
			evict()
		}
	}

	// Broadcast to WebSocket clients
	websocket.BroadcastUpdate(map[string]interface{}{
		"type":  "set",
		"key":   key,
		"value": value,
	})
}

func Delete(key string) {
	lruCache.Mutex.Lock()
	defer lruCache.Mutex.Unlock()

	if element, found := lruCache.Items[key]; found {
		lruCache.List.Remove(element)
		delete(lruCache.Items, key)

		// Broadcast to WebSocket clients
		websocket.BroadcastUpdate(map[string]interface{}{
			"type": "delete",
			"key":  key,
		})
	}
}
func GetAll() map[string]interface{} {
	lruCache.Mutex.Lock()
	defer lruCache.Mutex.Unlock()

	result := make(map[string]interface{})
	for key, element := range lruCache.Items {
		item := element.Value.(*structs.CacheItem)
		if item.Expiration == nil || item.Expiration.After(time.Now()) {
			result[key] = item.Value
		}
	}
	fmt.Println("result is ", result)
	return result
}

func evict() {
	element := lruCache.List.Back()
	if element != nil {
		lruCache.List.Remove(element)
		item := element.Value.(*structs.CacheItem)
		delete(lruCache.Items, item.Key)
	}
}
