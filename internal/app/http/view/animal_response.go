package view

import (
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/domain/animal"
	"github.com/vaberof/it-planet-task/internal/domain/vstlocation"
)

type AnimalResponse struct {
	Id                 int64   `json:"id"`
	AnimalTypes        []int64 `json:"animalTypes"`
	Weight             float32 `json:"weight"`
	Length             float32 `json:"length"`
	Height             float32 `json:"height"`
	Gender             string  `json:"gender"`
	LifeStatus         string  `json:"lifeStatus"`
	ChippingDateTime   string  `json:"chippingDateTime"`
	ChipperId          int32   `json:"chipperId"`
	ChippingLocationId int64   `json:"chippingLocationId"`
	VisitedLocations   []int64 `json:"visitedLocations"`
	DeathDateTime      *string `json:"deathDateTime"`
}

func RenderAnimalResponse(c *gin.Context, statusCode int, animal *animal.Animal) {
	RenderResponse(c, statusCode, buildAnimal(animal))
}

func RenderAnimalsResponse(c *gin.Context, statusCode int, animals []*animal.Animal) {
	animalsResponse := make([]*AnimalResponse, len(animals))

	for i := 0; i < len(animalsResponse); i++ {
		animalsResponse[i] = buildAnimal(animals[i])
	}

	RenderResponse(c, statusCode, animalsResponse)
}

func buildAnimal(animal *animal.Animal) *AnimalResponse {
	var responseDeathDateTime *string

	if animal.DeathDateTime != nil {
		convDeathDateTime := animal.DeathDateTime.Format("2006-01-02T15:04:05-07:00")
		responseDeathDateTime = &convDeathDateTime
	}

	return &AnimalResponse{
		Id:                 animal.Id,
		AnimalTypes:        animal.AnimalTypes,
		Weight:             animal.Weight,
		Length:             animal.Length,
		Height:             animal.Height,
		Gender:             animal.Gender,
		LifeStatus:         animal.LifeStatus,
		ChippingDateTime:   animal.ChippingDateTime.Format("2006-01-02T15:04:05-07:00"),
		ChipperId:          animal.ChipperId,
		ChippingLocationId: animal.ChippingLocationId,
		VisitedLocations:   convertVisitedLocations(animal.VisitedLocations),
		DeathDateTime:      responseDeathDateTime,
	}
}

func convertVisitedLocations(visitedLocations []*vstlocation.VisitedLocation) []int64 {
	visitedLocationsResponse := make([]int64, len(visitedLocations))

	for i := 0; i < len(visitedLocations); i++ {
		visitedLocationsResponse[i] = visitedLocations[i].Id
	}

	return visitedLocationsResponse
}
