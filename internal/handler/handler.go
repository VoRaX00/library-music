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

	api := router.Group("/api")
	{
		api.POST("/add", h.AddMusic)
		api.PUT("/update", h.UpdateMusic)
		api.DELETE("/delete", h.DeleteMusic)
		api.GET("/get:song", h.GetMusic)
		api.GET("/getAll", h.GetMusicList)
		api.GET("/getText", h.GetTextMusic)
	}
	return router
}
