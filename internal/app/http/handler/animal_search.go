package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/app/http/view"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"strconv"
	"time"
)

func (h *HttpHandler) SearchAnimals(c *gin.Context) {
	_, err := h.authService.AuthenticateAccount(c.Request)

	var errorWrapper *xerror.ErrorWrapper

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil && errorWrapper.Err.Error() != "unauthorized" {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	queryParams := c.Request.URL.Query()

	from, err := strconv.ParseInt(queryParams.Get("from"), 10, 32)
	if err != nil {
		from = 0
	}

	if from < 0 {
		view.RenderErrorResponse(c, 400, "invalid 'from' parameter")
		return
	}

	size, err := strconv.ParseInt(queryParams.Get("size"), 10, 32)
	if err != nil {
		size = 10
	}

	if size <= 0 {
		view.RenderErrorResponse(c, 400, "invalid 'size' parameter")
		return
	}

	startDateTimeParam := queryParams.Get("startDateTime")
	var startDateTime *time.Time

	if startDateTimeParam != "" {
		parsedStartDateTime, err := time.Parse("2006-01-02T15:04:05Z", startDateTimeParam)
		if err != nil {
			view.RenderErrorResponse(c, 400, "invalid 'startDateTime' parameter")
			return
		}
		startDateTime = &parsedStartDateTime
	}

	endDateTimeParam := queryParams.Get("endDateTime")
	var endDateTime *time.Time

	if endDateTimeParam != "" {
		parsedEndDateTime, err := time.Parse("2006-01-02T15:04:05Z", endDateTimeParam)
		if err != nil {
			view.RenderErrorResponse(c, 400, "invalid 'endDateTime' parameter")
			return
		}
		endDateTime = &parsedEndDateTime
	}

	chipperId, err := strconv.ParseInt(queryParams.Get("chipperId"), 10, 32)
	if err != nil {
		chipperId = -1
	} else if err == nil && chipperId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid 'chipperId' parameter")
		return
	}

	chippingLocationId, err := strconv.ParseInt(queryParams.Get("chippingLocationId"), 10, 64)
	if err != nil {
		chippingLocationId = -1
	} else if err == nil && chippingLocationId <= 0 {
		view.RenderErrorResponse(c, 400, "invalid 'chippingLocationId' parameter")
		return
	}

	animals, err := h.animalService.SearchAnimals(
		startDateTime,
		endDateTime,
		int32(chipperId),
		chippingLocationId,
		queryParams.Get("lifeStatus"),
		queryParams.Get("gender"),
		int32(from),
		int32(size))

	if errors.As(err, &errorWrapper) && errorWrapper.Err != nil {
		view.RenderErrorResponse(c, errorWrapper.StatusCode, errorWrapper.Message)
		return
	}

	view.RenderAnimalsResponse(c, 200, animals)
}
