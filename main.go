package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"mnc-stage2/src"
	"mnc-stage2/src/config"
	controller "mnc-stage2/src/handler"
	"mnc-stage2/src/repository"
	"mnc-stage2/src/service"
)

func main() {
	cfg := config.LoadConfig()
	db := cfg.ConnectDB()
	cfg.GenerateMocks()

	userRepository := repository.NewUserRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	userService := service.NewUserService(userRepository, transactionRepository, db)
	userController := controller.NewUserController(userService)
	r := gin.Default()

	src.RegisterRoutes(r, userController)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start the server: ", err)
	}

}
