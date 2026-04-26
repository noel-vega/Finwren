package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	ridHeaderKey  = "X-Request-ID"
	ridContextKey = "request_id"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid := ctx.GetHeader(ridHeaderKey)

		if err := uuid.Validate(rid); err != nil {
			rid = uuid.NewString()
		}

		ctx.Set(ridContextKey, rid)
		ctx.Writer.Header().Set(ridHeaderKey, rid)
		ctx.Next()
	}
}
