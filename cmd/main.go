package main

import (
	config "musiclib/internal/config"
	logger "musiclib/internal/infrastructure/logger"
	storage "musiclib/internal/infrastructure/storage"
	_ "musiclib/internal/models"
	repository "musiclib/internal/repository"
	service "musiclib/internal/service"
	transport "musiclib/internal/transport/api"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/v2"
	_ "github.com/swaggo/swag"

	_ "musiclib/docs"

	_ "github.com/sirupsen/logrus"
)

// @title Music Lib
// @version 1.0
// @description Sever for manage music lib.

// @host localhost:8080
// @BasePath /api/v1

func main() {

	log, _, err := logger.InitLogger()
	if err != nil {
		os.Exit(1)
	}

	connectionString, err := config.LoadConfig()
	if err != nil {
		log.Debug("Failed to load configuration file")
		log.Debug(connectionString)
		log.Info("Bad configuration file")
		os.Exit(1)
	}

	db, err := storage.InitDatabase(log, connectionString)
	if err != nil {
		log.Debug("Error initializing database: ", err)
		os.Exit(126)
	}

	err = storage.RunMigrations(log, connectionString)
	if err != nil {
		os.Exit(126)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)

	router := transport.NewRouter(log, service)
	router.Mux.HandleFunc("/swagger/*", httpSwagger.WrapHandler)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		router.Log.Fatal("Error starting server: ", err)
	}
}
