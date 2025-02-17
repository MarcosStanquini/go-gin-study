package usecase

import "go-api/model"

type ProductUsecase struct {
}

func newProductUsecase() productUsecase {
	return productUsecase{}
}

func (pu *productUsecase) GetProducts() ([]model.Product, error){
	return []model.Product{}, nil
}
