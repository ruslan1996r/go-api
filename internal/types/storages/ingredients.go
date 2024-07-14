package storages

import (
	"github.com/gin-gonic/gin"
	"go-api/internal/types/entities"
)

type IngredientsStorage interface {
	Search(ctx *gin.Context, ingredients []string) ([]*entities.Burger, error)
}
