package service

type repository interface {
}

type ProductService struct {
	repository repository
}

func NewProductService(repository repository) *ProductService {
	return &ProductService{repository: repository}
}
