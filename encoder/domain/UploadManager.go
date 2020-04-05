package domain

import (
	"context"
	"fmt"
	"runtime"

	"cloud.google.com/go/storage"
)

func ProcessUpload(file Video, storagePath string, bucketName string, doneUpload chan bool) {
	concurrency := 5
	in := make(chan int, runtime.NumCPU())
	ret := make(chan error)
	paths := file.GetVideoPaths()
	uploadClient, ctx := getClientUpload()
	if file.Status != "error" {
		fmt.Println("Uploandig", file.Uuid, "....")
		// Faz o loop pela concorrencia criando os workers
		for x := 0; x < concurrency; x++ {
			go upLoadWorker(file, storagePath, bucketName, in, ret, paths, uploadClient, ctx)
		}
		//Preenche o canal de "in" para o uploadWorker
		//Pega todos os videos que foram separados e joga no canal in
		go func() {
			for i := 0; i < len(paths); i++ {
				in <- i
			}
			//Fecha o canal
			close(in)
		}()
		// Verifica erros e para upload caso encontre
		for err := range ret {
			if err != nil {
				fmt.Println(err.Error())
				doneUpload <- true
				break
			}
		}
	}

}

func upLoadWorker(video Video, storagePath string, bucketName string, in chan int, returnChan chan error, paths []string, uploadClient *storage.Client, ctx context.Context) {
	for x := range in {
		fmt.Println("Object ", x)
		err := video.UploadObject(paths[x], storagePath, bucketName, uploadClient, ctx)
		returnChan <- err
	}
	returnChan <- fmt.Errorf("Upload completed")
}
func getClientUpload() (*storage.Client, context.Context) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("Failed to create Client", err.Error())
	}
	return client, ctx
}
