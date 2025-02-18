package controller

import (
	"go-api/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ProductController struct {  // Alterado para seguir a convenção de Go
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {  // Alterado para seguir a convenção de Go
	return ProductController{
		productUseCase: usecase,
	}
}

func (p *ProductController) GetProduct(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if(err != nil){
		ctx.JSON(http.StatusInternalServerError, err)

	}
	ctx.JSON(http.StatusOK, products)
}
