package handlers

import (
	"github.com/Jereyji/avito/internal/application/services"
	"github.com/Jereyji/avito/internal/presentation/DTO"
	"github.com/gin-gonic/gin"
	"log"
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
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		//auth.POST("/dummyLogin", h.DummyLogin)
	}

	//house := router.Group("/house")
	//{
	//	house.POST("/create", h.CreateHouse)
	//	house.GET("/:id", h.GetFlatsByHouseId)
	//	//house.POST("/:id/subscribe", h.HouseSubscribe)
	//}
	//
	//flat := router.Group("/flat")
	//{
	//	flat.POST("/create", h.CreateFlat)
	//	flat.POST("/update", h.UpdateFlat)
	//}

	return router
}

func (h *EstateHandler) Register(c *gin.Context) {
	var user DTO.UserDTO
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("Input error")
		return
	}

	err = h.service.Register(c, user.Username, user.Password, user.AccessLevel)

	log.Println("JHHOHOAH")

}

func (h *EstateHandler) Login(c *gin.Context) {

}
