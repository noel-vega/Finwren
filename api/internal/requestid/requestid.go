package requestid

import "context"

type ctxKey struct{}

var requestIDKey = ctxKey{}

func Set(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func FromContext(ctx context.Context) string {
	if v, ok := ctx.Value(requestIDKey).(string); ok {
		return v
	}

	return ""
}
