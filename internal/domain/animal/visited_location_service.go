package animal

import (
	"github.com/vaberof/it-planet-task/internal/domain/vstlocation"
)

type VisitedLocationService interface {
	CreateVisitedLocation(animalId, locationPointId int64, chipperId int32) (*vstlocation.VisitedLocation, error)
}
