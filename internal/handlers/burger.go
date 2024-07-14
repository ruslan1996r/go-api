package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-api/internal/types/entities"
	s "go-api/internal/types/storages"
)

type BurgerHandlers struct {
	*BasicHandler
	burgerStorage s.BurgerStorage
}

func NewBurgerHandlers(basicHandler *BasicHandler, burgerStorage s.BurgerStorage) *BurgerHandlers {
	return &BurgerHandlers{
		BasicHandler:  basicHandler,
		burgerStorage: burgerStorage,
	}
}

func (h *BurgerHandlers) InstallRoutes(r gin.IRouter) {
	g := r.Group("/burger")

	g.POST("/", h.GetAll)
	g.GET("/:id", h.GetById)
	g.POST("/create", h.Create)
	g.PUT("/order", h.OrderBurger)
	g.GET("/random", h.GetRandom)
	g.GET("/popular", h.GetPopular)
}

// Create creates a new burger
//
// @Summary Creates a new burger
// @Param 	burger body entities.BurgerRequest true "New burger payload"
// @Tags 		burger
// @Accept  json
// @Produce json
// @Failure	500	{object} ErrorResponse
// @Success 200 {integer} int "Burger ID"
// @Router /burger/create [post]
func (h *BurgerHandlers) Create(c *gin.Context) {
	var body entities.BurgerRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		h.sendInternalServerError(c, err)
		return
	}

	id, err := h.burgerStorage.Create(c, body)
	if err != nil {
		h.sendInternalServerError(c, err)
		return
	}

	h.sendOk(c, id)
}

// GetAll get all burgers with ingredients
//
// @Summary Creates a new burger
// @Param 	burger body entities.BurgerFilter true "Burgers filter"
// @Tags 		burger
// @Accept  json
// @Produce json
// @Failure	500	{object} ErrorResponse
// @Success 200 {object} []entities.Burger "Burgers array"
// @Router /burger [post]
func (h *BurgerHandlers) GetAll(c *gin.Context) {
	var body entities.BurgerFilter

	if err := c.ShouldBindJSON(&body); err != nil {
		h.sendInternalServerError(c, err)
		return
	}

	burgers, err := h.burgerStorage.GetAll(c, body)
	if err != nil {
		h.sendInternalServerError(c, err)
		return
	}

	h.sendOk(c, burgers)
}

// GetById get burger by id
//
// @Summary Get burger by ID
// @Param 	id path string true "Burger ID"
// @Tags 		burger
// @Accept  json
// @Produce json
// @Failure	500	{object} ErrorResponse
// @Failure	404	{object} string
// @Success 200 {object} entities.Burger
// @Router /burger/{id} [get]
func (h *BurgerHandlers) GetById(c *gin.Context) {
	burgerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.sendInternalServerError(c, err)
		return
	}

	burger, err := h.burgerStorage.GetById(c, burgerID)
	if err != nil {
		h.sendInternalServerError(c, err)
		return
	}

	if burger == nil {
		h.notFound(c, fmt.Sprintf("burger with ID '%s' was not found", c.Param("id")))
		return
	}

	h.sendOk(c, burger)
}

// OrderBurger order a burger
//
// @Summary Order a new burger
// @Description  Orders a burger and increments its internal state (ordered) by 1, which represents its popularity
// @Param 	burger body entities.BurgerOrderRequest true "New burger payload"
// @Tags 		burger
// @Accept  json
// @Produce json
// @Failure	500	{object} ErrorResponse
// @Success 200 {boolean} bool
// @Router /burger/order [put]
func (h *BurgerHandlers) OrderBurger(c *gin.Context) {
	var body entities.BurgerOrderRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		h.sendInternalServerError(c, err)
		return
	}

	ok, err := h.burgerStorage.OrderBurger(c, body)
	if err != nil {
		h.sendInternalServerError(c, err)
		return
	}

	h.sendOk(c, ok)
}

// GetRandom get random burger
//
// @Summary Get random burger
// @Tags 		burger
// @Accept  json
// @Produce json
// @Failure	500	{object} ErrorResponse
// @Success 200 {object} entities.Burger
// @Router /burger/random [get]
func (h *BurgerHandlers) GetRandom(c *gin.Context) {
	burger, err := h.burgerStorage.GetRandom(c)
	if err != nil {
		h.sendInternalServerError(c, err)
		return
	}

	h.sendOk(c, burger)
}

// GetPopular get popular burgers
//
// @Summary Get popular burgers
// @Description  Get top 5 burgers. Popularity is based on number of orders (ordered field)
// @Tags 		burger
// @Accept  json
// @Produce json
// @Failure	500	{object} ErrorResponse
// @Success 200 {object} []entities.Burger
// @Router /burger/popular [get]
func (h *BurgerHandlers) GetPopular(c *gin.Context) {
	burger, err := h.burgerStorage.GetPopular(c)
	if err != nil {
		h.sendInternalServerError(c, err)
		return
	}

	h.sendOk(c, burger)
}
