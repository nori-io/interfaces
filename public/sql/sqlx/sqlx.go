package sqlx

import (
	"github.com/jmoiron/sqlx"
	"github.com/nori-io/common/v3/errors"
	"github.com/nori-io/common/v3/plugin"
	"github.com/nori-io/common/v3/meta"
)

const (
	SqlxInterface meta.Interface = "public/sql/Sqlx@v1.2.0"
)

func GetSqlx(r plugin.Registry) (*sqlx.DB, error) {
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
