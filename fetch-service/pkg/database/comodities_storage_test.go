package database

import (
	"log"
	"testing"

	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestGetCommoditiesStorage(t *testing.T) {
	//Setting Up Viper configuration
	viperConfig := config.Config{
		Name: "config",
		Type: "json",
		Path: "../config",
	}
	err := viperConfig.Init()
	if err != nil {
		log.Fatal(err)
	}

	//Requirement test list
	testTable := []struct {
		Name     string
		Validate func(*testing.T, error)
	}{
		{
			Name: "Get Commodities",
			Validate: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
	}

	storage := &CommoditiesStorage{}
	for _, value := range testTable {
		t.Run(value.Name, func(t *testing.T) {
			_, err := storage.Get()
			value.Validate(t, err)
		})
	}
}
