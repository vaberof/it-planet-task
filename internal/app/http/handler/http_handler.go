package handler

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	accountService    AccountService
	locationService   LocationService
	animalTypeService AnimalTypeService
	animalService     AnimalService
	authService       AuthService
}

func NewHttpHandler(
	accountService AccountService,
	locationService LocationService,
	animalTypeService AnimalTypeService,
	animalService AnimalService,
	authService AuthService) *HttpHandler {

	return &HttpHandler{
		accountService:    accountService,
		locationService:   locationService,
		animalTypeService: animalTypeService,
		animalService:     animalService,
		authService:       authService,
	}
}

func (h *HttpHandler) InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/registration", h.Register)

	router.GET("/accounts/:accountId", h.GetAccount)
	router.PUT("/accounts/:accountId", h.UpdateAccount)
	router.GET("/accounts/search", h.SearchAccounts)
	router.DELETE("/accounts/:accountId", h.DeleteAccount)

	router.POST("/locations", h.CreateLocation)
	router.GET("/locations/:pointId", h.GetLocation)
	router.PUT("/locations/:pointId", h.UpdateLocation)
	router.DELETE("/locations/:pointId", h.DeleteLocation)

	router.POST("/animals/types", h.CreateAnimalType)
	router.GET("/animals/types/:typeId", h.GetAnimalType)
	router.PUT("/animals/types/:typeId", h.UpdateAnimalType)
	router.DELETE("/animals/types/:typeId", h.DeleteAnimalType)

	router.POST("/animals", h.CreateAnimal)
	router.GET("/animals/:animalId", h.GetAnimal)
	router.GET("/animals/search", h.SearchAnimals)
	router.PUT("/animals/:animalId", h.UpdateAnimal)
	router.DELETE("/animals/:animalId", h.DeleteAnimal)

	router.POST("/animals/:animalId/types/:typeId", h.AddAnimalsType)
	router.PUT("/animals/:animalId/types", h.UpdateAnimalsType)
	router.DELETE("/animals/:animalId/types/:typeId", h.DeleteAnimalsType)

	router.GET("/animals/:animalId/locations", h.GetAnimalsVisitedLocations)
	router.POST("/animals/:animalId/locations/:pointId", h.AddAnimalsVisitedLocation)
	router.PUT("/animals/:animalId/locations", h.UpdateAnimalsVisitedLocation)
	router.DELETE("/animals/:animalId/locations/:visitedPointId", h.DeleteAnimalsVisitedLocations)

	return router
}
