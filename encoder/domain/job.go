package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

//First method executed
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Job struct {
	ID               string    `json:"job_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OutputBucketPath string    `json:"output_bucket_path" valid:"notnull"`
	Status           string    `json:"status" valid:"notnull"`
	Video            *Video    `json:"video" valid:"-"`
	VideoID          string    `json:"-" valid:"-" gorm:"column:video_id;type:uuid;notnull"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `json:"created_at" valid:"-"`
	UpdatedAt        time.Time `json:"updated_at" valid:"-"`
}

//Função que funciona como um construtor retornado dois valores um job ou um errror
func NewJob(output string, status string, video *Video) (*Job, error) {
	//tipo construtor com parametros
	job := Job{
		OutputBucketPath: output,
		Status:           status,
		Video:            video,
	}
	//Roda o prepare preenchendo os campos
	job.prepare()
	err := job.Validate()
	if err != nil {
		return nil, err
	}

	return &job, nil

}

//Inicial minúscula visivel dentro packge, no caso domain
func (job *Job) prepare() {
	job.ID = uuid.NewV4().String()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
}

// (job *Job) Transforma a função em um methodo ligado a Job
func (job *Job) Validate() error {
	_, err := govalidator.ValidateStruct(job)
	if err != nil {
		return err
	}
	return nil

}
