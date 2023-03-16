package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"strconv"
)

func (h *HttpHandler) GetAnimal(c *gin.Context) {
	_, err := h.authService.AuthenticateAccount(c.Request)

	var errorWrapper *xerror.ErrorWrapper

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil && errorWrapper.Err.Error() != "unauthorized" {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	animalId, err := strconv.ParseInt(c.Param("animalId"), 10, 64)
	if err != nil || animalId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid 'animalId' parameter")
		return
	}

	animal, err := h.animalService.GetAnimalById(animalId)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderAnimalResponse(c, 200, animal)
}
