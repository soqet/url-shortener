package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/soqet/configjson"
	"net/http"
	"url-shortener/internal/api"
	database "url-shortener/internal/db"
)

type Config struct {
	ApiUrl string `json:"apiUrl"`
	Port   int    `json:"port"`
}

func main() {
	config := new(Config)
	configjson.ReadConfigFile("./config.json", config)
	db, err := database.NewDb("./urls.db")
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	s := router.Host(config.ApiUrl).Subrouter()
	api.Init(s, *db, config.ApiUrl)
	http.Handle("/", router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
