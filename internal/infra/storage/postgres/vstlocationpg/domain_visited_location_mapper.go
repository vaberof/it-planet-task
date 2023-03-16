package vstlocationpg

import (
	"github.com/vaberof/it-planet-task/internal/domain/vstlocation"
)

func BuildDomainVisitedLocation(postgresVisitedLocation *VisitedLocation) *vstlocation.VisitedLocation {
	return &vstlocation.VisitedLocation{
		Id:                           postgresVisitedLocation.Id,
		DateTimeOfVisitLocationPoint: postgresVisitedLocation.DateTimeOfVisitLocationPoint,
		LocationPointId:              postgresVisitedLocation.LocationPointId,
	}
}

func BuildDomainVisitedLocations(postgresVisitedLocations []*VisitedLocation) []*vstlocation.VisitedLocation {
	domainVisitedLocations := make([]*vstlocation.VisitedLocation, len(postgresVisitedLocations))

	for i := 0; i < len(postgresVisitedLocations); i++ {
		domainVisitedLocations[i] = BuildDomainVisitedLocation(postgresVisitedLocations[i])
	}

	return domainVisitedLocations
}
