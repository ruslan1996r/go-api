package entities

import "time"

type Burger struct {
	ID          int          `gorm:"primaryKey" json:"id"`
	Name        string       `json:"name"`
	Ordered     int          `gorm:"default:0" json:"ordered"`
	CreatedAt   *time.Time   `json:"created_at"`
	Ingredients []Ingredient `gorm:"foreignKey:BurgerID" json:"ingredients"`
}

type BurgerRequest struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

type BurgerFilter struct {
	Name *string `json:"name"`
}

type BurgerOrderRequest struct {
	BurgerID int `json:"burger_id"`
}
