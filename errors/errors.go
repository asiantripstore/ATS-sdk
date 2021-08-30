package errors

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	//ErrValidationFail error when token validation or authroziation fail
	ErrValidationFail = errors.New("validation fail")
)

//ErrHandler  to return an error to the handler with good status
func ErrHandler(c *gin.Context, err error) *gin.Context {
	msg := gin.H{"message": err.Error()}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, msg)
	} else if errors.Is(err, ErrValidationFail) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, msg)
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
	}
	return c
}
