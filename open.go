package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer f.Close()

	// 读取文件内容
	buf := make([]byte, 1024)
	n, err := f.Read(buf) //把f里的东西送到buf里
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}

	fmt.Println(string(buf), n)
}
