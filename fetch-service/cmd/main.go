package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/pkg/config"
	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/pkg/database"
	"github.com/briankliwon/efishery-recruitment-backend/fetch-service/services"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// Initial commodities service
var storage = &database.CommoditiesStorage{}
var service = services.CommoditiesService{Storage: storage}

func main() {
	//Setting Up Viper configuration
	viperConfig := config.Config{
		Name: "config",
		Type: "json",
		Path: "pkg/config",
	}
	err := viperConfig.Init()
	if err != nil {
		log.Fatal(err)
	}

	//Initial routes of services
	routes := mux.NewRouter()
	commoditiesRouter := routes.PathPrefix("/commodities").Subrouter()
	commoditiesRouter.HandleFunc("/", service.Get).Methods("GET")
	commoditiesRouter.HandleFunc("/aggregate", service.GetAggregate).Methods("GET")

	//http configuration
	serverConfig := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("server.port")),
		Handler: routes,
	}

	log.Println(viper.GetString("server.port"))
	err = serverConfig.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
