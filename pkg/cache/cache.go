package cache

import (
	"context"
	"time"
)

type ExpiredHook func(ctx context.Context, key string)

type Cache interface {
	SetExpireHook(resource string, hook ExpiredHook)
	Set(ctx context.Context, resource, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, resource, key string) (any, bool, error)
	Del(ctx context.Context, resource, key string) error
}
