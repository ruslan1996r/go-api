package storage

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go-api/internal/types/entities"
	"gorm.io/gorm"
)

type BurgerStorage struct {
	db *gorm.DB
}

func NewBurgerStorage(db *gorm.DB) *BurgerStorage {
	return &BurgerStorage{
		db: db,
	}
}

func (s *BurgerStorage) Create(ctx *gin.Context, burgerReq entities.BurgerRequest) (id int, err error) {
	createdAt := time.Now()

	newBurger := &entities.Burger{
		Name:        burgerReq.Name,
		CreatedAt:   &createdAt,
		Ingredients: burgerReq.Ingredients,
	}

	result := s.db.Create(newBurger).WithContext(ctx)
	if result.Error != nil {
		return 0, fmt.Errorf("failed to create task: %w", result.Error)
	}

	return newBurger.ID, nil
}

func (s *BurgerStorage) GetById(ctx *gin.Context, burgerID int) (*entities.Burger, error) {
	var burger entities.Burger

	result := s.db.Preload("Ingredients").
		First(&burger, burgerID).
		WithContext(ctx)

	if result.Error != nil {
		return nil, nil
	}

	return &burger, nil
}

func (s *BurgerStorage) GetAll(ctx *gin.Context, filter entities.BurgerFilter) ([]*entities.Burger, error) {
	var burgers []*entities.Burger

	query := s.db.WithContext(ctx)

	if filter.Name != nil {
		query = query.Where("name LIKE ?", "%"+*filter.Name+"%")
	}

	result := query.Preload("Ingredients").Find(&burgers)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch burgers: %w", result.Error)
	}

	return burgers, nil
}

func (s *BurgerStorage) GetRandom(ctx *gin.Context) (*entities.Burger, error) {
	var burger entities.Burger

	result := s.db.WithContext(ctx).
		Order("RANDOM()").
		Limit(1).
		Preload("Ingredients").
		Find(&burger)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch random burger: %w", result.Error)
	}

	return &burger, nil
}

func (s *BurgerStorage) OrderBurger(ctx *gin.Context, order entities.BurgerOrderRequest) (int, error) {
	var burger entities.Burger

	result := s.db.First(&burger, order.BurgerID).WithContext(ctx)
	if result.Error != nil {
		return 0, fmt.Errorf("failed to find burger: %w", result.Error)
	}

	ordered := burger.Ordered + 1
	burger.Ordered = ordered

	errUpdate := s.db.Save(&burger).Error
	if errUpdate != nil {
		return 0, fmt.Errorf("failed to update burger: %w", errUpdate)
	}

	return ordered, nil
}

func (s *BurgerStorage) GetPopular(ctx *gin.Context) ([]*entities.Burger, error) {
	var burgers []*entities.Burger

	result := s.db.WithContext(ctx).
		Order("ordered DESC").
		Limit(5).
		Preload("Ingredients").
		Find(&burgers)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch top ordered burgers: %w", result.Error)
	}

	return burgers, nil
}
