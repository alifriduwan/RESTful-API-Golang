package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //Create web server

	// http://example.com/hello GET POST PATCH DELETE
	// GET -> ส่งข้อมูลออกมาในรูปแบบ json / ส่ง status code
	//json {"key": "value"} => {"greeting":"Hello Server"}
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"greeting": "Hello Server"})
	})

	r.Run()
}
