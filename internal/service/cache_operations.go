package service

import (
	"github.com/Artheriom/GREG/internal/structures"
	"sync"
)

var cache []structures.Item
var mutex = &sync.Mutex{}

func Append(item structures.Item) {
	mutex.Lock()

	found := false

	for i, key := range cache {
		if key.Name == item.Name {
			found = true
			cache[i].Value = item.Value
			break
		}
	}

	if !found {
		cache = append(cache, item)
	}

	mutex.Unlock()
}

func Reset() {
	mutex.Lock()

	cache = []structures.Item{}

	mutex.Unlock()
}
