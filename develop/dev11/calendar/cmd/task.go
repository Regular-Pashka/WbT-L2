package main

import (
	"net/http"

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
	if err := InitConfig(); err != nil {
		logrus.Fatalf("Error reading config file, %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		Password: viper.GetString("db.password"), // тут опционально можно спрятать в .env файл в переменную окружения DB_PASSWORD и потом с помощью godotenv считать
		DBName:   viper.GetString("db.database"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	// mux := http.NewServeMux()
	// mux.HandleFunc("/hello")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	http.HandleFunc("/create_event", handlers.CreateEvent)

	port := "8000"


	logrus.Println("Server started on port " + port)
	logrus.Fatal(http.ListenAndServe(":" + port, nil))
}

func InitConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	return viper.ReadInConfig()
}