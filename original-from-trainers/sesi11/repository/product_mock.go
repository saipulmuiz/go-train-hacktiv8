package repository

import (
	"fmt"
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

func (m *ProductRepositoryMock) CreateProduct(product *entity.Product) bool {
	arguments := m.Mock.Called(product)
	fmt.Println(arguments.Get(0) != nil, arguments.Get(0))
	return arguments.Get(0) != nil
}

func (m *ProductRepositoryMock) GetAll() *[]entity.Product {
	arguments := m.Mock.Called()

	if arguments.Get(0) == nil {
		return nil
	}

	products := arguments.Get(0).([]entity.Product)
	return &products
}
