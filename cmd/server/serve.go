package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "go-api/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-api/internal/database"
	"go-api/internal/handlers"
	"go-api/internal/storage"
)

const dbUri = "DB_URI" //

func CreateApp() (*gin.Engine, error) {
	db, err := database.Init(dbUri)
	if err != nil {
		return nil, fmt.Errorf("failed to run database: %w", err)
	}

	err = database.RunMigration(db)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	r := gin.New()

	burgerStorage := storage.NewBurgerStorage(db)
	ingredientsStorage := storage.NewIngredientsStorage(db)

	basicHandler := handlers.NewBasicHandler()

	burgerHandlers := handlers.NewBurgerHandlers(basicHandler, burgerStorage)
	ingredientsHandlers := handlers.NewIngredientsHandlers(basicHandler, ingredientsStorage)

	// Handlers
	r.GET("/health", handlers.Health)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	burgerHandlers.InstallRoutes(r)
	ingredientsHandlers.InstallRoutes(r)

	return r, nil
}
