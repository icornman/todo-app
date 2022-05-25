package main

import (
	"github.com/spf13/viper"
	"log"
	todo "todo-app"
	"todo-app/pkg/handler"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	err := srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
