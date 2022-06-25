package cache

import (
	"github.com/sirupsen/logrus"

	"sync"
)

type CallbacksCache struct {
	lock      sync.RWMutex
	callbacks map[string]*Callback
}

type Callback struct {
	Response interface{}
	Err      chan error
}

func (cache *CallbacksCache) Append(id string, callback *Callback) {
	defer cache.lock.Unlock()
	cache.lock.Lock()
	if cache.callbacks[id] != nil {
		logrus.Warnf("Callback for request with id [%s] already present. Overrides",
			id)
	}
	cache.callbacks[id] = callback
}

func (cache *CallbacksCache) Poll(id string) *Callback {
	defer cache.lock.RUnlock()
	cache.lock.RLock()
	callback := cache.callbacks[id]
	delete(cache.callbacks, id)
	return callback
}
