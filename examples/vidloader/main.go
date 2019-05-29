package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/stek29/vk"
	"github.com/stek29/vk/vkapi"
)

func uploadVideo(uploadURL, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("video_file", filepath.Base(file.Name()))
	if err != nil {
		return err
	}
	io.Copy(part, file)
	writer.Close()

	req, err := http.NewRequest("POST", uploadURL, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Unexpected status code: Not 200 but %v", resp.StatusCode)
	}

	return nil
}

func main() {
	var vkToken string
	var videoPath string
	var params vkapi.VideoSaveParams

	flag.StringVar(&vkToken, "vk-token", "", "vk auth token")
	flag.StringVar(&videoPath, "video", "", "path to video to upload")

	flag.StringVar(&params.Name, "name", "", "video name")
	flag.StringVar(&params.Description, "description", "", "video description")
	flag.IntVar(&params.GroupID, "group-id", 0, "Group ID to post to")

	flag.Parse()

	client, err := vk.NewBaseAPI(vk.BaseAPIConfig{AccessToken: vkToken})
	if err != nil {
		panic(err)
	}
	resp, err := vkapi.Video{API: client}.Save(params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("INFO: Upload response: %+v\n", resp)

	fmt.Printf("Created Video: https://vk.com/video%v_%v\n", resp.OwnerID, resp.VideoID)
	fmt.Printf("INFO: Upload URL: %v\n", resp.UploadURL)

	err = uploadVideo(resp.UploadURL, videoPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("INFO: Video was uploaded")
}
