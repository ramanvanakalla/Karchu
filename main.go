package main

import (
	"Karchu/controllers"
	"Karchu/initializers"
	"log"

	_ "Karchu/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

// @title           Karchu API
// @version         1.0
// @description     All APIs related to Karchu.

// @host      karchu.onrender.com
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth
func main() {
	router := gin.Default()
	router.Use(cors.Default())
	v1 := router.Group("/v1")
	{
		v1.Use(controllers.AuthMiddleware)
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
			transactions.POST("/category", controllers.GetTransactionStringsOfCategory)
			transactions.POST("/categories", controllers.GetTransactionsOfCategory)
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
	v2 := router.Group("/v2")
	{
		v2.Use(controllers.AuthMiddleware)
		friends := v2.Group("/friends")
		{
			friends.POST("", controllers.CreateFriend)
			friends.POST("/get", controllers.GetFriends)

		}
		SplitTransaction := v2.Group("/split-transaction")
		{
			SplitTransaction.POST("", controllers.SplitTransaction)
			SplitTransaction.DELETE("", controllers.DeleteSplitTransaction)
			SplitTransaction.DELETE("/str", controllers.DeleteSplitTransactionString)
			SplitTransaction.POST("/one", controllers.SplitTransactionWithOneFriend)
			SplitTransaction.POST("/splits", controllers.GetSplitTransactions)
			SplitTransaction.POST("/unsettled-splits", controllers.GetUnSettledSplitTransactions)
			SplitTransaction.POST("/settled-splits", controllers.GetSettledSplitTransactions)
		}
		settleSplit := v2.Group("/settle")
		{
			settleSplit.POST("", controllers.SettleSplitTransaction)
			settleSplit.POST("/str", controllers.SettleSplitTransactionString)
			settleSplit.DELETE("", controllers.UnSettleSplitTransaction)
			settleSplit.DELETE("/str", controllers.UnSettleSplitTransactionString)
			settleSplit.POST("/friend", controllers.SettleSplitsOfFriend)
		}
		transactionAndSplit := v2.Group("/trans-split-with-one")
		{
			transactionAndSplit.POST("", controllers.CreateTransactionAndSplitWithOne)
		}
		moneyLent := v2.Group("/money-lent")
		{
			moneyLent.POST("", controllers.MoneyLentToFriend)
		}
	}
	router.GET("/", controllers.Home)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":3000")
	log.Println("Everything is setup")
}
