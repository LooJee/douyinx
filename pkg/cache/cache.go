package cache

import (
	"context"
	"time"
)

type ExpiredHook func(context.Context)

type Cache interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration, expireHook ExpiredHook) error
	Get(ctx context.Context, key string) (any, bool, error)
	Del(ctx context.Context, key string) error
}
