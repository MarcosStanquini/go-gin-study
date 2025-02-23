package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUsecase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecase)
	
	server.GET("/products", ProductController.GetProduct)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductsByID)
	server.DELETE("/product/:productId", ProductController.DeleteProduct)

	server.Run(":8000")
}
