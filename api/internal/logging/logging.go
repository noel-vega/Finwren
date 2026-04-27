package logging

import (
	"context"
	"io"
	"log/slog"
	"time"

	"github.com/lmittmann/tint"
	"github.com/noel-vega/finances/api/internal/config"
	"github.com/noel-vega/finances/api/internal/requestid"
)

type contextHandler struct {
	slog.Handler
}

func New(w io.Writer, env config.Environment) *slog.Logger {
	base := tint.NewHandler(w, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	})

	if env == config.EnvProduction {
		base = slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	return slog.New(&contextHandler{Handler: base})
}

func (h *contextHandler) Handle(ctx context.Context, r slog.Record) error {
	if requestID := requestid.FromContext(ctx); requestID != "" {
		r.AddAttrs(slog.String("request_id", requestID))
	}
	return h.Handler.Handle(ctx, r)
}
