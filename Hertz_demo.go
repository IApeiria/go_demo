package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func main() {
	h := server.Default()

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		name := ctx.Query("name")
		ctx.JSON(200, utils.H{"message": "pong"})

		ctx.JSON(200, utils.H{
			"message": "Hello " + name,
		})
	})
	h.POST("/add", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, utils.H{
			"success":       true,
			"add a message": "message"})
	})

	// 设置响应

	h.Spin()
}
