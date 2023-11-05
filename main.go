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
	log.Println("Everything is setup")
	r := gin.Default()
	r.GET("/categories", controllers.GetCategories)
	r.GET("/splitTags", controllers.GetSplitTags)
	r.POST("/createUser", controllers.CreateUser)
	r.POST("/createCategory", controllers.CreateCategory)
	r.Run()
}
