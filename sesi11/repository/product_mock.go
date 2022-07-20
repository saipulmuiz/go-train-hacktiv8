package repository

import (
	"unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func NewProductRepositoryMock(mock mock.Mock) ProductRepository {
	return &ProductRepositoryMock{Mock: mock}
}

func (m *ProductRepositoryMock) FindById(id int) *entity.Product {
	arguments := m.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(entity.Product)
	return &product
}

func (p *ProductRepositoryMock) CreateProduct(product *entity.Product) bool {
	return true
}

func (p *ProductRepositoryMock) GetAll() *[]entity.Product {
	return nil
}
