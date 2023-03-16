package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"strconv"
)

type updateAnimalsVisitedLocationRequestBody struct {
	VisitedLocationPointId int64 `json:"visitedLocationPointId" binding:"required"`
	LocationPointId        int64 `json:"locationPointId" binding:"required"`
}

func (h *HttpHandler) UpdateAnimalsVisitedLocation(c *gin.Context) {
	_, err := h.authService.AuthenticateAccount(c.Request)

	var errorWrapper *xerror.ErrorWrapper

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	animalId, err := strconv.ParseInt(c.Param("animalId"), 10, 64)
	if err != nil || animalId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid 'animalId' parameter")
		return
	}

	var updateAnimalsVisitedLocationReqBody updateAnimalsVisitedLocationRequestBody

	if err := c.Bind(&updateAnimalsVisitedLocationReqBody); err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	if updateAnimalsVisitedLocationReqBody.VisitedLocationPointId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	if updateAnimalsVisitedLocationReqBody.LocationPointId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	visitedLocation, err := h.animalService.UpdateAnimalsVisitedLocation(
		animalId,
		updateAnimalsVisitedLocationReqBody.VisitedLocationPointId,
		updateAnimalsVisitedLocationReqBody.LocationPointId)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderVisitedLocationResponse(c, 200, visitedLocation)
}
