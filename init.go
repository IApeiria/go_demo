package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func main() {
	// 从环境变量中获取访问凭证。运行本代码示例之前，请先配置环境变量。
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("http://oss-cn-beijing.aliyuncs.com", "LTAI5tLFqJRbUXgWLCB5uBX2", "XD0jzP0elGVOLMaspX1uYpwwfkgapf", oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	fmt.Printf("client:%#v\n", client)
}
