package main

import (
	"github.com/gin-gonic/gin"
	"github.com/visitor-verification-project/handler"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.Static("/img", "./images")

	r.GET("/", handler.ShowHomePage)
	r.GET("/stream", handler.ShowStreamPage)
	r.GET("/video_log", handler.ShowVideoLogPage)

	r.Run(":8080")
}

// video_log.html에서 캡처한 이미지(약 15장)전부 띄울 수 있게 구현필요
