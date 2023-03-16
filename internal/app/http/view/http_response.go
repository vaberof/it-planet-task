package view

import "github.com/gin-gonic/gin"

type HttpResponse struct {
	Data     interface{} `json:"data,omitempty"`
	ErrorMsg string      `json:"error,omitempty"`
}

func RenderResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

func RenderErrorResponse(c *gin.Context, statusCode int, errorMsg string) {
	c.JSON(statusCode, HttpResponse{ErrorMsg: errorMsg})
}
