package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"strconv"
)

type updateLocationRequestBody struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

func (h *HttpHandler) UpdateLocation(c *gin.Context) {
	_, err := h.authService.AuthenticateAccount(c.Request)

	var errorWrapper *xerror.ErrorWrapper

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	pointId, err := strconv.ParseInt(c.Param("pointId"), 10, 64)
	if err != nil || pointId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid 'pointId' parameter")
		return
	}

	var updateLocationReqBody updateLocationRequestBody

	if err := c.Bind(&updateLocationReqBody); err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	latitude := updateLocationReqBody.Latitude
	if latitude < -90 || latitude > 90 {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	longitude := updateLocationReqBody.Longitude
	if longitude < -180 || longitude > 180 {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	location, err := h.locationService.UpdateLocation(pointId, latitude, longitude)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderLocationResponse(c, 200, location)
}
