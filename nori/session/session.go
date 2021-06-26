package session

import (
	"time"

	"github.com/nori-io/common/v5/pkg/domain/meta"
	"github.com/nori-io/common/v5/pkg/domain/registry"
	"github.com/nori-io/common/v5/pkg/errors"
)

const (
	SessionInterface    meta.Interface = "nori/session/Session@2.0.0"
	SessionIdContextKey                = "nori.session.id"
)

type Session interface {
	Get(key []byte, data interface{}) error
	Save(key []byte, data interface{}, exp time.Duration) error
	Delete(key []byte) error
	IsExists(key []byte) (bool, error)
}

func GetSession(r registry.Registry) (Session, error) {
	instance, err := r.Interface(SessionInterface)
	if err != nil {
		return nil, err
	}
	i, ok := instance.(Session)
	if !ok {
		return nil, errors.InterfaceAssertError{
			Interface: SessionInterface,
		}
	}
	return i, nil
}
