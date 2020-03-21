package main

import (
	v "encoder/domain"

	log "github.com/sirupsen/logrus"
)

func main() {
	var video v.Video
	data := []byte("{\"uuid\":\"abcd1234\",\"path\":\"convite.mp4\",\"status\":\"Pending\"}")
	video.Unmarshal(data)
	video.Dowloand("codeeducation-test", "/tmp")
	video.Fragment("/tmp")
	log.Info(video.Path)

}
