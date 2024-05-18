package landing

import (
	"context"
	"gohx/domain/landing/view"
	"gohx/internal/helper"
	"net/http"
)

type Handler struct {
	svc LandingSvc
}

type LandingSvc interface {
	LongRunningMethod(ctx context.Context)
}

func NewHandler(svc LandingSvc) *Handler {
	return &Handler{svc}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	h.svc.LongRunningMethod(ctx)
	helper.RenderView(ctx, w, view.Home())
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	h.svc.LongRunningMethod(ctx)
	helper.RenderView(ctx, w, view.About())
}

func (h *Handler) Contact(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	h.svc.LongRunningMethod(ctx)
	helper.RenderView(ctx, w, view.Contact())
}
