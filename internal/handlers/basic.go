package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func Health(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}

type BasicHandler struct{}

func NewBasicHandler() *BasicHandler {
	return &BasicHandler{}
}

func (h *BasicHandler) sendNotFound(ctx *gin.Context, err error) {
	var errStr string
	if err != nil {
		_ = ctx.Error(err)
		errStr = err.Error()
	}

	ctx.JSON(http.StatusNotFound, ErrorResponse{Error: errStr})
}

func (h *BasicHandler) sendInternalServerError(ctx *gin.Context, err error) {
	_ = ctx.Error(err)

	ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
}

func (h *BasicHandler) notFound(ctx *gin.Context, val any) {
	ctx.JSON(http.StatusNotFound, val)
}

func (h *BasicHandler) sendOk(ctx *gin.Context, val any) {
	ctx.JSON(http.StatusOK, val)
}
