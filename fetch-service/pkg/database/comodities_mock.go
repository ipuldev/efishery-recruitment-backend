package database

import (
	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/models"
	"github.com/stretchr/testify/mock"
)

type CommoditiesStorageMock struct {
	Mock mock.Mock
}

func (storage *CommoditiesStorageMock) Get() ([]models.Commodities, error) {
	args := storage.Mock.Called()
	result := args.Get(0).([]models.Commodities)
	return result, nil
}
