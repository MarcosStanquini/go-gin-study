package controller

import (
	"go-api/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
	"go-api/model"
)

type ProductController struct {  
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {  
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

func(p *ProductController) CreateProduct(ctx *gin.Context){
	var product model.Product  
	err := ctx.BindJSON(&product)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	insertedProduct, err := p.productUseCase.CreateProduct(product)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, insertedProduct)
}
