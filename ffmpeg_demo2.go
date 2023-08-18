package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 设置视频源文件路径
	inputFile := "gura.mp4"
	// 设置转码后文件路径
	outputFile := "output.mp4"

	// 设置 ffmpeg 命令行参数

	args := []string{"-i", inputFile, "-c", "copy", outputFile}

	// 创建 *exec.Cmd
	cmd := exec.Command("ffmpeg", args...)

	// 运行 ffmpeg 命令
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("转码成功")
}
