package models

import (
	"WowrackCustomerAppRestfulAPI/database"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Status       bool       `json:"status"`
	ArticlesSize int        `json:"articles_size"`
	Articles     []Articles `json:"articles"`
}

type Articles struct {
	gorm.Model
	ID          int    `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

func (articles *Articles) CreateUserRecord() error {
	result := database.GlobalDB.Create(&articles)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
