package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"strconv"
)

func (h *HttpHandler) DeleteAnimalsVisitedLocations(c *gin.Context) {
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

	visitedPointId, err := strconv.ParseInt(c.Param("visitedPointId"), 10, 64)
	if err != nil || visitedPointId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid 'visitedPointId' parameter")
		return
	}

	err = h.animalService.DeleteAnimalsVisitedLocation(animalId, visitedPointId)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderResponse(c, 200, "successfully deleted visited location")
}
