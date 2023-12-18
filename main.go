package main

import (
	"Karchu/controllers"
	"Karchu/initializers"
	"log"

	_ "Karchu/docs"

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
			transactions.POST("", controllers.NewTransactionV2)
			transactions.POST("/get", controllers.GetTransactionsV2)
			transactions.POST("/all", controllers.GetTransactionsListOfUserV2)
			transactions.POST("/last-n", controllers.GetLastNTransactionsV2)
			transactions.POST("/category", controllers.GetTransactionOfCategoryV2)
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
		categories := v2.Group("/categories")
		{
			categories.POST("/merge", controllers.MergeCategoryV2)
		}
		transactions := v2.Group("/transactions")
		{
			transactions.POST("", controllers.NewTransactionV2)
			transactions.POST("/get", controllers.GetTransactionsV2)
			transactions.POST("/all", controllers.GetTransactionsListOfUserV2)
			transactions.POST("/last-n", controllers.GetLastNTransactionsV2)
			transactions.POST("/category", controllers.GetTransactionOfCategoryV2)
			transactions.DELETE("", controllers.DeleteTransaction)
			transactions.DELETE("/str", controllers.DeleteTransactionFromTransString)
		}
		netAmount := v2.Group("/net-amount")
		{
			netAmount.POST("/categories", controllers.GetNetMoneySpentByCategory2)
		}
	}
	router.GET("/", controllers.Home)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":3000")
	log.Println("Everything is setup")
}
