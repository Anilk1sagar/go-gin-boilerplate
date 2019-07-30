package utils

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// ValidateRequestBody for valdation
func ValidateRequestBody(c *gin.Context, fileds []string) error {

	fmt.Println("Fields : ", fileds)

	for i := 0; i < len(fileds); i++ {

		if c.PostForm(fileds[i]) == "" {
			return errors.New(fileds[i] + " is required")
		}

		if !govalidator.IsASCII(c.PostForm(fileds[i])) {
			return errors.New(fileds[i] + " should be in ASCII formate")
		}

	}

	return nil
}
