package database

import "github.com/briankliwon/efishery-recruitment-backend/fetch-service/models"

/*
commodities Interface
List Function will be used to accessing data from rosurce commodities
*/
type CommoditiesRepositories interface {
	Get() ([]models.Commodities, error) // Get all data from commodities resource
}
