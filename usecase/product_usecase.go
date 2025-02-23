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

func (pu *ProductUseCase) GetProductsByID(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductsByID(id_product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUseCase) DeleteProduct(id_product int) (error){
	err := pu.repository.DeleteProduct(id_product)
	if err != nil{
		return err
	}
	return nil
}
