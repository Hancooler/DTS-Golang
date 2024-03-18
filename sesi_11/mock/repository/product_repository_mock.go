package repository

import (
	"mock/entity"

	"github.com/stretchr/testify/mock"
)

type productrepositoryMock struct {
	Mock mock.Mock
}

func (repository *productrepositoryMock) FindByID(id string) *entity.Product {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil

	}
	product := arguments.Get(0).(entity.Product)

	return &product

}
