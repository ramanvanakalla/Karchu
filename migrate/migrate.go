package main

import (
	"Karchu/initializers"
	"Karchu/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}
func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Category{}, &models.CategoryTransactionMapping{})
}
