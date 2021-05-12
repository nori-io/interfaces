package sqlx

import (
	"github.com/jmoiron/sqlx"
	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	"github.com/nori-io/common/v4/pkg/errors"
)

const (
	SqlxInterface meta.Interface = "database/sqlx/Sqlx@v1.2.0"
)

func GetSqlx(r registry.Registry) (*sqlx.DB, error) {
	instance, err := r.Interface(SqlxInterface)
	if err != nil {
		return nil, err
	}
	i, ok := instance.(*sqlx.DB)
	if !ok {
		return nil, errors.InterfaceAssertError{
			Interface: SqlxInterface,
		}
	}
	return i, nil
}
