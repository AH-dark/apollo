package cache

import (
	"fmt"
	"github.com/AH-dark/apollo/config"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/samber/lo"
	"strings"
)

var Store Driver

func Init() {
	if config.Redis.Host != "" {
		Store = NewRedisCache(config.Redis.Network, config.Redis.Host, config.Redis.Port, config.Redis.Password, config.Redis.DB)
	} else {
		Store = NewMemoryCache()
	}
}

func Get(key string) (interface{}, bool) {
	log.Log().Debugf("get cache key: %s", key)
	return Store.Get(key)
}

func Gets(keys []string, prefix string) (map[string]interface{}, []string) {
	go log.Log().Debugf("get cache keys: %s", strings.Join(lo.Map(keys, func(key string, i int) string {
		return prefix + key
	}), ","))

	return Store.Gets(keys, prefix)
}

func Set(key string, value interface{}, ttl int) error {
	log.Log().Debugf("set cache key: %s=%v, ttl: %d", key, value, ttl)
	return Store.Set(key, value, ttl)
}

func Sets(values map[string]interface{}, prefix string) error {
	go log.Log().Debugf("set cache keys: %s", strings.Join(lo.MapToSlice(values, func(key string, value interface{}) string {
		return fmt.Sprintf("%s=%v", prefix+key, value)
	}), ","))

	return Store.Sets(values, prefix)
}

func Delete(key string) error {
	log.Log().Debugf("delete cache key: %s", key)
	return Store.Delete(key)
}

func Deletes(keys []string, prefix string) error {
	go log.Log().Debugf("delete cache keys: %s", strings.Join(lo.Map(keys, func(key string, i int) string {
		return prefix + key
	}), ","))

	return Store.Deletes(keys, prefix)
}
