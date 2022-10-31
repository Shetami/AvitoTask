package main

import (
	"AvitoTask"
	"AvitoTask/internal/handler"
	"AvitoTask/internal/repository"
	"AvitoTask/internal/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}
	db, err := repository.NewPSQL(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		NameDB:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed initialize db: %s", err.Error())
	}
	rep := repository.NewRepository(db)
	services := service.NewService(rep)
	handlers := response.NewHandler(services)

	server := new(avitotask.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
