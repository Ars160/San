package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
	Stock    int     `json:"stock"`
}
