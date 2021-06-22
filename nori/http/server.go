package http

import (
	"net/http"

	"github.com/nori-io/common/v5/pkg/domain/meta"
	"github.com/nori-io/common/v5/pkg/domain/registry"
	"github.com/nori-io/common/v5/pkg/errors"
)

const RouterInterface meta.Interface = "nori/http/Router@2.0.0"

type Router interface {
	http.Handler
	Handle(pattern string, h http.Handler)
	HandleFunc(pattern string, h http.HandlerFunc)
	Method(method, pattern string, h http.Handler)
	MethodFunc(method, pattern string, h http.HandlerFunc)
	Route(pattern string, fn func(h Router)) Router
	Mount(pattern string, h http.Handler)
	Connect(pattern string, h http.HandlerFunc)
	Delete(pattern string, h http.HandlerFunc)
	Get(pattern string, h http.HandlerFunc)
	Head(pattern string, h http.HandlerFunc)
	Options(pattern string, h http.HandlerFunc)
	Patch(pattern string, h http.HandlerFunc)
	Post(pattern string, h http.HandlerFunc)
	Put(pattern string, h http.HandlerFunc)
	Trace(pattern string, h http.HandlerFunc)
	Use(middlewares ...func(http.Handler) http.Handler)
	With(middlewares ...func(http.Handler) http.Handler) Router
	NotFound(h http.HandlerFunc)
	MethodNotAllowed(h http.HandlerFunc)
	//URLParam(r *http.Request, key string) string
}

func GetRouter(r registry.Registry) (Router, error) {
	instance, err := r.Interface(RouterInterface)
	if err != nil {
		return nil, err
	}
	i, ok := instance.(Router)
	if !ok {
		return nil, errors.InterfaceAssertError{
			Interface: RouterInterface,
		}
	}
	return i, nil
}
