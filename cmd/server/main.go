package main

import (
	"context"
	"fmt"
	ExceptionView "gohx/domain/exception/view"
	"gohx/domain/landing"
	"gohx/domain/product"
	"gohx/domain/static"
	"gohx/domain/todolist"
	"gohx/internal/helper"
	"gohx/internal/middleware"
	dbrepo "gohx/internal/repo/db"
	"gohx/internal/router"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	ctx := context.Background()
	mux := http.NewServeMux()
	router := router.NewMuxRouter(mux)

	db := connect(ctx)
	repo := dbrepo.New(db)

	// load services
	landingSvc := landing.NewService()
	productSvc := product.NewService()
	todolistSvc := todolist.NewService(repo)

	// load http web ui
	landingWeb := landing.NewWeb(landingSvc, todolistSvc)
	productWeb := product.NewWeb(productSvc)

	// load web ui
	router.HandleFunc("GET /{$}", landingWeb.Home)
	router.HandleFunc("GET /about", landingWeb.About).WithTrailingSlash()
	router.HandleFunc("GET /products", productWeb.Index).WithTrailingSlash()
	router.HandleFunc("GET /contact", landingWeb.Contact).WithTrailingSlash()

	// load assets
	staticHandler := static.NewHandler()
	router.HandleFunc("GET /assets/*", staticHandler.PublicFile)

	// not found handler
	router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		helper.RenderView(r.Context(), w, ExceptionView.NotFound())
	})

	server := http.Server{
		Addr:    ":4001",
		Handler: middleware.CheckHtmx(router.Mux),
	}

	fmt.Println("Server listening on port :4001")
	server.ListenAndServe()
}
