package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"strconv"
)

func (h *HttpHandler) DeleteAnimalType(c *gin.Context) {
	_, err := h.authService.AuthenticateAccount(c.Request)

	var errorWrapper *xerror.ErrorWrapper

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	typeId, err := strconv.ParseInt(c.Param("typeId"), 10, 64)
	if err != nil || typeId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid 'typeId' parameter")
		return
	}

	err = h.animalTypeService.DeleteAnimalType(typeId)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderResponse(c, 200, "successfully deleted animal type")
}
