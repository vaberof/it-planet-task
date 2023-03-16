package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"strconv"
)

func (h *HttpHandler) DeleteAnimal(c *gin.Context) {
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

	err = h.animalService.DeleteAnimal(animalId)
	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderResponse(c, 200, "successfully deleted animal")
}
