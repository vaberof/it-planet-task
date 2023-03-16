package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"strconv"
)

type updateAnimalsTypeRequestBody struct {
	OldTypeId int64 `json:"oldTypeId" binding:"required"`
	NewTypeId int64 `json:"newTypeId" binding:"required"`
}

func (h *HttpHandler) UpdateAnimalsType(c *gin.Context) {
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

	var updateAnimalsTypeReqBody updateAnimalsTypeRequestBody

	if err := c.Bind(&updateAnimalsTypeReqBody); err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	if updateAnimalsTypeReqBody.OldTypeId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	if updateAnimalsTypeReqBody.NewTypeId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	animal, err := h.animalService.UpdateAnimalsType(
		animalId,
		updateAnimalsTypeReqBody.OldTypeId,
		updateAnimalsTypeReqBody.NewTypeId)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}
	
	view.RenderAnimalResponse(c, 200, animal)
}
