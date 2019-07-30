package utils

import (
	"github.com/gin-gonic/gin"
)

// RespondJSON makes the response with payload as json format
func RespondJSON(c *gin.Context, status int, payload interface{}) {

	c.JSON(status, payload)
}

// RespondError makes the error response with payload as json format
func RespondError(c *gin.Context, code int, message string) {

	RespondJSON(c, code, gin.H{"status": code, "error": message})
}
