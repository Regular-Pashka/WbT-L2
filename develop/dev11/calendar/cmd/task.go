package main

import (
	"net/http"
	"os"

	"github.com/Regular-Pashka/WbT-L2/develop/dev11/calendar/internal/handler"
	"github.com/Regular-Pashka/WbT-L2/develop/dev11/calendar/internal/repository"
	"github.com/Regular-Pashka/WbT-L2/develop/dev11/calendar/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.database"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	port := "8000"

	http.ListenAndServe(":"+port, mux)

}
