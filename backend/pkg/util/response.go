package util

import "github.com/gin-gonic/gin"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int		 			`json:"status"`
}

func NewResponse(message string, data interface{}, status int) *Response {
	return &Response{
		Message: message,
		Data:    data,
		Status:  status,
	}
}

func (r *Response) ReturnGin(c *gin.Context) {
	c.JSON(r.Status, r)
}