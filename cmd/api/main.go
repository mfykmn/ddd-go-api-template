package main

import (
	"fmt"
	"log"

	"github.com/mafuyuk/ddd-go-api-template/application"
	"github.com/mafuyuk/ddd-go-api-template/handler"
	"github.com/mafuyuk/ddd-go-api-template/infrastructure"
)

func main() {
	// Initialize repository
	userRepo := infrastructure.NewUserRepository()

	// Initialize application service
	userService := application.NewUserService(userRepo)

	// Run App server
	server := handler.New(":8080", &handler.Services{
		UserService: userService,
	})
	log.Println("Start server")
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("Failed ListenAndServe. err: %v", err))
	}
}
