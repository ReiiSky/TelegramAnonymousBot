package module

import "sync"

// SafeCache ..
type SafeCache struct {
	Record map[string]string
	sync.Mutex
}

var userCache = SafeCache{
	Record: make(map[string]string),
}

// Read record from cache
func (cache *SafeCache) Read(k string) (string, bool) {
	cache.Lock()
	x, y := cache.Record[k]
	cache.Unlock()
	return x, y
}

// Insert record from cache
func (cache *SafeCache) Insert(k, v string) {
	cache.Lock()
	cache.Record[k] = v
	cache.Unlock()
}

// Delete record from cache
func (cache *SafeCache) Delete(k string) {
	cache.Lock()
	delete(cache.Record, k)
	cache.Unlock()
}
