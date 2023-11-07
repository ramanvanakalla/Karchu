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
	r.POST("/getCategories", controllers.GetCategories)
	r.GET("/splitTags", controllers.GetSplitTags)
	r.POST("/createUser", controllers.CreateUser)
	r.POST("/createCategory", controllers.CreateCategory)
	r.POST("/newTransaction", controllers.NewTransaction)
	r.Run(":3000")
	log.Println("Everything is setup")
}
