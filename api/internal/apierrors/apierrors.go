package apierrors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProblemDetailError struct {
	Code    string            `json:"code"`
	Detail  string            `json:"detail"`
	Pointer string            `json:"pointer"`
	Params  map[string]string `json:"params,omitempty"`
}

type ProblemDetail struct {
	Type   string               `json:"type"`
	Status int                  `json:"status"`
	Title  string               `json:"title"`
	Detail string               `json:"detail"`
	Errors []ProblemDetailError `json:"errors"`
}

func NewBadRequest(ctx *gin.Context, detail string, problems []ProblemDetailError) {
	problem := ProblemDetail{
		Type:   "about:blank",
		Status: http.StatusBadRequest,
		Title:  http.StatusText(http.StatusBadRequest),
		Detail: detail,
		Errors: problems,
	}
	ctx.AbortWithStatusJSON(problem.Status, problem)
}

func NewConflict(ctx *gin.Context, detail string, problems []ProblemDetailError) {
	problem := ProblemDetail{
		Type:   "about:blank",
		Status: http.StatusConflict,
		Title:  http.StatusText(http.StatusConflict),
		Detail: detail,
		Errors: problems,
	}
	ctx.AbortWithStatusJSON(problem.Status, problem)
}

func NewInternalServerError(ctx *gin.Context) {
	problem := ProblemDetail{
		Type:   "about:blank",
		Status: http.StatusInternalServerError,
		Title:  http.StatusText(http.StatusInternalServerError),
		Detail: "Unexpected internal server error",
	}
	ctx.AbortWithStatusJSON(problem.Status, problem)
}
