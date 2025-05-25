package models

type Product struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	CategoryID uint    `json:"category_id"`
	UserID     uint    `json:"user_id"`
}
