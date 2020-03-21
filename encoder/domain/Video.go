package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"cloud.google.com/go/storage"
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
func (video *Video) Dowloand(bucketName string, storagePath string) (Video, error) {

	// Cria contexo
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		video.Status = "Error"
		fmt.Println(err.Error())
		return *video, err
	}

	bkt := client.Bucket(bucketName)
	obj := bkt.Object(video.Path)
	r, err := obj.NewReader(ctx)

	if err != nil {
		video.Status = "Error"
		fmt.Println(err.Error())
		return *video, err
	}

	defer r.Close()
	body, err := ioutil.ReadAll(r)

	if err != nil {
		video.Status = "Error"
		fmt.Println(err.Error())
		return *video, err
	}

	f, err := os.Create(storagePath + "/" + video.Path + ".mp4")
	if err != nil {
		video.Status = "Error"
		fmt.Println(err.Error())
		return *video, err
	}

	_, err = f.Write(body)
	if err != nil {
		video.Status = "Error"
		fmt.Println(err.Error())
		return *video, err
	}

	defer f.Close()

	fmt.Println("Video ", video.Uuid, " has been stored")

	return *video, nil

}
