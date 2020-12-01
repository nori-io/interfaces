package http

import (
	"context"
	"net/http"

	"github.com/nori-io/common/v3/pkg/errors"
	"github.com/nori-io/common/v3/pkg/domain/meta"
	"github.com/nori-io/common/v3/pkg/domain/registry"
)

const (
	TransportInterface meta.Interface = "nori/http/Transport@1.0.0"
	AuthTokenContextKey                = "nori.auth.token"
	)

type AccessTokener interface {
	GetAccessToken() string
}


type BeforeFunc func(context.Context, *http.Request) context.Context

type ServerAfterFunc func(context.Context, http.ResponseWriter) context.Context

type Transport interface {
	ToContext() BeforeFunc
	ToTransport(ctx context.Context, w http.ResponseWriter, at AccessTokener) ServerAfterFunc
}

func GetTransport(r registry.Registry) (Transport, error) {
	instance, err := r.Interface(TransportInterface)
	if err != nil {
		return nil, err
	}
	i, ok := instance.(Transport)
	if !ok {
		return nil, errors.InterfaceAssertError{
			Interface: TransportInterface,
		}
	}
	return i, nil
}
