package database

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/models"
	"github.com/spf13/viper"
)

type CommoditiesStorage struct{}

func (storage *CommoditiesStorage) Get() (dataCommodities []models.Commodities, err error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", viper.GetString("commodities.url"), nil)
	if err != nil {
		return []models.Commodities{}, err
	}
	response, err := client.Do(request)
	if err != nil {
		return []models.Commodities{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, &dataCommodities); err != nil {
		return []models.Commodities{}, err
	}
	return dataCommodities, nil
}
