package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wwchacalww/go-cem304/domain/repository"
	"wwchacalww/go-cem304/domain/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type ParentHandler struct {
	Repo repository.ParentRepositoryInterface
}

func MakeParentHandler(r *chi.Mux, repo repository.ParentRepositoryInterface) {
	handler := &ParentHandler{
		Repo: repo,
	}
	jwtoken := jwtauth.New("HS256", []byte("secret_jwt"), nil)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtoken))
		r.Use(jwtauth.Authenticator)
		r.Route("/parents", func(r chi.Router) {
			r.Get("/{id}", handler.FindById)
		})
	})

}

func (p *ParentHandler) FindById(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin", "director", "secretary", "coordinator", "teather"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	id := chi.URLParam(r, "id")
	student, err := p.Repo.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
