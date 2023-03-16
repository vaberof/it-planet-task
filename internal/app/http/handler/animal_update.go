package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"strconv"
)

type updateAnimalRequestBody struct {
	Weight             float32 `json:"weight" binding:"required"`
	Length             float32 `json:"length" binding:"required"`
	Height             float32 `json:"height" binding:"required"`
	Gender             string  `json:"gender" binding:"required"`
	LifeStatus         string  `json:"lifeStatus" binding:"required"`
	ChipperId          int32   `json:"chipperId" binding:"required"`
	ChippingLocationId int64   `json:"chippingLocationId" binding:"required"`
}

func (h *HttpHandler) UpdateAnimal(c *gin.Context) {
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

	var updateAnimalReqBody updateAnimalRequestBody

	if err := c.Bind(&updateAnimalReqBody); err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	if updateAnimalReqBody.Weight <= 0 ||
		updateAnimalReqBody.Length <= 0 ||
		updateAnimalReqBody.Height <= 0 ||
		updateAnimalReqBody.ChipperId <= 0 ||
		updateAnimalReqBody.ChippingLocationId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	animal, err := h.animalService.UpdateAnimal(
		animalId,
		updateAnimalReqBody.Weight,
		updateAnimalReqBody.Length,
		updateAnimalReqBody.Height,
		updateAnimalReqBody.Gender,
		updateAnimalReqBody.LifeStatus,
		updateAnimalReqBody.ChipperId,
		updateAnimalReqBody.ChippingLocationId)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderAnimalResponse(c, 200, animal)
}
