package config

import (
	"log"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestInitConfig(t *testing.T) {
	//Requirement test list
	testTable := []struct {
		Name     string
		Params   Config
		Validate func(*testing.T, error)
	}{
		{
			Name: "Set Config - Empty",
			Params: Config{
				Name: "",
				Type: "",
				Path: "",
			},
			Validate: func(t *testing.T, result error) {
				require.NotEmpty(t, result)
			},
		},
		{
			Name: "Set Config - Empty Name",
			Params: Config{
				Name: "",
				Type: "json",
				Path: ".",
			},
			Validate: func(t *testing.T, result error) {
				require.Empty(t, result)
			},
		},
		{
			Name: "Set Config - Empty Type",
			Params: Config{
				Name: "config",
				Type: "",
				Path: ".",
			},
			Validate: func(t *testing.T, result error) {
				require.Empty(t, result)
			},
		},
		{
			Name: "Set Config - Empty Path",
			Params: Config{
				Name: "config",
				Type: "json",
				Path: "",
			},
			Validate: func(t *testing.T, result error) {
				require.Empty(t, result)
			},
		},
		{
			Name: "Set Config",
			Params: Config{
				Name: "config",
				Type: "json",
				Path: ".",
			},
			Validate: func(t *testing.T, result error) {
				log.Println(viper.GetString("helper.currency.api"))
				require.Empty(t, result)
			},
		},
	}

	for _, value := range testTable {
		t.Run(value.Name, func(t *testing.T) {
			result := value.Params.Init()
			value.Validate(t, result)
		})
	}
}
