package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"`
	ResourceID string    `json:"resouce_id" valid:"notnull" gorm:"type:varchar(255)"`
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"_" valid:"-"`                           ///não valida
	Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"` //Slice de jobs - FK VideoID
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
