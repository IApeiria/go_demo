package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	videoURL := "https://sidade.oss-accelerate.aliyuncs.com/video/698087575478341632.mp4"

	//output := "gura4.jpg"

	cmd := exec.Command("ffmpeg", "-i", videoURL, "-vframes", "1", "-f", "image2pipe", "-")

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	imgBytes := buf.Bytes()
	imgBase64 := base64.StdEncoding.EncodeToString(imgBytes)

	fmt.Println(imgBase64)
}
