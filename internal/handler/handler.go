package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "library-music/docs"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
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
		api.PATCH("/update", h.UpdatePartialMusic)
		api.DELETE("/delete", h.DeleteMusic)
		api.GET("/getMusic", h.GetMusic)
		api.GET("/getAll:page", h.GetAllMusic)
		api.GET("/getText:page", h.GetTextMusic)
	}

	return router
}
