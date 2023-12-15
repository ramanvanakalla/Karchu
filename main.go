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
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      https://karchu.onrender.com
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":3000")
	log.Println("Everything is setup")
}
