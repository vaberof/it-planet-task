package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"strconv"
)

func (h *HttpHandler) AddAnimalsVisitedLocation(c *gin.Context) {
	account, err := h.authService.AuthenticateAccount(c.Request)

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

	locationId, err := strconv.ParseInt(c.Param("pointId"), 10, 64)
	if err != nil || locationId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid 'pointId' parameter")
		return
	}

	visitedLocation, err := h.animalService.AddAnimalsVisitedLocation(animalId, locationId, account.Id)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderVisitedLocationResponse(c, 201, visitedLocation)
}
