package entities

type Ingredient struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	BurgerID int    `json:"burger_id"`
}
