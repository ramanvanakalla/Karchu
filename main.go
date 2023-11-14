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
	r := gin.Default()
	// User
	r.POST("/v1/user", controllers.CreateUser)
	// Categories
	r.POST("/v1/:route/categories/all", controllers.GetCategories)
	r.POST("/v1/categories", controllers.CreateCategory)

	// Transactions
	r.POST("/v1/transactions", controllers.NewTransaction)

	// SplitTags
	r.GET("/v1/split-tags", controllers.GetSplitTags)
	//Home
	r.GET("/", controllers.Home)

	r.Run(":3000")
	log.Println("Everything is setup")
}
