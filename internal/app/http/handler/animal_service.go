package handler

import (
	"github.com/vaberof/it-planet-task/internal/domain/animal"
	"github.com/vaberof/it-planet-task/internal/domain/vstlocation"
	"time"
)

type AnimalService interface {
	CreateAnimal(
		chipperId int32,
		chippingLocationId int64,
		animalTypes []int64,
		weight float32,
		length float32,
		height float32,
		gender string) (*animal.Animal, error)
	GetAnimalById(id int64) (*animal.Animal, error)
	SearchAnimals(
		startDateTime,
		endDateTime *time.Time,
		chipperId int32,
		chippingLocationId int64,
		lifeStatus,
		gender string,
		from,
		size int32) ([]*animal.Animal, error)
	UpdateAnimal(
		animalId int64,
		weight float32,
		length float32,
		height float32,
		gender,
		lifeStatus string,
		chipperId int32,
		chippingLocationId int64) (*animal.Animal, error)
	DeleteAnimal(animalId int64) error

	AddAnimalsType(animalId int64, animalTypeId int64) (*animal.Animal, error)
	UpdateAnimalsType(animalId int64, oldTypeId, newTypeId int64) (*animal.Animal, error)
	DeleteAnimalsType(animalId, animalTypeId int64) (*animal.Animal, error)

	AddAnimalsVisitedLocation(animalId, locationId int64, chipperId int32) (*vstlocation.VisitedLocation, error)
	GetAnimalsVisitedLocations(
		animalId int64,
		startDateTime,
		endDateTime *time.Time,
		from,
		size int32) ([]*vstlocation.VisitedLocation, error)
	UpdateAnimalsVisitedLocation(animalId, visitedLocationId, locationId int64) (*vstlocation.VisitedLocation, error)
	DeleteAnimalsVisitedLocation(animalId, visitedLocationId int64) error
}
