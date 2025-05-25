package models

type Outfit struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	UserID   uint      `json:"user_id"`
	Products []Product `gorm:"many2many:outfit_products;" json:"products"`
}
