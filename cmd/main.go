package main

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	productController := controller.NewProductController() 

	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	} )

	server.GET("/products", productController.GetProduct)

	server.Run(":8000")

}