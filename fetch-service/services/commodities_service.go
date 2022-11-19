package services

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/models"
	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/pkg/database"
	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/pkg/helper"
)

type CommoditiesService struct {
	Storage database.CommoditiesRepositories
}

func (service *CommoditiesService) Get(w http.ResponseWriter, r *http.Request) {
	commodities, err := service.Storage.Get()
	if err != nil {
		httpResponse := &models.HttpResponseMessage{
			Message: "Get data Failed",
		}
		helper.HttpResponse(w, httpResponse, http.StatusBadRequest)
		return
	}
	rateUsd, err := helper.GetCurrency("USD", "IDR")
	if err != nil {
		log.Println(err)
		httpResponse := &models.HttpResponseMessage{
			Message: "Get data Failed",
		}
		helper.HttpResponse(w, httpResponse, http.StatusBadRequest)
		return
	}

	var waitGroup sync.WaitGroup
	for index, value := range commodities {
		waitGroup.Add(1)
		go func(idx int, price string) {
			priceIdr, err := strconv.ParseFloat(price, 64)
			if err == nil {
				priceUsd := priceIdr / rateUsd
				commodities[idx].PriceUsd = fmt.Sprintf("%.2f", priceUsd)
			}
			waitGroup.Done()
		}(index, value.Price)
	}
	waitGroup.Wait()
	helper.HttpResponse(w, commodities, http.StatusOK)
}

func (service *CommoditiesService) GetAggregate(w http.ResponseWriter, r *http.Request) {
	commodities, err := service.Storage.Get()
	if err != nil {
		log.Println(err)
		httpResponse := &models.HttpResponseMessage{
			Message: "Get data Failed",
		}
		helper.HttpResponse(w, httpResponse, http.StatusBadRequest)
		return
	}

	//grouping commodities by province and weekly
	provinceCommodities := make(map[string]map[string]map[string][]int)
	for _, value := range commodities {
		tglParsed, err := time.Parse(time.RFC3339, value.TglParsed)
		if err != nil {
			continue
		}
		//Extract year & week number as a weekly group key
		yearNumber, weekNumber := tglParsed.ISOWeek()
		weeklyKey := fmt.Sprintf("%d-%d", yearNumber, weekNumber)

		//Parsing commodity price from string to int
		commodityPrice, err := strconv.Atoi(value.Price)
		if err != nil {
			continue
		}

		//Parsing commodity size from string to int
		commoditySize, err := strconv.Atoi(value.Size)
		if err != nil {
			continue
		}

		if provinceCommodities[value.AreaProvinsi] == nil {
			weeklyCommodities := make(map[string]map[string][]int)
			//Initial daily list as a slice int
			dailycommodities := make(map[string][]int)
			dailycommodities["price"] = append(dailycommodities["price"], commodityPrice)
			dailycommodities["size"] = append(dailycommodities["size"], commoditySize)
			//set province & weekly first item
			weeklyCommodities[weeklyKey] = dailycommodities
			provinceCommodities[value.AreaProvinsi] = weeklyCommodities
		} else {
			if provinceCommodities[value.AreaProvinsi][weeklyKey] == nil {
				//Initial daily list as a slice int
				dailycommodities := make(map[string][]int)
				dailycommodities["price"] = append(dailycommodities["price"], commodityPrice)
				dailycommodities["size"] = append(dailycommodities["size"], commoditySize)
				//add week number to the province items
				provinceCommodities[value.AreaProvinsi][weeklyKey] = dailycommodities
			} else {
				//add price & size to the weekly items
				provinceCommodities[value.AreaProvinsi][weeklyKey]["price"] = append(provinceCommodities[value.AreaProvinsi][weeklyKey]["price"], commodityPrice)
				provinceCommodities[value.AreaProvinsi][weeklyKey]["size"] = append(provinceCommodities[value.AreaProvinsi][weeklyKey]["size"], commoditySize)
			}
		}
	}

	aggCommodities := make([]map[string]interface{}, 0)
	for provinceKey, province := range provinceCommodities {
		provinceData := make(map[string]interface{})
		provinceWeeklyData := make([]map[string]interface{}, 0)
		for weeklyKey, week := range province {
			weeklyData := make(map[string]interface{})
			//Set & sorting price list
			priceList := week["price"]
			helper.BubbleSort(&priceList, len(priceList))
			//Set & sorting size list
			sizeList := week["size"]
			helper.BubbleSort(&sizeList, len(sizeList))
			//Split weekly key to year & weekly number
			weeklyKeyArr := strings.Split(weeklyKey, "-")

			weeklyData["year"] = weeklyKeyArr[0]
			weeklyData["week"] = weeklyKeyArr[1]
			weeklyData["price_max"] = priceList[len(priceList)-1]
			weeklyData["price_min"] = priceList[0]
			weeklyData["price_avg"] = helper.AvgArray(priceList)
			weeklyData["price_median"] = helper.MedianArray(priceList)
			weeklyData["size_max"] = sizeList[len(sizeList)-1]
			weeklyData["size_min"] = sizeList[0]
			weeklyData["size_avg"] = helper.AvgArray(sizeList)
			weeklyData["size_median"] = helper.MedianArray(sizeList)
			provinceWeeklyData = append(provinceWeeklyData, weeklyData)
		}
		provinceData["province"] = provinceKey
		provinceData["data"] = provinceWeeklyData
		aggCommodities = append(aggCommodities, provinceData)
	}

	helper.HttpResponse(w, aggCommodities, http.StatusOK)
}
