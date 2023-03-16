package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"strconv"
)

func (h *HttpHandler) SearchAccounts(c *gin.Context) {
	_, err := h.authService.AuthenticateAccount(c.Request)

	var errorWrapper *xerror.ErrorWrapper

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil && errorWrapper.Err.Error() != "unauthorized" {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	queryParams := c.Request.URL.Query()

	from, err := strconv.Atoi(queryParams.Get("from"))
	if err != nil {
		from = 0
	}

	if from < 0 {
		view.RenderErrorResponse(c, 400, "invalid 'from' parameter")
		return
	}

	size, err := strconv.Atoi(queryParams.Get("size"))
	if err != nil {
		size = 10
	}

	if size <= 0 {
		view.RenderErrorResponse(c, 400, "invalid 'size' parameter")
		return
	}

	accounts, err := h.accountService.SearchAccounts(
		queryParams.Get("firstName"),
		queryParams.Get("lastName"),
		queryParams.Get("email"),
		from,
		size)

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderAccountsResponse(c, 200, accounts)
}
