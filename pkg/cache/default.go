package cache

import (
	"context"
	"sync"
	"time"
)

type DefaultCache struct {
	c sync.Map
}

func NewDefaultCache() *DefaultCache {
	return &DefaultCache{
		c: sync.Map{},
	}
}

type valueItem struct {
	Value      any
	CreateAt   int64
	Expiration time.Duration
}

func (c *DefaultCache) Set(ctx context.Context, key string, value any, expiration time.Duration, expireHook ExpiredHook) error {
	// createdAt 作为版本号，用于判断过期时的数据是否是和插入时的一样
	createdAt := time.Now().UnixMicro()
	v := valueItem{
		Value:      value,
		CreateAt:   createdAt,
		Expiration: expiration,
	}

	c.c.Store(key, &v)

	// 如果没有设置超时事件，则不处理超时
	if expiration <= 0 {
		return nil
	}

	go func() {
		<-time.After(expiration)
		val, ok, _ := c.GetRawValue(ctx, key)
		if !ok {
			return
		}

		if val == nil {
			return
		}

		// 如果数据是和插入时的一样，则删除
		if val.CreateAt == createdAt {
			c.Del(ctx, key)
			expireHook(ctx)
		}
	}()

	return nil
}

func (c *DefaultCache) GetRawValue(ctx context.Context, key string) (*valueItem, bool, error) {
	val, ok := c.c.Load(key)

	if !ok {
		return nil, false, nil
	}

	return val.(*valueItem), ok, nil
}

func (c *DefaultCache) Get(ctx context.Context, key string) (any, bool, error) {
	val, ok := c.c.Load(key)

	if !ok {
		return nil, false, nil
	}

	return val.(*valueItem).Value, ok, nil
}

func (c *DefaultCache) Del(ctx context.Context, key string) error {
	c.c.Delete(key)
	return nil
}
