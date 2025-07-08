package mangadex

import (
	"context"
	"encoding/json"
	goCache "github.com/patrickmn/go-cache"
	redis "github.com/redis/go-redis/v9"
	"log/slog"
	"sync"
)

// Cache is a simple key-value store for caching responses.
type Cache interface {
	Get(key string) ([]byte, bool)
	Set(key string, value interface{})
}

// MemCache is an in-memory implementation of Cache.
type MemCache struct{ sync.Map }

func (m *MemCache) Get(key string) (interface{}, bool) { return m.Load(key) }
func (m *MemCache) Set(key string, value interface{})  { m.Store(key, value) }

// RedisCache is a Redis-backed implementation of Cache.
type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisCache creates a RedisCache.
// addr is e.g. "localhost:6379", password="" if none, db=0..
func NewRedisCache(addr, password string, db int) Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisCache{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (r *RedisCache) Get(key string) ([]byte, bool) {
	data, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil || err != nil {
		return nil, false
	}
	// try JSON unmarshal, else return raw string
	return []byte(data), true
	//var v interface{}
	//if json.Unmarshal([]byte(data), &v) == nil {
	//	return v, true
	//}
	//return data, true
}

func (r *RedisCache) Set(key string, value interface{}) {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return
		}
		data = b
	}
	// no expiration; change `0` to a time.Duration if you'd like TTL
	r.client.Set(r.ctx, key, data, 0)
}

// HybridCache tries the local cache first, then falls back to Redis.
type HybridCache struct {
	Local  Cache
	Remote Cache
}

func NewHybridCache(local Cache, remote Cache) Cache {
	return &HybridCache{Local: local, Remote: remote}
}

func (h *HybridCache) Get(key string) ([]byte, bool) {
	if v, ok := h.Local.Get(key); ok {
		return v, true
	}
	if v, ok := h.Remote.Get(key); ok {
		// warm up local cache
		h.Local.Set(key, v)
		return v, true
	}
	return nil, false
}

func (h *HybridCache) Set(key string, value interface{}) {
	h.Local.Set(key, value)
	h.Remote.Set(key, value)
}

// HybridCache tries the local cache first, then falls back to Redis.
type LocalCache struct {
	Local *goCache.Cache
}

func NewLocalCache(local *goCache.Cache) Cache {
	return &LocalCache{Local: local}
}

func (h *LocalCache) Get(key string) ([]byte, bool) {
	if v, ok := h.Local.Get(key); ok {
		return v.([]byte), true
	}

	return nil, false
}

func (h *LocalCache) Set(key string, value interface{}) {
	data, err := json.Marshal(value)
	if err != nil {
		slog.Error("Error marshalling value to json", "err", err)
		return
	}
	h.Local.Set(key, data, goCache.DefaultExpiration)
}
