package main

import (
	"Karchu/initializers"
	"Karchu/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	initializers.LoadEnvVariables()
}

type Transaction struct {
	gorm.Model
	UserId           uint
	Amount           int
	Time             time.Time
	Description      string
	SplitTag         string
	CategoryMappings []models.CategoryTransactionMapping `gorm:"foreignKey:TransactionId"`
}

func main() {
	sourceDB, err := gorm.Open(postgres.Open(os.Getenv("DB_URL_PROD")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer sourceDB.DB()
	fmt.Println("Connected to source DB")
	destinationDB, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer destinationDB.DB()
	fmt.Println("Connected to destination DB")
	migrateUsersTable(sourceDB, destinationDB)
	migrateCategoriesTable(sourceDB, destinationDB)
	migrateTransactionsTable(sourceDB, destinationDB)
	migrateCategoryTransactionMappingTable(sourceDB, destinationDB)
}

func migrateUsersTable(sourceDB *gorm.DB, destinationDB *gorm.DB) {
	// Auto Migrate the User model to the destination database
	err := destinationDB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
	// Query data from the source database
	var users []models.User
	sourceDB.Find(&users)

	// Insert data into the destination database
	for _, user := range users {
		destinationDB.Create(&user)
	}
	fmt.Println("Data migration successful for users table")
}

func migrateCategoriesTable(sourceDB *gorm.DB, destinationDB *gorm.DB) {
	// Auto Migrate the Category model to the destination database
	err := destinationDB.AutoMigrate(&models.Category{})
	if err != nil {
		log.Fatal(err)
	}
	// Query data from the source database
	var categories []models.Category
	sourceDB.Find(&categories)

	// Insert data into the destination database
	for _, category := range categories {
		destinationDB.Create(&category)
	}
	fmt.Println("Data migration successful for category table")
}

func migrateTransactionsTable(sourceDB *gorm.DB, destinationDB *gorm.DB) {
	// Auto Migrate the Transaction model to the destination database
	err := destinationDB.AutoMigrate(&Transaction{})
	if err != nil {
		log.Fatal(err)
	}
	// Query data from the source database
	var transactions []Transaction
	sourceDB.Find(&transactions)

	// Insert data into the destination database
	for _, transaction := range transactions {
		destinationDB.Create(&transaction)
	}
	fmt.Println("Data migration successful for transaction table")
}

func migrateCategoryTransactionMappingTable(sourceDB *gorm.DB, destinationDB *gorm.DB) {
	// Auto Migrate the CategoryTransactionMapping model to the destination database
	err := destinationDB.AutoMigrate(&models.CategoryTransactionMapping{})
	if err != nil {
		log.Fatal(err)
	}
	// Query data from the source database
	var transactions []models.Transaction
	sourceDB.Find(&transactions)

	// Insert data into the destination database
	for _, transaction := range transactions {
		categoryTransactionMap := models.CategoryTransactionMapping{TransactionId: transaction.ID, CategoryId: transaction.CategoryId}
		destinationDB.Create(&categoryTransactionMap)
	}
	fmt.Println("Data migration successful for categoryTransactionMapping table")
}
