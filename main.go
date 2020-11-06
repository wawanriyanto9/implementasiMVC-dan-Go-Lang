package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/imadedwisuryadinata/implementasiMvcGolang/app/config"
	"github.com/imadedwisuryadinata/implementasiMvcGolang/app/controller"
	"github.com/imadedwisuryadinata/implementasiMvcGolang/app/middleware"
)

func main() {
	db := config.DBInit()
	accountController := controller.AccountController{DB: db}
	transactionController := controller.TransactionController{DB: db}

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/api/v1/account/add", accountController.CreateAccount)
	router.POST("/api/v1/login", accountController.Login)
	router.GET("/api/v1/account", middleware.Auth, accountController.CreateAccount)

	router.POST("/api/v1/transfer", middleware.Auth, transactionController.Transfer)
	router.POST("/api/v1/deposit", middleware.Auth, transactionController.Deposit)
	router.POST("/api/v1/withdraw", middleware.Auth, transactionController.Withdraw)

	router.GET("/", accountController.CreateAccount)
	router.Run(":8000")
}
