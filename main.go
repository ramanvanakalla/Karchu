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
	router := gin.Default()
	router.Use(controllers.AuthMiddleware)
	v1 := router.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("", controllers.CreateUser)
			user.POST("/auth", controllers.AuthUser)
		}
		categories := v1.Group("/categories")
		{
			categories.POST("/all", controllers.GetCategories)
			categories.POST("/n", controllers.GetCategories)
			categories.POST("", controllers.CreateCategory)
			categories.DELETE("", controllers.DeleteCategory)
			categories.PATCH("", controllers.RenameCategory)
			categories.POST("/merge", controllers.MergeCategory)
		}
		transactions := v1.Group("/transactions")
		{
			transactions.POST("", controllers.NewTransaction)
			transactions.POST("/get", controllers.GetTransactions)
			transactions.POST("/all", controllers.GetTransactionsListOfUser)
			transactions.POST("/last-n", controllers.GetLastNTransactions)
			transactions.POST("/category", controllers.GetTransactionOfCategory)
			transactions.DELETE("", controllers.DeleteTransaction)
			transactions.DELETE("/str", controllers.DeleteTransactionFromTransString)
		}
		splitTags := v1.Group("/split-tags")
		{
			splitTags.GET("", controllers.GetSplitTags)
		}
		netAmount := v1.Group("/net-amount")
		{
			netAmount.POST("/categories", controllers.GetNetMoneySpentByCategory)
		}
	}
	router.GET("/", controllers.Home)

	router.Run(":3000")
	log.Println("Everything is setup")
}
