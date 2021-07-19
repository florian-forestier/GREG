package service

import (
	"sync"
)

var cache = []Item{}
var mutex = &sync.Mutex{}

func Append(item Item) {
	mutex.Lock()
	
	found := false
	
	for i, key := range(cache) {
		if key.Name == item.Name {
			found = true
			cache[i].Value = item.Value
		}
	}
	
	if !found {
		cache = append(cache, item)
	}

	mutex.Unlock()
}

func Reset() {
	mutex.Lock()
	
	cache = []Item{}
	
	mutex.Unlock()
}