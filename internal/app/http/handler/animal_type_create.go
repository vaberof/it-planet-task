package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"github.com/vaberof/it-planet-task/internal/pkg/xvalidator"
)

type createAnimalTypeRequestBody struct {
	Type string `json:"type" binding:"required"`
}

func (h *HttpHandler) CreateAnimalType(c *gin.Context) {
	account, err := h.authService.AuthenticateAccount(c.Request)

	var errorWrapper *xerror.ErrorWrapper

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	var createAnimalTypeReqBody createAnimalTypeRequestBody

	if err := c.Bind(&createAnimalTypeReqBody); err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	err = xvalidator.ConsistsOfSpaces([]string{createAnimalTypeReqBody.Type})
	if err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	animalType, err := h.animalTypeService.CreateAnimalType(account.Id, createAnimalTypeReqBody.Type)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderAnimalTypeResponse(c, 201, animalType)
}
