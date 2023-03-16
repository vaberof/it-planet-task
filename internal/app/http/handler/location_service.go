package handler

import "github.com/vaberof/it-planet-task/internal/domain/location"

type LocationService interface {
	CreateLocation(chipperId int32, latitude, longitude float64) (*location.Location, error)
	GetLocation(pointId int64) (*location.Location, error)
	UpdateLocation(pointId int64, latitude, longitude float64) (*location.Location, error)
	DeleteLocation(pointId int64) error
}
