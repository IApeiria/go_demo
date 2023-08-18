package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("http://apeiria.xyz/411", func(context *gin.Context) {
		context.String(200, "value:%v", "hello")
	})
	r.Run("http://apeiria.xyz/411")
}
