package pg

import (
	"github.com/nori-io/common/v5/pkg/domain/meta"
	"github.com/nori-io/common/v5/pkg/domain/registry"
	"github.com/nori-io/common/v5/pkg/errors"
	"gorm.io/gorm"
)

const (
	GormInterface meta.Interface = "database/gorm/Gorm@v1.21.10"
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
