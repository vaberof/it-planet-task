package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"github.com/vaberof/it-planet-task/internal/pkg/xvalidator"
)

type createAnimalRequestBody struct {
	AnimalTypes        []*int64 `json:"animalTypes" binding:"required"`
	Weight             float32  `json:"weight" binding:"required"`
	Length             float32  `json:"length" binding:"required"`
	Height             float32  `json:"height" binding:"required"`
	Gender             string   `json:"gender" binding:"required"`
	ChipperId          int32    `json:"chipperId" binding:"required"`
	ChippingLocationId int64    `json:"chippingLocationId" binding:"required"`
}

func (h *HttpHandler) CreateAnimal(c *gin.Context) {
	_, err := h.authService.AuthenticateAccount(c.Request)

	var errorWrapper *xerror.ErrorWrapper

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	var createAnimalReqBody createAnimalRequestBody

	if err := c.Bind(&createAnimalReqBody); err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	animalTypes, err := xvalidator.ValidateAndConvertAnimalTypes(createAnimalReqBody.AnimalTypes)
	if err != nil {
		view.RenderErrorResponse(c, 400, err.Error())
		return
	}

	if createAnimalReqBody.Weight <= 0 ||
		createAnimalReqBody.Length <= 0 ||
		createAnimalReqBody.Height <= 0 ||
		createAnimalReqBody.ChipperId <= 0 ||
		createAnimalReqBody.ChippingLocationId <= 0 {

		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	animal, err := h.animalService.CreateAnimal(
		createAnimalReqBody.ChipperId,
		createAnimalReqBody.ChippingLocationId,
		animalTypes,
		createAnimalReqBody.Weight,
		createAnimalReqBody.Length,
		createAnimalReqBody.Height,
		createAnimalReqBody.Gender)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderAnimalResponse(c, 201, animal)
}
