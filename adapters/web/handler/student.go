package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"wwchacalww/go-cem304/domain/repository"

	"github.com/go-chi/chi/v5"
)

type StudentHandler struct {
	Repo repository.StudentRepositoryInterface
}

func MakeStudentHandlers(r *chi.Mux, repo repository.StudentRepositoryInterface) {
	handler := &StudentHandler{
		Repo: repo,
	}

	r.Route("/students", func(r chi.Router) {
		r.Post("/", handler.Store)
		// r.Post("/import", handler.Import)
		r.Get("/{id}", handler.GetStudent)
		r.Get("/search", handler.FindByName)
		r.Get("/list", handler.List)
		// r.Patch("/enable/{id}", handler.Enable)
		// r.Patch("/disable/{id}", handler.Disable)
		// r.Patch("/anne", handler.ANNE)
	})
}

func (s *StudentHandler) Store(w http.ResponseWriter, r *http.Request) {
	var input repository.StudentInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	class, err := s.Repo.Create(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(class)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	class, err := s.Repo.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(class)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *StudentHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	Students, err := s.Repo.FindByName(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(Students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *StudentHandler) List(w http.ResponseWriter, r *http.Request) {
	year := r.URL.Query().Get("year")
	log.Println(year)
	Students, err := s.Repo.List(year)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(Students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *StudentHandler) Enable(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	class, err := s.Repo.Enable(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(class)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *StudentHandler) Disable(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	class, err := s.Repo.Disable(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(class)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *StudentHandler) ANNE(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID   string `json:"id"`
		ANNE string `json:"ANNE"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	class, err := s.Repo.ANNE(input.ID, input.ANNE)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(class)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// func (s *StudentHandler) Import(w http.ResponseWriter, r *http.Request) {
// 	f, fh, err := r.FormFile("file")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write(jsonError(err.Error()))
// 		return
// 	}
// 	ct := fh.Header.Get("Content-Type")
// 	if ct != "text/csv" {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write(jsonError("File type invalid"))
// 		return
// 	}
// 	defer f.Close()

// 	list, err := utils.CsvToStudents(f)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write(jsonError(err.Error()))
// 		return
// 	}

// 	result, err := s.Repo.AddMass(list)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write(jsonError(err.Error()))
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(result)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write(jsonError(err.Error()))
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }
