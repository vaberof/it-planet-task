package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"strconv"
)

func (h *HttpHandler) DeleteAccount(c *gin.Context) {
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

	err = h.accountService.DeleteAccount(account.Id, int32(accountId))
	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderResponse(c, 200, "successfully deleted account")
}
