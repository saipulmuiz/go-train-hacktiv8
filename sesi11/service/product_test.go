package service

import (
	"testing"
	"unit-test/entity"
	"unit-test/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepo = repository.NewProductRepositoryMock(mock.Mock{})
var productService = NewProductService(productRepo)

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepo.(*repository.ProductRepositoryMock).Mock.On("FindById", 1).Return(nil)

	product, err := productService.GetOneProductById(1)

	assert.Nil(t, product)
	assert.NotNil(t, err)
}

func TestProductServiceGetOneProduct(t *testing.T) {
	p := entity.Product{
		Id:   2,
		Name: "Product 2",
	}

	productRepo.(*repository.ProductRepositoryMock).Mock.On("FindById", p.Id).Return(p)

	product, err := productService.GetOneProductById(p.Id)

	t.Run("must be not error", func(t *testing.T) {
		assert.Nil(t, err)
	})
	t.Run("product is exist", func(t *testing.T) {
		assert.NotNil(t, product)
	})
	t.Run("product name is "+product.Name, func(t *testing.T) {
		assert.Equal(t, p.Name, product.Name)
	})
}

func TestProductServiceGetAllProduct(t *testing.T) {
	p := []entity.Product{
		{
			Id:   1,
			Name: "Product 1",
		},
		{
			Id:   2,
			Name: "Product 2",
		},
	}

	productRepo.(*repository.ProductRepositoryMock).Mock.On("GetAll").Return(p)

	product, err := productService.GetAll()

	t.Run("FAIL", func(t *testing.T) {
		assert.Nil(t, err)
	})
	t.Run("Success", func(t *testing.T) {
		assert.NotNil(t, product)
	})
}

func TestProductServiceCreateProduct(t *testing.T) {
	p := entity.Product{
		Id:   1,
		Name: "Product 3",
	}

	productRepo.(*repository.ProductRepositoryMock).Mock.On("CreateProduct").Return(p)

	product, err := productService.GetAll()

	t.Run("FAIL", func(t *testing.T) {
		assert.Nil(t, err)
	})
	t.Run("Success", func(t *testing.T) {
		assert.NotNil(t, product)
	})
}
