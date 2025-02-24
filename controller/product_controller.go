package controller

import (
	"database/sql"
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: usecase,
	}
}

// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Product
// @Failure 500 {object} model.Response
// @Router /products [get]
func (p *ProductController) GetProduct(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body model.Product true "Product to create"
// @Success 201 {object} model.Product
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /product [post]
func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, insertedProduct)
}

// @Summary Get product by ID
// @Description Get product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param productId path int true "Product ID"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /product/{productId} [get]
func (p *ProductController) GetProductsByID(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id do produto precisa ser um número"})
		return
	}

	product, err := p.productUseCase.GetProductsByID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return

	}

	ctx.JSON(http.StatusOK, product)
}

// @Summary Delete product by ID
// @Description Delete product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param productId path int true "Product ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /product/{productId} [delete]
func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id do produto precisa ser um número"})
		return
	}

	err = p.productUseCase.DeleteProduct(productID)
	if err != nil {
		if err == sql.ErrNoRows {
			response := model.Response{
				Message: "Produto não encontrado na base de dados",
			}
			ctx.JSON(http.StatusNotFound, response)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	response := model.Response{
		Message: "Produto deletado com sucesso",
	}
	ctx.JSON(http.StatusOK, response)
}
