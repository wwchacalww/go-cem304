package server

import (
	"net/http"
	"wwchacalww/go-cem304/adapters/web/handler"
	"wwchacalww/go-cem304/domain/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

type WebServer struct {
	AuthRepository      repository.AuthRepositoryInterface
	UserRepository      repository.UserRepositoryInterface
	ClassroomRepository repository.ClassroomRepositoryInterface
	StudentRepository   repository.StudentRepositoryInterface
	ParentRepository    repository.ParentRepositoryInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Server() {
	r := chi.NewRouter()
	c := cors.New(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(c.Handler)
	r.Use(middleware.Logger)
	handler.MakeAuthHandlers(r, w.AuthRepository)
	handler.MakeUserHandlers(r, w.UserRepository)
	handler.MakeClassroomHandlers(r, w.ClassroomRepository)
	handler.MakeStudentHandlers(r, w.StudentRepository)
	handler.MakeParentHandler(r, w.ParentRepository)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":9000", r)
}
