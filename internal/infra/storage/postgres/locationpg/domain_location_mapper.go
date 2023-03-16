package locationpg

import (
	"github.com/vaberof/it-planet-task/internal/domain/location"
)

func BuildDomainLocation(postgresLocation *Location) *location.Location {
	return &location.Location{
		Id:        postgresLocation.Id,
		Latitude:  postgresLocation.Latitude,
		Longitude: postgresLocation.Longitude,
	}
}
