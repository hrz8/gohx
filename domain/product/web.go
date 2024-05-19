package product

import (
	"context"
	"gohx/domain/product/view"
	"gohx/internal/helper"
	"net/http"
)

type Web struct {
	svc ProductSvc
}

type ProductSvc interface {
	GetProducts(ctx context.Context)
}

func NewWeb(svc ProductSvc) *Web {
	return &Web{svc: svc}
}

func (web *Web) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	web.svc.GetProducts(ctx)
	helper.RenderView(ctx, w, view.Index())
}
