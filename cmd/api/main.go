package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mafuyuk/ddd-go-api-template/application"
	"github.com/mafuyuk/ddd-go-api-template/config"
	"github.com/mafuyuk/ddd-go-api-template/handler"
	"github.com/mafuyuk/ddd-go-api-template/infrastructure"
	"github.com/mafuyuk/ddd-go-api-template/infrastructure/db"
)

func main() {
	// Parse runtime options
	confPath := flag.String("c", "config", "config path")
	flag.Parse()

	// Setting Config
	conf := &config.Config{}
	log.Println(*confPath)
	if err := config.New(conf, *confPath); err != nil {
		panic(err)
	}

	// Initialize db connection
	dbmClient, err := db.NewMySQL(&conf.DBMaster)
	if err != nil {
		panic(err)
	}
	defer dbmClient.Close()

	dbsClient, err := db.NewMySQL(&conf.DBSlave)
	if err != nil {
		panic(err)
	}
	defer dbsClient.Close()

	// Initialize repository
	userRepo := infrastructure.NewUserRepository(dbmClient, dbsClient)

	// Initialize application service
	userService := application.NewUserService(userRepo)

	// Run App server
	server := handler.New(conf.Server.Port, &handler.Services{
		UserService: userService,
	})
	log.Println("Start server")
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("Failed ListenAndServe. err: %v", err))
	}
}
