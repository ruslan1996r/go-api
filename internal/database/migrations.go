package database

import (
	"go-api/internal/types/entities"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.Burger{},
		&entities.Ingredient{},
	)

	if err != nil {
		return err
	}

	return nil
}
