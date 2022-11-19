package services

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/models"
	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/pkg/config"
	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/pkg/database"
	"github.com/stretchr/testify/require"
)

// Setting Up commodities service
var storage = &database.CommoditiesStorageMock{}
var service = CommoditiesService{Storage: storage}

func TestGetCommodities(t *testing.T) {
	//Set dummy data to testing logic and respone
	commodities := []models.Commodities{
		{
			Uuid:         "0f8ca8a9-108e-44f5-8f2d-81e1fc8348ff",
			Komoditas:    "LELE",
			AreaProvinsi: "ACEH",
			AreaKota:     "ACEH KOTA",
			Size:         "120",
			Price:        "68000",
			TglParsed:    "2022-01-07T22:47:27Z",
			Timestamp:    "1641595647520",
		},
		{
			Uuid:         "5abe591a-4075-480d-b38e-40847cd63cb7",
			Komoditas:    "MAS",
			AreaProvinsi: "JAWA TENGAH",
			AreaKota:     "PURWOREJOL",
			Size:         "30",
			Price:        "27000",
			TglParsed:    "2022-03-02T02:17:54Z",
			Timestamp:    "1646187474809",
		},
	}
	storage.Mock.On("Get").Return(commodities)

	//Setting Up Viper configuration
	viperConfig := config.Config{
		Name: "config",
		Type: "json",
		Path: "../pkg/config",
	}
	err := viperConfig.Init()
	if err != nil {
		log.Fatal(err)
	}

	//Requirement test list
	testTable := []struct {
		Name     string
		Uri      string
		Validate func(*testing.T, int)
	}{
		{
			Name: "Get Commodities",
			Uri:  "/commodities",
			Validate: func(t *testing.T, result int) {
				require.Equal(t, result, http.StatusOK)
			},
		},
	}

	for _, value := range testTable {
		t.Run(value.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, value.Uri, nil)
			w := httptest.NewRecorder()
			service.Get(w, req)
			res := w.Result()
			defer res.Body.Close()
			value.Validate(t, res.StatusCode)
		})
	}
}

func TestGetCommoditiesAggregation(t *testing.T) {
	commodities := []models.Commodities{
		{
			Uuid:         "0f8ca8a9-108e-44f5-8f2d-81e1fc8348ff",
			Komoditas:    "LELE",
			AreaProvinsi: "ACEH",
			AreaKota:     "ACEH KOTA",
			Size:         "120",
			Price:        "68000",
			TglParsed:    "2022-01-07T22:47:27Z",
			Timestamp:    "1641595647520",
		},
		{
			Uuid:         "5abe591a-4075-480d-b38e-40847cd63cb7",
			Komoditas:    "MAS",
			AreaProvinsi: "JAWA TENGAH",
			AreaKota:     "PURWOREJOL",
			Size:         "30",
			Price:        "27000",
			TglParsed:    "2022-03-02T02:17:54Z",
			Timestamp:    "1646187474809",
		},
	}
	storage.Mock.On("Get").Return(commodities)
	//Requirement test list
	testTable := []struct {
		Name     string
		Uri      string
		Validate func(*testing.T, int)
	}{
		{
			Name: "Get Commodities Aggregation",
			Uri:  "/aggregate",
			Validate: func(t *testing.T, result int) {
				require.Equal(t, result, http.StatusOK)
			},
		},
	}

	for _, value := range testTable {
		t.Run(value.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, value.Uri, nil)
			w := httptest.NewRecorder()
			service.GetAggregate(w, req)
			res := w.Result()
			defer res.Body.Close()
			value.Validate(t, res.StatusCode)
		})
	}
}
