package service

import (
	"fmt"
	"unit-test/entity"
	"unit-test/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetOneProductById(id int) (*entity.Product, error) {
	product := s.repo.FindById(id)
	if product == nil {
		return nil, fmt.Errorf("product with id %d not found", id)
	}

	return product, nil
}

func (s *ProductService) CreateProduct(p *entity.Product) (bool, error) {
	success := s.repo.CreateProduct(p)
	if !success {
		return false, fmt.Errorf("create product with id %d fail", p.Id)
	}
	return true, nil
}

func (s *ProductService) GetAllProduct() (*[]entity.Product, error) {
	product := s.repo.GetAll()
	if product == nil {
		return nil, fmt.Errorf("there is no product")
	}

	return product, nil
}
