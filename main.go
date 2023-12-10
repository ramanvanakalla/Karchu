package main

import (
	"Karchu/controllers"
	"Karchu/initializers"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(controllers.AuthMiddleware)
	// User
	router.POST("/v1/user", controllers.CreateUser)
	// Categories
	router.POST("/v1/categories/all", controllers.GetCategories)
	router.POST("/v1/categories/n", controllers.GetCategories)
	router.POST("/v1/categories", controllers.CreateCategory)
	router.DELETE("/v1/categories", controllers.DeleteCategory)
	// Transactions
	router.POST("/v1/transactions", controllers.NewTransaction)
	router.POST("/v1/transactions/all", controllers.GetTransactionsListOfUser)
	router.POST("/v1/transactions/last-n", controllers.GetLastNTransactions)
	router.POST("/v1/transactions/category", controllers.GetTransactionOfCategory)
	router.DELETE("/v1/transactions", controllers.DeleteTransaction)
	router.DELETE("/v1/transactions/str", controllers.DeleteTransactionFromTransString)
	// SplitTags
	router.GET("/v1/split-tags", controllers.GetSplitTags)

	// Net
	router.POST("/v1/net-amount/categories", controllers.GetNetMoneySpentByCategory)
	//Home
	router.GET("/", controllers.Home)

	router.Run(":3000")
	log.Println("Everything is setup")
}
