package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"github.com/vaberof/it-planet-task/internal/pkg/xvalidator"
	"strconv"
)

type updateAnimalTypeRequestBody struct {
	Type string `json:"type" binding:"required"`
}

func (h *HttpHandler) UpdateAnimalType(c *gin.Context) {
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
	
	var updateAnimalTypeReqBody updateAnimalTypeRequestBody

	if err := c.Bind(&updateAnimalTypeReqBody); err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	err = xvalidator.ConsistsOfSpaces([]string{updateAnimalTypeReqBody.Type})

	if err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	animalType, err := h.animalTypeService.UpdateAnimalType(typeId, updateAnimalTypeReqBody.Type)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderAnimalTypeResponse(c, 200, animalType)
}
