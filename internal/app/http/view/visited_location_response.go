package view

import (
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/domain/vstlocation"
)

type VisitedLocationResponse struct {
	Id                           int64  `json:"id"`
	DateTimeOfVisitLocationPoint string `json:"dateTimeOfVisitLocationPoint"`
	LocationPointId              int64  `json:"locationPointId"`
}

func RenderVisitedLocationResponse(c *gin.Context, statusCode int, visitedLocation *vstlocation.VisitedLocation) {
	RenderResponse(c, statusCode, buildVisitedLocation(visitedLocation))
}

func RenderVisitedLocationsResponse(c *gin.Context, statusCode int, visitedLocations []*vstlocation.VisitedLocation) {
	visitedLocationsResponse := make([]*VisitedLocationResponse, len(visitedLocations))

	for i := 0; i < len(visitedLocationsResponse); i++ {
		visitedLocationsResponse[i] = buildVisitedLocation(visitedLocations[i])
	}

	RenderResponse(c, statusCode, visitedLocationsResponse)
}

func buildVisitedLocation(visitedLocation *vstlocation.VisitedLocation) *VisitedLocationResponse {
	return &VisitedLocationResponse{
		Id:                           visitedLocation.Id,
		DateTimeOfVisitLocationPoint: visitedLocation.DateTimeOfVisitLocationPoint.Format("2006-01-02T15:04:05-07:00"),
		LocationPointId:              visitedLocation.LocationPointId,
	}
}
