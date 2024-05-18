package static

import (
	"gohx/assets"
	ExceptionView "gohx/domain/exception/view"
	"gohx/internal/helper"
	"io"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) PublicFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filePath := r.URL.Path[len("/assets/"):]

	file, err := assets.PublicFiles.Open("public/" + filePath)
	if err != nil {
		helper.RenderView(ctx, w, ExceptionView.NotFound())
		return
	}
	defer file.Close()

	ff, ok := file.(io.ReadSeeker)
	if !ok {
		helper.RenderView(ctx, w, ExceptionView.NotFound())
		return
	}

	fileInfo, _ := file.Stat()
	http.ServeContent(w, r, filePath, fileInfo.ModTime(), ff)
}
