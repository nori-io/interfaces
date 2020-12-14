package pg

import (
	"github.com/go-pg/pg"
	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	"github.com/nori-io/common/v4/pkg/errors"
)

const (
	PgInterface meta.Interface = "public/sql/PG@v8.0.7"
)

func GetPG(r registry.Registry) (*pg.DB, error) {
	instance, err := r.Interface(PgInterface)
	if err != nil {
		return nil, err
	}
	i, ok := instance.(*pg.DB)
	if !ok {
		return nil, errors.InterfaceAssertError{
			Interface: PgInterface,
		}
	}
	return i, nil
}
