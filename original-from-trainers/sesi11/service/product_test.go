package service

import (
	"fmt"
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

func TestProductServiceCreateProductFail(t *testing.T) {
	p := entity.Product{
		Id:   2,
		Name: "Product 2",
	}
	productRepo.(*repository.ProductRepositoryMock).Mock.On("CreateProduct", &p).Return(nil)

	success, err := productService.CreateProduct(&p)

	assert.NotNil(t, err)
	assert.False(t, success)
	assert.Equal(t, fmt.Sprintf("create product with id %d fail", p.Id), err.Error())

}

func TestProductServiceCreateProductSuccess(t *testing.T) {
	p := entity.Product{
		Id:   2,
		Name: "Product 2",
	}
	productRepo.(*repository.ProductRepositoryMock).Mock.On("CreateProduct", &p).Return("")

	success, err := productService.CreateProduct(&p)

	assert.Nil(t, err)
	assert.True(t, success)

}

func TestProductServiceGetAllFail(t *testing.T) {
	productRepo.(*repository.ProductRepositoryMock).Mock.On("GetAll").Return(nil)

	products, err := productService.GetAllProduct()
	assert.NotNil(t, err)
	assert.Nil(t, products)
	assert.Equal(t, fmt.Sprintf("there is no product"), err.Error())

}

func TestProductServiceGetAllSuccess(t *testing.T) {
	p := []entity.Product{
		{
			Id:   10,
			Name: "",
		},
	}
	productRepo.(*repository.ProductRepositoryMock).Mock.On("GetAll").Return(p)

	products, err := productService.GetAllProduct()
	assert.Nil(t, err)
	assert.NotNil(t, products)

	pr := *products
	assert.NotEmpty(t, pr[0].Name)

}
