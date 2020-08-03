package mail

import (

	"github.com/nori-io/common/v3/errors"
	"github.com/nori-io/common/v3/meta"
	"github.com/nori-io/common/v3/plugin"
)

const (
	MailInterface meta.Interface = "nori/mail/Mail@1.0.0"
)

type Mail interface {
	RegisterBeforeSend(f func(interface{}))
	RegisterAfterSend(f func(interface{}))
	RegisterBeforeDelivery(f func(interface{}))
	RegisterAfterDelivery(f func(interface{}))
	RegisterOnDeliveryError(f func(interface{}))
	Send(msg interface{}) error
}

func GetMail(r plugin.Registry) (Mail, error) {
	instance, err := r.Interface(MailInterface)
	if err != nil {
		return nil, err
	}
	i, ok := instance.(Mail)
	if !ok {
		return nil, errors.InterfaceAssertError{
			Interface: MailInterface,
		}
	}
	return i, nil
}
