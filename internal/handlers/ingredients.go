package handlers

import (
	"strings"

	"github.com/gin-gonic/gin"
	s "go-api/internal/types/storages"
)

type IngredientsHandlers struct {
	*BasicHandler
	ingredientsStorage s.IngredientsStorage
}

func NewIngredientsHandlers(basicHandler *BasicHandler, ingredientsStorage s.IngredientsStorage) *IngredientsHandlers {
	return &IngredientsHandlers{
		BasicHandler:       basicHandler,
		ingredientsStorage: ingredientsStorage,
	}
}

func (h *IngredientsHandlers) InstallRoutes(r gin.IRouter) {
	g := r.Group("/ingredients")

	g.GET("/", h.Search)
}

// Search get a list of burgers with selected ingredients
//
// @Summary Get a list of burgers containing the selected ingredients
// @Param 	search query string false "Ingredients with a lowercase letter separated by commas"
// @Tags 		burger
// @Accept  json
// @Produce json
// @Failure	500	{object} ErrorResponse
// @Success 200 {object} []entities.Burger
// @Router /ingredients [get]
func (h *IngredientsHandlers) Search(c *gin.Context) {
	limitParam := c.DefaultQuery("search", "")

	var ingredients []string

	if limitParam != "" {
		ingredients = strings.Split(limitParam, ",")
	}

	burgers, err := h.ingredientsStorage.Search(c, ingredients)
	if err != nil {
		h.sendInternalServerError(c, err)
		return
	}

	h.sendOk(c, burgers)
}
