// main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gin-json-response/models"  // นำเข้า models ที่เราเปลี่ยนเป็น package models
)

var db *gorm.DB
var err error

// Initialize the database
func initDB() {
	// Open the SQLite database
	db, err = gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		fmt.Println("Error opening database: ", err)
	}

	// Migrate the schema (create tables)
	db.AutoMigrate(&models.User{}, &models.Post{})

	// Create some sample data if users are empty
	if db.Table("users").Where("name = ?", "Alice").First(&models.User{}).RecordNotFound() {
		user := models.User{Name: "Alice"}
		db.Create(&user)
		db.Create(&models.Post{Title: "Post 1", Body: "This is the first post", UserID: user.ID})
		db.Create(&models.Post{Title: "Post 2", Body: "This is the second post", UserID: user.ID})
	}

	if db.Table("users").Where("name = ?", "Bob").First(&models.User{}).RecordNotFound() {
		user := models.User{Name: "Bob"}
		db.Create(&user)
		db.Create(&models.Post{Title: "Bob's First Post", Body: "Bob's post content", UserID: user.ID})
	}
}

// Fetch all users with related posts
func getUsers(c *gin.Context) {
	var users []models.User
	// Preload posts (this will fetch users along with their related posts)
	if err := db.Preload("Posts").Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": "Unable to fetch users"})
		return
	}
	// Return the users as JSON
	c.JSON(200, users)
}

func main() {
	// Initialize the database
	initDB()

	// Set up Gin router
	r := gin.Default()

	// Define routes
	r.GET("/users", getUsers)

	// Start the server
	r.Run(":8080")
}
