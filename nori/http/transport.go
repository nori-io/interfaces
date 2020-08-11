package http

import (
	"context"
	"net/http"

	"github.com/nori-io/common/v3/errors"
	"github.com/nori-io/common/v3/meta"
	"github.com/nori-io/common/v3/plugin"
)

const (
	TransportInterface meta.Interface = "nori/http/Transport@1.0.0"
)

type BeforeFunc func(context.Context, *http.Request) context.Context

type ServerAfterFunc func(context.Context, http.ResponseWriter) context.Context

type Transport interface {
	ToContext() BeforeFunc
	ToTransport(ctx context.Context, w http.ResponseWriter /*at interfaces.AccessTokener*/) ServerAfterFunc
}

func GetTransport(r plugin.Registry) (Transport, error) {
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
