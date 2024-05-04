package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func newJSONResponse(c *gin.Context, statusCode int, message string, err string) {
	logrus.Errorf(message, err)
	c.JSON(statusCode, gin.H{message: err})
}
