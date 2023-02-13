package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"wwchacalww/go-cem304/domain/repository"
	"wwchacalww/go-cem304/domain/utils"
	"wwchacalww/go-cem304/usecase/check"
	reportpdf "wwchacalww/go-cem304/usecase/report-pdf"

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
		r.Post("/import", handler.Import)
		r.Get("/{id}", handler.GetStudent)
		r.Get("/educar/{id}", handler.FindByEducar)
		r.Get("/search", handler.FindByName)
		r.Get("/list", handler.List)
		r.Put("/change", handler.ChangeClassroom)
		r.Put("/checkclass", handler.CheckStudentsInClass)
		r.Post("/statement/schooling", handler.StatementSchooling)
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

func (s *StudentHandler) ChangeClassroom(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID          string `json:"id"`
		ClassroomID string `json:"classroom_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = s.Repo.ChangeClassroom(input.ID, input.ClassroomID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	w.WriteHeader(201)
	w.Write(jsonError("classroom updated"))
}

func (s *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	student, err := s.Repo.FindById(id)
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

func (s *StudentHandler) FindByEducar(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Println(id)
	educar, err := strconv.ParseInt(id, 10, 64)
	log.Println(educar)
	student, err := s.Repo.FindByEducar(educar)
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
	classroom_id := r.URL.Query().Get("class")
	log.Println(classroom_id)
	Students, err := s.Repo.List(classroom_id)
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

func (s *StudentHandler) Import(w http.ResponseWriter, r *http.Request) {
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

	list, err := utils.CsvToStudents(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	result, err := s.Repo.AddMass(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *StudentHandler) CheckStudentsInClass(w http.ResponseWriter, r *http.Request) {
	f, fh, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	classroom_id := r.FormValue("classroom_id")
	students, err := s.Repo.List(classroom_id)
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

	list, err := utils.StudentsInClass(f, classroom_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	output, err := check.CheckStudentsInClass(list, classroom_id, students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	log.Println(output)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *StudentHandler) StatementSchooling(w http.ResponseWriter, r *http.Request) {
	var input reportpdf.InputSS
	err := json.NewDecoder(r.Body).Decode(&input)
	student, err := s.Repo.FindByEducar(input.Educar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	var inputState reportpdf.InputStatementSchooling
	inputState.Student = student
	inputState.Matricula = input.Matricula
	inputState.FathersName = input.FathersName
	inputState.MothersName = input.MothersName
	inputState.Address = input.Address
	inputState.Phones = input.Phones
	inputState.Nationality = input.Nationality
	inputState.Birthplace = input.Birthplace
	inputState.CPF = input.CPF
	inputState.Neighborhood = input.Neighborhood
	inputState.City = input.City
	inputState.ParentCPF = input.ParentCPF

	err = reportpdf.StatementSchooling(inputState)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	fileD, err := os.Open("pdf/statement-" + student.GetID() + ".pdf")
	if err != nil {
		log.Panic(err)
	}
	_, err = ioutil.ReadAll(fileD)
	if err != nil {
		log.Panic(err)
	}
	// w.Header().Add("content-type", "application/pdf")
	// w.Write(file_bytes)
	err = json.NewEncoder(w).Encode(inputState)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
