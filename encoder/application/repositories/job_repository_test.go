package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRespositoryDbInsert(t *testing.T) {
	/*
	  Roteiro do teste
	  Obter uma conexão
	  Obter um video
	  Anexar o video a job
	  Inserir a Job
	  Pesquisar e testar o id
	*/
	db := database.NewDbTest()
	defer db.Close() //Fecha db depois que usar

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "filepath"
	video.ResourceID = "resourceId"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)
	job, err := domain.NewJob("output_pat", "Pending", video)
	require.Nil(t, err)
	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)
	j, err := repoJob.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, video.ID, job.VideoID)

}
func TestJobRepositoryDbUpdate(t *testing.T) {
	/*
	  Roteiro do teste
	  Obter uma conexão
	  Obter um video
	  Anexar o video a job
	  Inserir a Job
	  Alterar Status para Complete
	  Pesquisar e testar o id
	*/
	db := database.NewDbTest()
	defer db.Close() //Fecha db depois que usar

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "filepath"
	video.ResourceID = "resourceId"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)
	job, err := domain.NewJob("output_pat", "Pending", video)
	require.Nil(t, err)
	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)
	job.Status = "Complete"
	repoJob.Update(job)
	j, err := repoJob.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)

}
