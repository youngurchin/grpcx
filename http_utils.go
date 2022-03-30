package grpcx

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONResult json result
type JSONResult struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"msg,omitempty"`
}

// NewJSONBodyHTTPHandle returns a http handle JSON body
func NewJSONBodyHTTPHandle(factory func() interface{}, handler func(interface{}) (*JSONResult, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		value := factory()
		err := ReadJSONFromBody(c, value)
		if err != nil {
			c.JSON(http.StatusOK, &JSONResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		result, err := handler(value)
		if err != nil {
			c.JSON(http.StatusOK, &JSONResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, result)
		return
	}
}

// NewGetHTTPHandle return get http handle
func NewGetHTTPHandle(factory func(*gin.Context) (interface{}, error), handler func(interface{}) (*JSONResult, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		value, err := factory(c)
		if err != nil {
			c.JSON(http.StatusOK, &JSONResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		result, err := handler(value)
		if err != nil {
			c.JSON(http.StatusOK, &JSONResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, result)
		return
	}
}

// ReadJSONFromBody read json body
func ReadJSONFromBody(ctx *gin.Context, value interface{}) error {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}

	if len(data) > 0 {
		err = json.Unmarshal(data, value)
		if err != nil {
			return err
		}
	}

	return nil
}
