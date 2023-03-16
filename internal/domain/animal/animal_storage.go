package animal

import (
	"github.com/vaberof/it-planet-task/internal/domain/vstlocation"
	"time"
)

type AnimalStorage interface {
	CreateAnimal(
		chipperId int32,
		chippingLocationId int64,
		animalTypes []int64,
		weight float32,
		length float32,
		height float32,
		gender string) (*Animal, error)
	GetAnimalById(id int64) (*Animal, error)
	SearchAnimals(
		startDateTime *time.Time,
		endDateTime *time.Time,
		chipperId int32,
		chippingLocationId int64,
		lifeStatus string,
		gender string,
		from int32,
		size int32) ([]*Animal, error)
	UpdateAnimal(
		animalId int64,
		weight float32,
		length float32,
		height float32,
		gender string,
		lifeStatus string,
		chipperId int32,
		chippingLocationId int64) (*Animal, error)
	FindAnimalById(typeId int64) error
	DeleteAnimal(animalId int64) error

	AddAnimalsType(animalId, typeId int64) (*Animal, error)
	UpdateAnimalsType(animalId, oldTypeId, newTypeId int64) (*Animal, error)
	DeleteAnimalsType(animalId, typeId int64) (*Animal, error)

	GetAnimalsVisitedLocations(
		animalId int64,
		startDateTime *time.Time,
		endDateTime *time.Time,
		from int32,
		size int32) ([]*vstlocation.VisitedLocation, error)
	UpdateAnimalsVisitedLocation(visitedLocationId, locationId int64) (*vstlocation.VisitedLocation, error)
	DeleteAnimalsVisitedLocation(animalId int64, index int) error
}
