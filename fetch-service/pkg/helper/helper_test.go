package helper

import (
	"log"
	"testing"

	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestGetCurrency(t *testing.T) {
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
		Name   string
		Params struct {
			Base string
			To   string
		}
		Validate func(*testing.T, float64)
	}{
		{
			Name: "Get Currency",
			Params: struct {
				Base string
				To   string
			}{
				Base: "USD",
				To:   "IDR",
			},
			Validate: func(t *testing.T, result float64) {
				require.Greater(t, result, 0.0)
			},
		},
		{
			Name: "Get Currency - Without Base",
			Params: struct {
				Base string
				To   string
			}{
				To: "IDR",
			},
			Validate: func(t *testing.T, result float64) {
				require.Equal(t, result, 0.0)
			},
		},
		{
			Name: "Get Currency - Without To",
			Params: struct {
				Base string
				To   string
			}{
				Base: "USD",
			},
			Validate: func(t *testing.T, result float64) {
				require.Equal(t, result, 0.0)
			},
		},
		{
			Name: "Get Currency - Empty",
			Params: struct {
				Base string
				To   string
			}{},
			Validate: func(t *testing.T, result float64) {
				require.Equal(t, result, 0.0)
			},
		},
	}

	for _, value := range testTable {
		t.Run(value.Name, func(t *testing.T) {
			result, _ := GetCurrency(value.Params.Base, value.Params.To)
			value.Validate(t, result)
		})
	}
}

func TestBubbleSort(t *testing.T) {
	//Requirement test list
	testTable := []struct {
		Name   string
		Params struct {
			List   []int
			Length int
		}
		Validate func(*testing.T, []int)
	}{
		{
			Name: "Set Bubble Sort",
			Params: struct {
				List   []int
				Length int
			}{
				List:   []int{2, 3, 1},
				Length: 3,
			},
			Validate: func(t *testing.T, result []int) {
				require.ElementsMatch(t, result, []int{1, 2, 3})
			},
		},
		{
			Name: "Set Bubble Sort - Empty list",
			Params: struct {
				List   []int
				Length int
			}{
				List:   []int{},
				Length: 0,
			},
			Validate: func(t *testing.T, result []int) {
				require.ElementsMatch(t, result, []int{})
			},
		},
		{
			Name: "Set Bubble Sort - List and length not match",
			Params: struct {
				List   []int
				Length int
			}{
				List:   []int{3, 2, 1},
				Length: 2,
			},
			Validate: func(t *testing.T, result []int) {
				require.ElementsMatch(t, result, []int{2, 3, 1})
			},
		},
	}

	for _, value := range testTable {
		t.Run(value.Name, func(t *testing.T) {
			BubbleSort(&value.Params.List, value.Params.Length)
			value.Validate(t, value.Params.List)
		})
	}
}

func TestAvgArray(t *testing.T) {
	//Requirement test list
	testTable := []struct {
		Name   string
		Params struct {
			List []int
		}
		Validate func(*testing.T, float64)
	}{
		{
			Name: "Get Average",
			Params: struct {
				List []int
			}{
				List: []int{1, 2, 3, 4, 5},
			},
			Validate: func(t *testing.T, result float64) {
				require.Equal(t, 3.0, result)
			},
		},
		{
			Name: "Get Average - One item list",
			Params: struct {
				List []int
			}{
				List: []int{1},
			},
			Validate: func(t *testing.T, result float64) {
				require.Equal(t, 1.0, result)
			},
		},
		{
			Name: "Set Average - Empty",
			Params: struct {
				List []int
			}{
				List: []int{},
			},
			Validate: func(t *testing.T, result float64) {
				require.Equal(t, 0.0, result)
			},
		},
	}

	for _, value := range testTable {
		t.Run(value.Name, func(t *testing.T) {
			result := AvgArray(value.Params.List)
			value.Validate(t, result)
		})
	}
}

func TestMedianArray(t *testing.T) {
	//Requirement test list
	testTable := []struct {
		Name   string
		Params struct {
			List []int
		}
		Validate func(*testing.T, float64)
	}{
		{
			Name: "Get Median",
			Params: struct {
				List []int
			}{
				List: []int{1, 2, 3, 4, 5},
			},
			Validate: func(t *testing.T, result float64) {
				require.Equal(t, 3.0, result)
			},
		},
		{
			Name: "Get Median - One item list",
			Params: struct {
				List []int
			}{
				List: []int{1},
			},
			Validate: func(t *testing.T, result float64) {
				require.Equal(t, 1.0, result)
			},
		},
		{
			Name: "Get Median - Even list",
			Params: struct {
				List []int
			}{
				List: []int{1, 2, 3, 4},
			},
			Validate: func(t *testing.T, result float64) {
				require.Equal(t, 2.5, result)
			},
		},
		{
			Name: "Set Median - Empty",
			Params: struct {
				List []int
			}{
				List: []int{},
			},
			Validate: func(t *testing.T, result float64) {
				require.Equal(t, 0.0, result)
			},
		},
	}

	for _, value := range testTable {
		t.Run(value.Name, func(t *testing.T) {
			result := MedianArray(value.Params.List)
			value.Validate(t, result)
		})
	}
}
