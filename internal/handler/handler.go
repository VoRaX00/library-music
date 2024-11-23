package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "library-music/docs"
	"library-music/internal/services"
	"log/slog"
)

type Handler struct {
	log     *slog.Logger
	service *services.Service
}

func NewHandler(log *slog.Logger, service *services.Service) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		api.POST("/add", h.AddMusic)
		api.PUT("/update", h.UpdateMusic)
		api.DELETE("/delete", h.DeleteMusic)
		api.GET("/info", h.GetMusic)
		api.GET("/getAll", h.GetMusicList)
		api.GET("/getText", h.GetTextMusic)
	}
	return router
}
