package handlers

import (
	"github.com/Jereyji/avito/internal/application/services"
	"github.com/gin-gonic/gin"
)

type EstateHandler struct {
	service *services.EstateService
}

func NewEstateHandler(serv *services.EstateService) *EstateHandler {
	return &EstateHandler{
		service: serv,
	}
}

func (h *EstateHandler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.service.Register)
		auth.POST("/login", h.service.Login)
		auth.POST("/dummyLogin", h.service.DummyLogin)
	}

	house := router.Group("/house")
	{
		house.POST("/create", h.service.CreateHouse)
		house.GET("/:id", h.service.GetFlatsByHouseId)
		//house.POST("/:id/subscribe", h.service.HouseSubscribe)
	}

	flat := router.Group("/flat")
	{
		flat.POST("/create", h.service.CreateFlat)
		flat.POST("/update", h.service.UpdateFlat)
	}

	return router
}
