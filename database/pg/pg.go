package pg

import (
	"github.com/go-pg/pg/v10"
	"github.com/nori-io/common/v5/pkg/domain/meta"
	"github.com/nori-io/common/v5/pkg/domain/registry"
	"github.com/nori-io/common/v5/pkg/errors"
)

const (
	PgInterface meta.Interface = "database/pg/PG@v10.9.3"
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
