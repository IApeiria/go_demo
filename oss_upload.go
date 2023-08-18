/*package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
func main() {
	// yourBucketName填写存储空间名称。
	bucketName := "sixxxde"
	// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
	objectName := "/video/79059953_p0.png"
	// yourLocalFileName填写本地文件的完整路径。
	localFileName := "F:\\ACGN\\fantasy.pic\\79059953_p0.png"

	// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("http://oss-cn-beijing.aliyuncs.com", "LTAI5tLFxxxxxxxB5uBX2", "XD0jzP0elGVOLMxxxxxxxgapf")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		handleError(err)
	}
}*/

package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func main() {

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("https://oss-cn-beijing.aliyuncs.com", "LTAI5xxxxxxxxxxjuVxxxs", "tBd0rIMwxxxxxxxxxEOkTxxx")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("sidade")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObjectFromFile("video/gura2.mp4", "gura.mp4")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
