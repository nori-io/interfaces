package session


import (
	"github.com/nori-io/common/v3/errors"
	"github.com/nori-io/common/v3/meta"
	"github.com/nori-io/common/v3/plugin"
	"google.golang.org/grpc"
)

const (
	SessionInterface meta.Interface = "nori/session/Session@1.0.0"
)

type Session interface {
	Register(func(*grpc.Server)) // в качестве параметра функции функция и хэндлер реализующий интерфейс grpc сервиса
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
