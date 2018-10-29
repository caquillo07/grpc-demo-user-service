package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	"github.com/gofrs/uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("id", uid.String())
}
