package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(dbKey string) (*gorm.DB, error) {
	uri := os.Getenv(dbKey)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: NewLogger(),
	})

	if err != nil {
		return nil, fmt.Errorf("cannot connect to Database: %w", err)
	}
	fmt.Println("database connection success!")

	return db, nil
}

func NewLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
}
