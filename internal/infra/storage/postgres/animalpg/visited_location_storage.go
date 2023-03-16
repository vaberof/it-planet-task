package animalpg

import (
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/vstlocationpg"
)

type VisitedLocationStorage interface {
	GetVisitedLocationById(visitedLocationId int64) (*vstlocationpg.VisitedLocation, error)
}
