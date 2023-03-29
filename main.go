package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //Create web server
	serveRoutes(r)
	r.Run()
}
