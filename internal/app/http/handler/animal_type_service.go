package handler

import (
	"github.com/vaberof/it-planet-task/internal/domain/animaltype"
)

type AnimalTypeService interface {
	CreateAnimalType(chipperId int32, animalType string) (*animaltype.AnimalType, error)
	GetAnimalTypeById(typeId int64) (*animaltype.AnimalType, error)
	UpdateAnimalType(typeId int64, animalType string) (*animaltype.AnimalType, error)
	DeleteAnimalType(typeId int64) error
}
