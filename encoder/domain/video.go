package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID         string    `valid:"uuid"`
	ResourceID string    `valid:"notnull"`
	FilePath   string    `valid:"notnull"`
	CreatedAt  time.Time `valid:"-"` ///não valida
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}
func NewVideo() *Video {
	return &Video{}
}
func (v *Video) Validate() error {
	valid, err := govalidator.ValidateStruct(v)
	if !valid {
		return err
	}
	return nil
}
