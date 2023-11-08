package cache

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type DefaultCache struct {
	c       sync.Map
	hookLok sync.RWMutex
	hooks   map[string]ExpiredHook
}

func NewDefaultCache() Cache {
	return &DefaultCache{
		c:     sync.Map{},
		hooks: make(map[string]ExpiredHook),
	}
}

type valueItem struct {
	Value      any
	CreateAt   int64
	Expiration time.Duration
}

func (c *DefaultCache) SetExpireHook(resource string, hook ExpiredHook) {
	c.hookLok.Lock()
	defer c.hookLok.Unlock()

	c.hooks[resource] = hook
}

func (c *DefaultCache) Set(ctx context.Context, resource, key string, value any, expiration time.Duration) error {
	// createdAt 作为版本号，用于判断过期时的数据是否是和插入时的一样
	createdAt := time.Now().UnixMicro()
	v := valueItem{
		Value:      value,
		CreateAt:   createdAt,
		Expiration: expiration,
	}

	c.c.Store(c.keyGen(resource, key), &v)

	// 如果没有设置超时事件，则不处理超时
	if expiration <= 0 {
		return nil
	}

	go func() {
		<-time.After(expiration)
		val, ok, _ := c.GetRawValue(ctx, resource, key)
		if !ok {
			return
		}

		if val == nil {
			return
		}

		// 如果数据是和插入时的一样，则删除
		if val.CreateAt == createdAt {
			c.Del(ctx, resource, key)
			if hook, ok := c.hooks[resource]; ok {
				hook(ctx, key)
			}
		}
	}()

	return nil
}

func (c *DefaultCache) GetRawValue(ctx context.Context, resource, key string) (*valueItem, bool, error) {
	val, ok := c.c.Load(c.keyGen(resource, key))

	if !ok {
		return nil, false, nil
	}

	return val.(*valueItem), ok, nil
}

func (c *DefaultCache) Get(ctx context.Context, resource, key string) (any, bool, error) {
	val, ok := c.c.Load(c.keyGen(resource, key))

	if !ok {
		return nil, false, nil
	}

	return val.(*valueItem).Value, ok, nil
}

func (c *DefaultCache) Del(ctx context.Context, resource, key string) error {
	c.c.Delete(c.keyGen(resource, key))
	return nil
}

func (c *DefaultCache) keyGen(resource, key string) string {
	return fmt.Sprintf("%s:%s", resource, key)
}
