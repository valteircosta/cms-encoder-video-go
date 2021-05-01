package domain_test

import (
	"encoder/domain"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	job, err := domain.NewJob("path", "converted", video)
	//test job !nil
	require.NotNil(t, job)
	//tst err = nil
	require.Nil(t, err)
}
