package pg

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-io/common/v3/pkg/domain/meta"
	"github.com/nori-io/common/v3/pkg/errors"
	"github.com/nori-io/common/v3/pkg/domain/registry"

)

const (
	GormInterface meta.Interface = "public/sql/Gorm@v1.9.15"
)

func GetGorm(r registry.Registry) (*gorm.DB, error) {
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
