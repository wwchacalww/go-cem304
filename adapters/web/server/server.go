package server

import (
	"net/http"
	"wwchacalww/go-cem304/adapters/web/handler"
	"wwchacalww/go-cem304/application"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	UsersService application.UserServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Server() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	handler.MakeUserHandlers(r, w.UsersService)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":9000", r)
}
