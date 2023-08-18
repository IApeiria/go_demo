package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("请输入一段文本: ")

	// 创建一个带缓冲的读取器，读取标准输入
	reader := bufio.NewReader(os.Stdin)

	// 读取一行文本
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入出错:", err)
		return
	}

	fmt.Println("你输入的文本是:", text)
}
