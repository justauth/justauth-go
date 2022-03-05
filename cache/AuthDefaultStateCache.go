package cache

import (
	"sync"
	"time"
)

//AuthDefaultStateCache  default StateCache
type AuthDefaultStateCache struct {
	stateCache sync.Map
}

var timeout = time.Second * 180

type StateCache struct {
	state  string
	expire int64
}

func NewAuthDefaultStateCache() *AuthDefaultStateCache {
	defaultStateCache := &AuthDefaultStateCache{}
	go func() {
		ticker := time.NewTicker(timeout)
		defer ticker.Stop()
		for range ticker.C {
			defaultStateCache.clearExpiredCache()
		}
	}()
	return defaultStateCache
}

func (defaultStateCache *StateCache) isExpired() bool {
	return time.Now().Unix() > defaultStateCache.expire
}

func (defaultStateCache *AuthDefaultStateCache) Cache(key, value string) {
	defaultStateCache.stateCache.Store(key, &StateCache{
		state:  value,
		expire: time.Now().Unix() + int64(timeout.Seconds()),
	})
}

func (defaultStateCache *AuthDefaultStateCache) CacheTimeOut(key, value string, ttl int64) {
	defaultStateCache.stateCache.Store(key, &StateCache{
		state:  value,
		expire: time.Now().Unix() + ttl,
	})
}

func (defaultStateCache *AuthDefaultStateCache) Get(key string) *StateCache {
	value, ok := defaultStateCache.stateCache.Load(key)
	if !ok {
		return nil
	}
	return value.(*StateCache)
}

func (defaultStateCache *AuthDefaultStateCache) ContainsKey(key string) bool {
	_, ok := defaultStateCache.stateCache.Load(key)
	if !ok {
		return false
	}
	return true
}

func (defaultStateCache *AuthDefaultStateCache) clearExpiredCache() {
	defaultStateCache.stateCache.Range(func(key, value interface{}) bool {
		if value.(*StateCache).isExpired() {
			defaultStateCache.stateCache.Delete(key)
		}
		return true
	})
}
