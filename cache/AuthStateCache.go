package cache

type AuthStateCache interface {
	Cache(key, value string)
	CacheTimeOut(key, value string, ttl int64)
	Get(key string) *StateCache
	ContainsKey(key string) bool
	clearExpiredCache()
}
