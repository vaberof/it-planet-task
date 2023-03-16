package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"github.com/vaberof/it-planet-task/internal/pkg/xvalidator"
)

type registerAccountRequestBody struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"  binding:"required"`
	Email     string `json:"email"  binding:"required,email"`
	Password  string `json:"password"  binding:"required"`
}

func (h *HttpHandler) Register(c *gin.Context) {
	_, err := h.authService.AuthenticateAccount(c.Request)
	if err == nil {
		view.RenderErrorResponse(c, 403, "request from authorized account")
		return
	}

	var registerAccountReqBody registerAccountRequestBody

	if err := c.Bind(&registerAccountReqBody); err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	err = xvalidator.ConsistsOfSpaces([]string{
		registerAccountReqBody.FirstName,
		registerAccountReqBody.LastName,
		registerAccountReqBody.Email,
		registerAccountReqBody.Password})

	if err != nil {
		view.RenderErrorResponse(c, 400, "invalid request body")
		return
	}

	account, err := h.accountService.CreateAccount(
		registerAccountReqBody.FirstName,
		registerAccountReqBody.LastName,
		registerAccountReqBody.Email,
		registerAccountReqBody.Password)

	var errorWrapper *xerror.ErrorWrapper

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderAccountResponse(c, 201, account)
}
