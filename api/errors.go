package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProblemDetailError struct {
	Detail  string `json:"detail"`
	Pointer string `json:"pointer"`
}

type ProblemDetail struct {
	Type   string          `json:"type"`
	Status int             `json:"status"`
	Title  string          `json:"title"`
	Detail string          `json:"detail"`
	Errors []ProblemDetail `json:"Errors"`
}

type APIError struct {
	Code       string         // stable, e.g. "USER_NOT_FOUND"
	Message    string         // human-readable
	HTTPStatus int            // 404, 422, etc.
	Details    map[string]any // optional
	Cause      error          // wrapped, never serialized
}

func NewAPIError(c *gin.Context, code string, httpCode int, message string) {
	c.AbortWithStatusJSON(httpCode, APIError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpCode,
	})
	return
}

func NewNotFoundError(c *gin.Context, code string) {
	NewAPIError(c, code, http.StatusNotFound)
}
