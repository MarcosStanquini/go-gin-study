package main

import (
    _ "go-api/cmd/docs"
    "go-api/controller"
    "go-api/db"
    "go-api/repository"
    "go-api/usecase"
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
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

    server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    server.Run(":8000")
}