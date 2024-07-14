package storage

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go-api/internal/types/entities"
	"gorm.io/gorm"
)

type IngredientsStorage struct {
	db *gorm.DB
}

func NewIngredientsStorage(db *gorm.DB) *IngredientsStorage {
	return &IngredientsStorage{
		db: db,
	}
}

func (s *IngredientsStorage) Search(ctx *gin.Context, ingredients []string) ([]*entities.Burger, error) {
	var burgers []*entities.Burger

	if len(ingredients) == 0 {
		return burgers, nil
	}

	subQuery := s.db.Model(&entities.Ingredient{}).
		Select("burger_id").
		Where("name IN ?", ingredients).
		Group("burger_id").
		Having("COUNT(DISTINCT name) = ?", len(ingredients))

	result := s.db.Where("id IN (?)", subQuery).
		Preload("Ingredients").
		Find(&burgers).
		WithContext(ctx)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch burgers by ingredients: %w", result.Error)
	}

	return burgers, nil
}
