package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
)

type createLocationRequestBody struct {
	Latitude  *float64 `json:"latitude" binding:"required"`
	Longitude *float64 `json:"longitude" binding:"required"`
}

func (h *HttpHandler) CreateLocation(c *gin.Context) {
	account, err := h.authService.AuthenticateAccount(c.Request)

	var errorWrapper *xerror.ErrorWrapper

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	var createLocationReqBody createLocationRequestBody

	if err := c.Bind(&createLocationReqBody); err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	latitude := *createLocationReqBody.Latitude
	if latitude < -90 || latitude > 90 {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	longitude := *createLocationReqBody.Longitude
	if longitude < -180 || longitude > 180 {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	location, err := h.locationService.CreateLocation(account.Id, latitude, longitude)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderLocationResponse(c, 201, location)
}
