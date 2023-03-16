package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"github.com/vaberof/it-planet-task/internal/pkg/xvalidator"
	"strconv"
)

type updateAccountRequestBody struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"  binding:"required"`
	Email     string `json:"email"  binding:"required,email"`
	Password  string `json:"password"  binding:"required"`
}

func (h *HttpHandler) UpdateAccount(c *gin.Context) {
	account, err := h.authService.AuthenticateAccount(c.Request)

	var errorWrapper *xerror.ErrorWrapper

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	accountId, err := strconv.ParseInt(c.Param("accountId"), 10, 32)
	if err != nil || accountId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid 'accountId' parameter")
		return
	}

	var updateAccountReqBody updateAccountRequestBody

	if err := c.Bind(&updateAccountReqBody); err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	err = xvalidator.ConsistsOfSpaces([]string{
		updateAccountReqBody.FirstName,
		updateAccountReqBody.LastName,
		updateAccountReqBody.Email,
		updateAccountReqBody.Password})

	if err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	account, err = h.accountService.UpdateAccount(
		account.Id,
		account.Email,
		int32(accountId),
		updateAccountReqBody.FirstName,
		updateAccountReqBody.LastName,
		updateAccountReqBody.Email,
		updateAccountReqBody.Password)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderAccountResponse(c, 200, account)
}
