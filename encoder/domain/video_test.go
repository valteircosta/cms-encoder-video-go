package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()
	require.Error(t, err)

}

func TestVideoIDisNotUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "abc"
	video.ResourceID = "ResourcID-1"
	video.FilePath = "FilePath"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestValidation(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "ResourcID-1"
	video.FilePath = "FilePath"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.NoError(t, err)
}
