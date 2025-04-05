package cache

import (
	"fmt"
	"sync"
	"time"
)

type CachedResponse struct {
	Body       []byte
	Header     map[string][]string
	Expiration time.Time
}

var cacheMutex = &sync.Mutex{}

var cache = make(map[string]CachedResponse)

const ExpireTime = 5 * time.Minute

func Get(key string) ([]byte, map[string][]string, bool) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	cachedResp, found := cache[key]
	if !found || time.Now().After(cachedResp.Expiration) {
		return nil, nil, false
	}
	fmt.Println("Cache hit: " + key)
	return cachedResp.Body, cachedResp.Header, true
}
func Set(key string, body []byte, Header map[string][]string) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	cache[key] = CachedResponse{Body: body, Expiration: time.Now().Add(ExpireTime), Header: Header}
}
