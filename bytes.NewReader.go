package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	data := []byte("Hello World")

	reader := bytes.NewReader(data)

	// 读取一个字节
	oneByte := make([]byte, 1)
	reader.Read(oneByte) //读出一个，少一个
	fmt.Println(string(oneByte))

	// 读取剩余所有字节
	restBytes := make([]byte, len(data))
	reader.Read(restBytes)
	fmt.Println(string(restBytes))

	// 重新读
	reader = bytes.NewReader(data)
	allBytes := make([]byte, len(data))
	io.ReadFull(reader, allBytes)
	fmt.Println(string(allBytes))
}
