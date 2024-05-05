package main

import (
	"back"
	"back/pkg/handler"
	"back/pkg/repository"
	"back/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(back.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("configs")
	viper.SetConfigName("config.yml")
	return viper.ReadInConfig()
}
