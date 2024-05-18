package main

import (
	"fmt"
	ExceptionView "gohx/domain/exception/view"
	"gohx/domain/landing"
	"gohx/domain/product"
	"gohx/domain/static"
	"gohx/internal/helper"
	"gohx/internal/middleware"
	"gohx/internal/router"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	router := router.NewMuxRouter(mux)

	// load services
	landingSvc := landing.NewService()
	productSvc := product.NewService()

	// load http handlers
	landingHandler := landing.NewHandler(landingSvc)
	staticHandler := static.NewHandler()
	productWeb := product.NewWeb(productSvc)

	// load web ui
	router.HandleFunc("GET /{$}", landingHandler.Home)
	router.HandleFunc("GET /about", landingHandler.About).WithTrailingSlash()
	router.HandleFunc("GET /products", productWeb.Index).WithTrailingSlash()
	router.HandleFunc("GET /contact", landingHandler.Contact).WithTrailingSlash()

	// load assets
	router.HandleFunc("GET /assets/*", staticHandler.PublicFile)

	// not found handler
	router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		helper.RenderView(r.Context(), w, ExceptionView.NotFound())
	})

	server := http.Server{
		Addr:    ":4001",
		Handler: middleware.CheckHtmx(router.Mux),
	}

	fmt.Println("Server listening on port :8080")
	server.ListenAndServe()
}
