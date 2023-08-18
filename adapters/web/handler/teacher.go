package handler

import (
	"encoding/json"
	"net/http"
	"wwchacalww/go-cem304/domain/repository"
	"wwchacalww/go-cem304/domain/utils"

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
		r.Post("/attach/class/subject", handler.AttachClassSubject)
		r.Post("/import", handler.ImportTeachers)
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

func (t *TeacherHandler) AttachClassSubject(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Id           string `json:"teacher_id"`
		Classroom_Id string `json:"classroom_id"`
		Subject_Id   string `json:"subject_id"`
		Slug         string `json:"slug"`
		Wch          int32  `json:"wch"`
		StartCourse  string `json:"start_course"`
		EndCourse    string `json:"end_course"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	teacher, err := t.Repo.AttachClassroomSubject(input.Id, input.Classroom_Id, input.Subject_Id, input.Slug, input.StartCourse, input.EndCourse, input.Wch)
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
	w.WriteHeader(http.StatusAccepted)
}

func (t *TeacherHandler) ImportTeachers(w http.ResponseWriter, r *http.Request) {
	f, fh, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	ct := fh.Header.Get("Content-Type")
	if ct != "text/csv" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError("File type invalid"))
		return
	}
	defer f.Close()
	classrooms, err := t.Repo.Classrooms("2023")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	result, err := utils.CsvToTeachers(f, classrooms)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	_ = t.Repo.Import(result)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(201)
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
