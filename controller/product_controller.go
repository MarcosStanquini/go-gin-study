package controller

import (
	"go-api/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.product
}

func NewProductController() productController {
	return productController{}

}

func (p *productController) GetProduct(ctx *gin.Context) {
	products := []model.Product{
		{
		ID:    1,
		Name:  "Batata Frita",
		Price: 20,
		},
	}	
		ctx.JSON(http.StatusOK, products)

}
