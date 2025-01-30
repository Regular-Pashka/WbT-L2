package main

import (
	"fmt"
	"net/http"
)

/* 
	Создать handler в папке handlers
	там обрабатывать запрос через стороннюю функцию
	
*/

func main() {
	/* viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	} */
	// db, err := 
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", simpleResponse)

	port := "8000"

	http.ListenAndServe(":" + port, mux)



}