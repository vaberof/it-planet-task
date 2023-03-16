package view

import (
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/domain/animaltype"
)

type AnimalTypeResponse struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

func RenderAnimalTypeResponse(c *gin.Context, statusCode int, animalType *animaltype.AnimalType) {
	RenderResponse(c, statusCode, buildAnimalType(animalType))
}

func buildAnimalType(animalType *animaltype.AnimalType) *AnimalTypeResponse {
	return &AnimalTypeResponse{
		Id:   animalType.Id,
		Type: animalType.Type,
	}
}
