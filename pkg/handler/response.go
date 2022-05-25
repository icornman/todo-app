package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Printf(message)
	c.AbortWithStatusJSON(statusCode, Error{message})
}
