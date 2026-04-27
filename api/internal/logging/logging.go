package logging

import (
	"context"
	"io"
	"log/slog"

	"github.com/noel-vega/finances/api/internal/requestid"
)

type contextHandler struct {
	slog.Handler
}

func New(w io.Writer) *slog.Logger {
	base := slog.NewJSONHandler(w, nil)
	return slog.New(&contextHandler{Handler: base})
}

func (h *contextHandler) Handle(ctx context.Context, r slog.Record) error {
	if requestID := requestid.FromContext(ctx); requestID != "" {
		r.AddAttrs(slog.String("request_id", requestID))
	}
	return h.Handler.Handle(ctx, r)
}
