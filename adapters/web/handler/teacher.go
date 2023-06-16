package handler

import (
	"encoding/json"
	"net/http"
	"wwchacalww/go-cem304/domain/repository"

	"github.com/go-chi/chi/v5"
)

type TeacherHandler struct {
	Repo repository.TeacherRepositoryInterface
}

func MakeTeacherHandlers(r *chi.Mux, repo repository.TeacherRepositoryInterface) {
	handler := &TeacherHandler{
		Repo: repo,
	}

	r.Route("/teachers", func(r chi.Router) {
		r.Post("/", handler.Store)
		r.Get("/{id}", handler.GetTeacher)
		r.Get("/search", handler.FindByName)
	})
}

func (t *TeacherHandler) Store(w http.ResponseWriter, r *http.Request) {
	var input repository.TeacherInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	teacher, err := t.Repo.Create(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(teacher)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (t *TeacherHandler) GetTeacher(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	teacher, err := t.Repo.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(teacher)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (t *TeacherHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	teachers, err := t.Repo.FindByName(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(teachers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
