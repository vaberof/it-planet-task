package view

import (
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/domain/location"
)

type LocationResponse struct {
	Id        int64   `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func RenderLocationResponse(c *gin.Context, statusCode int, location *location.Location) {
	RenderResponse(c, statusCode, buildLocation(location))
}

func buildLocation(location *location.Location) *LocationResponse {
	return &LocationResponse{
		Id:        location.Id,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}
}
