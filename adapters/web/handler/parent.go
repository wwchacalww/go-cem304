package handler

import (
	"encoding/json"
	"net/http"
	"wwchacalww/go-cem304/domain/repository"

	"github.com/go-chi/chi/v5"
)

type ParentHandler struct {
	Repo repository.ParentRepositoryInterface
}

func MakeParentHandler(r *chi.Mux, repo repository.ParentRepositoryInterface) {
	handler := &ParentHandler{
		Repo: repo,
	}

	r.Route("/parents", func(r chi.Router) {
		r.Get("/{id}", handler.FindById)
	})
}

func (p *ParentHandler) FindById(w http.ResponseWriter, r *http.Request) {
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
