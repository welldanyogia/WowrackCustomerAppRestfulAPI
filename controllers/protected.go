package controllers

import (
	"WowrackCustomerAppRestfulAPI/database"
	"WowrackCustomerAppRestfulAPI/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Profile is a controller function that retrieves the user profile from the database
// based on the email provided in the authorization middleware.
// It returns a 404 status code if the user is not found,
// and a 500 status code if an error occurs while retrieving the user profile.

func Profile(c *gin.Context) {
	// Initialize a user model
	var user models.User
	// Get the email from the authorization middleware
	email, _ := c.Get("email")
	// Query the database for the user
	result := database.GlobalDB.Where("email = ?", email.(string)).First(&user)
	// If the user is not found, return a 404 status code
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"Error": "User Not Found",
		})
		c.Abort()
		return
	}
	// If an error occurs while retrieving the user profile, return a 500 status code
	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get User Profile",
		})
		c.Abort()
		return
	}
	// Set the user's password to an empty string
	user.Password = ""
	// Return the user profile with a 200 status code
	c.JSON(200, user)
}
func Article(c *gin.Context) {
	var articles []models.Articles
	var user models.User

	// Get the email from the authorization middleware
	email, _ := c.Get("email")
	result := database.GlobalDB.Where("email = ?", email.(string)).First(&user)

	// Retrieve all articles from the database
	result = database.GlobalDB.Find(&articles)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get Articles",
		})
		c.Abort()
		return
	}

	count := int64(len(articles))
	c.JSON(200, gin.H{
		"status":       true,
		"article_size": count,
		"articles":     articles,
	})
}
func Hotspot(c *gin.Context) {
	var hotspots []models.Hotspot
	var user models.User

	// Get the email from the authorization middleware
	email, _ := c.Get("email")
	result := database.GlobalDB.Where("email = ?", email.(string)).First(&user)

	// Retrieve all articles from the database
	result = database.GlobalDB.Find(&hotspots)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get Hotspot",
		})
		c.Abort()
		return
	}

	count := int64(len(hotspots))
	c.JSON(200, gin.H{
		"status":       true,
		"hotspot_size": count,
		"hotspot":      hotspots,
	})
}
