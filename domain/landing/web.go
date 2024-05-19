package landing

import (
	"context"
	"gohx/domain/landing/view"
	"gohx/internal/helper"
	dbrepo "gohx/internal/repo/db"
	"net/http"
)

type Web struct {
	svc         LandingSvc
	todolistSvc TodolistSvc
}

type LandingSvc interface {
	LongRunningMethod(ctx context.Context)
}

type TodolistSvc interface {
	GetAllTasks(ctx context.Context) ([]*dbrepo.GetAllTasksRow, error)
}

func NewWeb(svc LandingSvc, todolistSvc TodolistSvc) *Web {
	return &Web{svc: svc}
}

func (web *Web) Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	web.svc.LongRunningMethod(ctx)
	helper.RenderView(ctx, w, view.Home())
}

func (web *Web) About(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	web.svc.LongRunningMethod(ctx)
	helper.RenderView(ctx, w, view.About())
}

func (web *Web) Contact(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	web.svc.LongRunningMethod(ctx)
	helper.RenderView(ctx, w, view.Contact())
}
