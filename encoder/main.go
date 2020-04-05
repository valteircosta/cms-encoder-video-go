package main

import (
	"encoder/domain"
	v "encoder/domain"

	log "github.com/sirupsen/logrus"
)

func main() {
	var video v.Video
	doneUpload := make(chan bool)

	data := []byte("{\"uuid\":\"convite\",\"path\":\"convite.mp4\",\"status\":\"Pending\"}")
	video.Unmarshal(data)
	video.Dowloand("codeeducation-test", "/tmp")
	video.Fragment("/tmp")
	video.Encode("/tmp")

	go domain.ProcessUpload(video, "/tmp", "codeeducation-test", doneUpload)
	<-doneUpload

	video.Finish("/tmp")

	log.Info(video.Path)

}
