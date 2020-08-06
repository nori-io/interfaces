package pg

import (
	"github.com/jinzhu/gorm"

	"github.com/nori-io/common/v3/errors"
	"github.com/nori-io/common/v3/plugin"
	"github.com/nori-io/common/v3/meta"
)

const (
	GormInterface meta.Interface = "public/sql/Gorm@v1.9.15"
)

func GetGorm(r plugin.Registry) (*gorm.DB, error) {
	instance, err := r.Interface(GormInterface)
	if err != nil {
		return nil, err
	}
	i, ok := instance.(*gorm.DB)
	if !ok {
		return nil, errors.InterfaceAssertError{
			Interface: GormInterface,
		}
	}
	return i, nil
}
