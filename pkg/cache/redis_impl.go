package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/AH-dark/apollo/pkg/log"
	"github.com/go-redis/redis/v8"
	"github.com/samber/lo"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

var _ Driver = (*RedisCache)(nil)

func NewRedisCache(network, host string, port int, password string, db int) Driver {
	return &RedisCache{
		client: redis.NewClient(&redis.Options{
			Network:     network,
			Addr:        fmt.Sprintf("%s:%d", host, port),
			Password:    password,
			DB:          db,
			IdleTimeout: time.Second * 10,
		}),
		ctx: context.Background(),
	}
}

type redisItem struct {
	Value interface{}
}

func serializer(value interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	storeValue := redisItem{
		Value: value,
	}
	err := enc.Encode(storeValue)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func deserializer(value []byte) (interface{}, error) {
	var res redisItem
	buffer := bytes.NewReader(value)
	dec := gob.NewDecoder(buffer)
	err := dec.Decode(&res)
	if err != nil {
		return nil, err
	}
	return res.Value, nil
}

// Set 存储值
func (store *RedisCache) Set(key string, value interface{}, ttl int) error {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().WithError(err).Error("Redis 关闭连接错误")
		}
	}(rc)

	serialized, err := serializer(value)
	if err != nil {
		return err
	}

	if ttl > 0 {
		err = rc.SetEX(store.ctx, key, serialized, time.Second*time.Duration(ttl)).Err()
	} else {
		err = rc.Set(store.ctx, key, serialized, 0).Err()
	}

	if err != nil {
		return err
	}
	return nil

}

// Get 取值
func (store *RedisCache) Get(key string) (interface{}, bool) {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().WithError(err).Error("Redis 关闭连接错误")
		}
	}(rc)

	v, err := rc.Get(store.ctx, key).Bytes()
	if err != nil || v == nil {
		return nil, false
	}

	finalValue, err := deserializer(v)
	if err != nil {
		return nil, false
	}

	return finalValue, true

}

// Gets 批量取值
func (store *RedisCache) Gets(keys []string, prefix string) (map[string]interface{}, []string) {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().WithError(err).Error("Redis 关闭连接错误")
		}
	}(rc)

	v, err := rc.MGet(store.ctx, lo.Map(keys, func(key string, index int) string {
		return prefix + key
	})...).Result()
	if err != nil {
		return nil, keys
	}

	res := make(map[string]interface{})
	missed := make([]string, 0, len(keys))

	for key, value := range v {
		decoded, err := deserializer([]byte(value.(string)))
		if err != nil || decoded == nil {
			missed = append(missed, keys[key])
		} else {
			res[keys[key]] = decoded
		}
	}
	return res, missed
}

// Sets 批量设置值
func (store *RedisCache) Sets(values map[string]interface{}, prefix string) error {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().WithError(err).Error("Redis 关闭连接错误")
		}
	}(rc)
	var setValues = make(map[string]interface{})

	// 编码待设置值
	for key, value := range values {
		serialized, err := serializer(value)
		if err != nil {
			return err
		}
		setValues[prefix+key] = serialized
	}

	_, err := rc.MSet(store.ctx, setValues).Result()
	if err != nil {
		return err
	}
	return nil

}

// Delete 删除给定的键
func (store *RedisCache) Delete(key string) error {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().WithError(err).Error("Redis 关闭连接错误")
		}
	}(rc)

	_, err := rc.Del(store.ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

// Deletes 批量删除给定的键
func (store *RedisCache) Deletes(keys []string, prefix string) error {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().WithError(err).Error("Redis 关闭连接错误")
		}
	}(rc)

	// 处理前缀
	keys = lo.Map[string, string](keys, func(key string, index int) string {
		return prefix + key
	})

	_, err := rc.Del(store.ctx, keys...).Result()
	if err != nil {
		return err
	}
	return nil
}

// DeleteAll 批量所有键
func (store *RedisCache) DeleteAll() error {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().WithError(err).Error("Redis 关闭连接错误")
		}
	}(rc)

	_, err := rc.FlushDB(store.ctx).Result()

	return err
}
