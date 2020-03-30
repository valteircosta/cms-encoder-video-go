package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

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

	f, err := os.Create(storagePath + "/" + video.Uuid + ".mp4")

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
func (video *Video) Fragment(storedPath string) Video {

	err := os.Mkdir(storedPath+"/"+video.Uuid, os.ModePerm)
	if err != nil {
		video.Status = "Error"
		fmt.Println(err.Error())
	}
	fmt.Println("Make fragment " + video.Uuid)
	// Origem e destino
	source := storedPath + "/" + video.Uuid + ".mp4"
	target := storedPath + "/" + video.Uuid + ".frag"
	// Comando de fragmentação
	cmd := exec.Command("mp4fragment", source, target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		video.Status = "Error"
		fmt.Println(err.Error())
	}
	printOutput(output)
	return *video
}
func printOutput(out []byte) {
	if len(out) > 0 {
		fmt.Printf("===> Output %s\n", string(out))
	}
}
func (video *Video) Encode(storedPath string) Video {

	cmdArgs := []string{}
	cmdArgs = append(cmdArgs, storedPath+"/"+video.Uuid+".frag")
	cmdArgs = append(cmdArgs, "--use-segment-timeline")
	cmdArgs = append(cmdArgs, "-o")
	cmdArgs = append(cmdArgs, storedPath+"/"+video.Uuid)
	cmdArgs = append(cmdArgs, "-f")
	cmdArgs = append(cmdArgs, "--exec-dir")
	cmdArgs = append(cmdArgs, "/usr/local/bin")
	cmd := exec.Command("mp4dash", cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		video.Status = "Error"
		fmt.Println(err.Error())
	}
	printOutput(output)
	return *video
}
func (video *Video) UploadObject(completePath string, storagePath string, bucketName string, client *storage.Client, ctx context.Context) error {
	path := strings.Split(completePath, storagePath+"/")

	f, err := os.Open(completePath)
	if err != nil {
		fmt.Println("Error during the upload", err.Error())
		return err
	}
	defer f.Close()
	wc := client.Bucket(bucketName).Object(path[1]).NewWriter(ctx)
	wc.ACL := []storage.ACLRule

	## visto até  o minuto 7 do video

}
