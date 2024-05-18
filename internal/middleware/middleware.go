package middleware

import (
	"context"
	"net/http"
)

func setRequestContext(ctx context.Context, r *http.Request, key, val any) {
	req := r.WithContext(context.WithValue(ctx, key, val))
	*r = *req
}
