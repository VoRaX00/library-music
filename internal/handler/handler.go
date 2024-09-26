package handler

import (
	"github.com/gin-gonic/gin"
	"library-music/internal/application"
)

type Handler struct {
	service *application.Service
}

func NewHandler(service *application.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	return router
}
