package apiHelpers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func Descriptive(verr validator.ValidationErrors) map[string][]string {
	var errs = make(map[string][]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		if _, ok := errs[f.Field()]; !ok {
			errs[strings.ToLower(f.Field())] = []string{err}
		} else {
			_ = append(errs[f.Field()], err)
		}
	}

	return errs
}

func HandleError(err error, c *gin.Context) {
	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		c.JSON(http.StatusBadRequest, Descriptive(verr))
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	return
}
