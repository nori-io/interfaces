package http

import (
	"github.com/nori-io/common/v3/pkg/domain/registry"
	"github.com/nori-io/common/v3/pkg/errors"
	"net/http"

	"github.com/nori-io/common/v3/pkg/domain/meta"
)

const HttpInterface meta.Interface = "nori/http/HTTP"

type Http interface {
	http.Handler

	Handle(pattern string, h http.Handler)
	HandleFunc(pattern string, h http.HandlerFunc)

	Method(method, pattern string, h http.Handler)
	MethodFunc(method, pattern string, h http.HandlerFunc)

	Connect(pattern string, h http.HandlerFunc)
	Delete(pattern string, h http.HandlerFunc)
	Get(pattern string, h http.HandlerFunc)
	Head(pattern string, h http.HandlerFunc)
	Options(pattern string, h http.HandlerFunc)
	Patch(pattern string, h http.HandlerFunc)
	Post(pattern string, h http.HandlerFunc)
	Put(pattern string, h http.HandlerFunc)
	Trace(pattern string, h http.HandlerFunc)
}

func GetHttp(r registry.Registry) (Http, error) {
	instance, err := r.Interface(HttpInterface)
	if err != nil {
		return nil, err
	}
	i, ok := instance.(Http)
	if !ok {
		return nil, errors.InterfaceAssertError{
			Interface: HttpInterface,
		}
	}
	return i, nil
}
