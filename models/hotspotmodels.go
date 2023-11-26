package models

import "gorm.io/gorm"

type Hotspot struct {
	gorm.Model
	ID   int     `gorm:"primaryKey"`
	Name string  `json:"title"`
	Lang float64 `json:"lang"`
	Long float64 `json:"long"`
}
