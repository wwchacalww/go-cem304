package handler

import (
	"encoding/json"
	"net/http"
	"wwchacalww/go-cem304/domain/repository"

	"github.com/go-chi/chi/v5"
)

type SubjectHandler struct {
	Repo repository.SubjectRepositoryInterface
}

func MakeSubjectHandlers(r *chi.Mux, repo repository.SubjectRepositoryInterface) {
	handler := &SubjectHandler{
		Repo: repo,
	}

	r.Route("/subjects", func(r chi.Router) {
		r.Post("/", handler.Store)
		r.Get("/{id}", handler.GetSubject)
		r.Get("/search", handler.FindByName)
		r.Get("/license", handler.FindByLicense)
	})
}

func (t *SubjectHandler) Store(w http.ResponseWriter, r *http.Request) {
	var input repository.SubjectInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	subject, err := t.Repo.Create(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(subject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (t *SubjectHandler) GetSubject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	subject, err := t.Repo.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(subject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (t *SubjectHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	subjects, err := t.Repo.FindByName(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(subjects)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (t *SubjectHandler) FindByLicense(w http.ResponseWriter, r *http.Request) {
	license := r.URL.Query().Get("license")
	subjects, err := t.Repo.FindByLicense(license)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(subjects)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
