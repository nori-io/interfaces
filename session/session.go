package session

import (
	"context"
	"github.com/nori-io/common/v3/endpoint"
	"github.com/nori-io/common/v3/errors"
	"github.com/nori-io/common/v3/meta"
	"github.com/nori-io/common/v3/plugin"
	"time"
)

const (
	SessionInterface meta.Interface = "nori/session/Session@1.0.0"
	SessionIdContextKey                = "nori.session.id"
)

type Session interface {
	Get(key []byte, data interface{}) error
	Save(key []byte, data interface{}, exp time.Duration) error
	Delete(key []byte) error
	SessionId(ctx context.Context) []byte
	Verify() endpoint.Middleware
}

func GetSession(r plugin.Registry) (Session, error) {
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
