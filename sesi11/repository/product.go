package repository

import "unit-test/entity"

type ProductRepository interface {
	FindById(id int) *entity.Product
	CreateProduct(product *entity.Product) bool
	GetAll() *[]entity.Product
}

type productRepo struct {
}

func NewProductRepository() ProductRepository {
	return &productRepo{}
}

func (p *productRepo) FindById(id int) *entity.Product {
	/*
		SELECT * FROM products
		WHERE id = `id`
	*/
	return &entity.Product{}
}

func (p *productRepo) CreateProduct(product *entity.Product) bool {
	return true
}

func (p *productRepo) GetAll() *[]entity.Product {
	return &[]entity.Product{}
}
