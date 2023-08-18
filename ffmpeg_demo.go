package main

import (
	"fmt"
	"os/exec"
)

func main() {
	inputFile := "gura.mp4"
	outputFile := "gura.jpg"

	// 执行 FFmpeg 命令
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-vframes", "1", outputFile)

	// 开始执行命令
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("截取第一帧完成！")
}
