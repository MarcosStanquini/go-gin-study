package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository 
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
	repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productID, err := pu.repository.CreateProduct(product)
	if err != nil {
		 return model.Product{}, err
	}
	product.ID = productID
	return product, nil
}