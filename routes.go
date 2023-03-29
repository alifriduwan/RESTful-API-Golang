package main

import (
	"codezard-pos/controller"

	"github.com/gin-gonic/gin"
)

func serveRoutes(r *gin.Engine) {
	// /products GET
	// /products POST
	// /products/1 GET
	// /products/1 PATCH
	// /products/1 DELETE
	productController := controller.Product{}

	productGroup := r.Group("/products")
	productGroup.GET("", productController.FindAll)
	productGroup.GET("/:id", productController.FindOne)
	productGroup.DELETE("", productController.Create)
	productGroup.PATCH("/:id", productController.Update)
	productGroup.DELETE("/:id", productController.Delete)
}
