package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
}

func (p Product) FindAll(ctx *gin.Context) { //method หาสินค้าทั้งหมด
	search := ctx.Query("search")
	categoryId := ctx.Query("categoryId")
	ctx.JSON(http.StatusOK, gin.H{
		"FindAll":    "OK",
		"Search":     search,
		"CatagoryID": categoryId,
	})
}

func (p Product) FindOne(ctx *gin.Context) { //method หาสินค้าตัวเดียว
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (p Product) Create(ctx *gin.Context) {

}

func (p Product) Update(ctx *gin.Context) {

}

func (p Product) Delete(ctx *gin.Context) {

}
