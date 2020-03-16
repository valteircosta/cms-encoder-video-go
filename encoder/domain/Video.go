package domain

import (
	"encoding/json"
)

type Video struct {
	Uuid   string `json:"uuid"`
	Path   string `json:"path"`
	Status string `json:"status"`
}

func (video *Video) Unmarshal(payload []byte) Video {
	err := json.Unmarshal(payload, &video)
	if err != nil {
		panic(err)
	}
	return *video

}
