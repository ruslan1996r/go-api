package storages

import (
	"github.com/gin-gonic/gin"
	"go-api/internal/types/entities"
)

type BurgerStorage interface {
	Create(ctx *gin.Context, burgerReq entities.BurgerRequest) (id int, err error)
	GetById(ctx *gin.Context, id int) (*entities.Burger, error)
	GetAll(ctx *gin.Context, filter entities.BurgerFilter) ([]*entities.Burger, error)
	OrderBurger(ctx *gin.Context, order entities.BurgerOrderRequest) (int, error)
	GetRandom(ctx *gin.Context) (*entities.Burger, error)
	GetPopular(ctx *gin.Context) ([]*entities.Burger, error)
}
