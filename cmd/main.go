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

	server.Run(":8000")
}
