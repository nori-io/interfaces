package cache

import (
	"time"

	e "errors"

	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	"github.com/nori-io/common/v4/pkg/errors"
)

const (
	CacheInterface meta.Interface = "nori/cache/Cache@1.0.0"
)

var (
	CacheKeyNotFound = e.New("CacheKeyNotFound")
)

type Cache interface {
	Clear() error
	Delete(key []byte) error
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte, ttl time.Duration) error
}

func GetCache(r registry.Registry) (Cache, error) {
	instance, err := r.Interface(CacheInterface)
	if err != nil {
		return nil, err
	}
	i, ok := instance.(Cache)
	if !ok {
		return nil, errors.InterfaceAssertError{
			Interface: CacheInterface,
		}
	}
	return i, nil
}
