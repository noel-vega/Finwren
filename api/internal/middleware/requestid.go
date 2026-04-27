package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/noel-vega/finances/api/internal/requestid"
)

const (
	ridHeaderKey = "X-Request-ID"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.GetHeader(ridHeaderKey)

		if err := uuid.Validate(requestID); err != nil {
			requestID = uuid.NewString()
		}

		newCtx := requestid.Set(ctx.Request.Context(), requestID)
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Writer.Header().Set(ridHeaderKey, requestID)
		ctx.Next()
	}
}
