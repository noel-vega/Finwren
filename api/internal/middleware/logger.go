package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()

		level := slog.LevelInfo
		status := ctx.Writer.Status()

		switch {
		case status >= 500:
			level = slog.LevelError
		case status >= 400:
			level = slog.LevelWarn
		}

		attrs := []slog.Attr{
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.Request.URL.Path),
			slog.Int("status", status),
			slog.Duration("latency", time.Since(start)),
		}

		if len(ctx.Errors) > 0 {
			attrs = append(attrs, slog.String("errors", ctx.Errors.String()))
		}

		slog.LogAttrs(ctx.Request.Context(), level, "request", attrs...)
	}
}
