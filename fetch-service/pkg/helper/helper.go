package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

/*
Get currency rate by all country around the world
Requirement params:

	-base (what is the base country code to convert)
	-to (exchange destination country)
*/
func GetCurrency(base string, to string) (float64, error) {
	if base == "" || to == "" {
		return 0, errors.New("params is required")
	}

	client := http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", viper.GetString("helper.currency.api"), fmt.Sprintf("base=%s&symbols=%s&amount=1", base, to)), nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return 0, err
	}

	//Parsing api response to mapping data
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	//get rates currency and convert to the float64
	rates := result["rates"].(map[string]interface{})
	return rates[to].(float64), nil
}

/*
Send response body and status code to the client or end user
*/
func HttpResponse(w http.ResponseWriter, body interface{}, status int) {
	bodyResponse, err := json.Marshal(body)
	if err != nil {
		//Set to default value if parsing process is failed
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bodyResponse)
}

/*
Bubble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in the wrong order.
This algorithm is not suitable for large data sets as its average and worst-case time complexity is quite high.
*/
func BubbleSort(arr *[]int, n int) {
	if arr != nil && (n == 0 || n == 1) {
		return
	}
	var temp int
	for i := 0; i < n-1; i++ {
		if (*arr)[i] > (*arr)[i+1] {
			temp = (*arr)[i]
			(*arr)[i] = (*arr)[i+1]
			(*arr)[i+1] = temp
		}
	}
	BubbleSort(arr, n-1)
}

/*
Calculation average of array
*/
func AvgArray(arr []int) float64 {
	if len(arr) == 0 {
		return 0.0
	}
	var total int
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}

/*
The median of a sorted array of size N is defined as the middle element when N is odd and average of middle two elements when N is even
*/
func MedianArray(arr []int) float64 {
	if len(arr) == 0 {
		return 0.0
	}
	lengthArr := len(arr)
	if lengthArr%2 != 0 {
		return float64(arr[lengthArr/2])
	}
	return float64(arr[(lengthArr-1)/2]+arr[lengthArr/2]) / 2
}
