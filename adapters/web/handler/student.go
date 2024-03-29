package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"wwchacalww/go-cem304/domain/repository"
	"wwchacalww/go-cem304/domain/utils"
	"wwchacalww/go-cem304/usecase/check"
	reportpdf "wwchacalww/go-cem304/usecase/report-pdf"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type StudentHandler struct {
	Repo repository.StudentRepositoryInterface
}

func MakeStudentHandlers(r *chi.Mux, repo repository.StudentRepositoryInterface) {
	handler := &StudentHandler{
		Repo: repo,
	}

	jwtoken := jwtauth.New("HS256", []byte("secret_jwt"), nil)
	r.Group(func(r chi.Router) {
		r.Route("/student", func(r chi.Router) {
			r.Get("/student-transcript", handler.StudentTranscript)
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtoken))
		r.Use(jwtauth.Authenticator)
		r.Route("/students", func(r chi.Router) {
			r.Post("/", handler.Store)
			r.Post("/import", handler.Import)
			r.Post("/import/report", handler.ImportReport)
			r.Get("/{id}", handler.GetStudent)
			r.Get("/educar/{id}", handler.FindByEducar)
			r.Get("/search", handler.FindByName)
			r.Get("/list", handler.List)
			r.Get("/carro", handler.StudentTranscript)
			r.Put("/change", handler.ChangeClassroom)
			r.Put("/checkclass", handler.CheckStudentsInClass)
			r.Post("/statement/schooling", handler.StatementSchooling)
			r.Put("/check/list", handler.CheckStudentsList)
			r.Put("/update/list/educadf", handler.UpdateEducaDFList)
			// r.Patch("/enable/{id}", handler.Enable)
			// r.Patch("/disable/{id}", handler.Disable)
			// r.Patch("/anne", handler.ANNE)
		})
	})
}

func (s *StudentHandler) Store(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"secretary", "admin"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	var input repository.StudentInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	class, err := s.Repo.Save(input)
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
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"secretary", "admin"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	var input struct {
		ID          string `json:"id"`
		ClassroomID string `json:"classroom_id"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
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
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin", "director", "secretary", "coordinator"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
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
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin", "director", "secretary", "coordinator"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	name := r.URL.Query().Get("name")
	Students, err := s.Repo.FindByName(strings.ToUpper(name))
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
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin", "director", "secretary", "coordinator", "teather"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
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
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin", "secretary"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
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
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin", "secretary"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
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
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin", "secretary"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	var input struct {
		ID   string `json:"id"`
		ANNE string `json:"ANNE"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
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
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
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

func (s *StudentHandler) ImportReport(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	f, fh, err := r.FormFile("file")
	classroom_id := r.FormValue("classroom_id")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	ct := fh.Header.Get("Content-Type")
	if ct != "text/plain" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError("File type invalid"))
		return
	}
	defer f.Close()
	result, err := utils.ReportToStudents(f, classroom_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	students, err := s.Repo.AddMassReport(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	// list, err := utils.CsvToClassrooms(f)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write(jsonError(err.Error()))
	// 	return
	// }

	// result, err := c.Repo.AddMass(list)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write(jsonError(err.Error()))
	// 	return
	// }

	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *StudentHandler) CheckStudentsInClass(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
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

func (s *StudentHandler) StudentTranscript(w http.ResponseWriter, r *http.Request) {
	err := reportpdf.StudentTranscript()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	fileD, err := os.Open("pdf/student-transcript.pdf")
	if err != nil {
		log.Panic(err)
	}
	file_bytes, err := ioutil.ReadAll(fileD)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Add("content-type", "application/pdf")
	w.Write(file_bytes)
}

func (s *StudentHandler) StatementSchooling(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin", "secretary"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	// id := chi.URLParam(r, "id")
	// log.Println(id)
	// educar, err := strconv.ParseInt(id, 10, 64)
	var input reportpdf.InputSS
	err = json.NewDecoder(r.Body).Decode(&input)
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
	inputState.CEP = input.CEP
	inputState.PhoneOne = input.PhoneOne
	inputState.PhoneTwo = input.PhoneTwo
	inputState.PhoneThree = input.PhoneThree
	inputState.Nationality = input.Nationality
	inputState.Birthplace = input.Birthplace
	inputState.CPF = input.CPF
	inputState.Neighborhood = input.Neighborhood
	inputState.City = input.City
	inputState.ParentCPF = input.ParentCPF
	// inputState.Matricula = "55555-5"
	// inputState.FathersName = "Fulando de Tal"
	// inputState.MothersName = "Maria da Silva"
	// inputState.Address = "QR 304 conjunto 13 casa 6"
	// inputState.Phones = "(61) 9999-9999"
	// inputState.Nationality = "Brasileira"
	// inputState.Birthplace = "Brasília-DF"
	// inputState.CPF = "111.111.111-11"
	// inputState.Neighborhood = "Samambaia Sul"
	// inputState.City = "Samambaia"
	// inputState.ParentCPF = "222.222.222-22"

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

func (s *StudentHandler) CheckStudentsList(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	f, fh, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	ct := fh.Header.Get("Content-Type")
	if ct != "text/plain" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError("File type invalid"))
		return
	}
	defer f.Close()
	result, err := utils.CheckStudentsList(f)
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
	w.WriteHeader(http.StatusOK)
}

func (s *StudentHandler) UpdateEducaDFList(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
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
	classroom_id := r.FormValue("classroom_id")
	classroom, err := s.Repo.List(classroom_id)
	result, err := utils.CheckEducaDFList(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	for _, std := range result {
		student, err := utils.FindByName(std.Name, classroom)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		log.Println(student.GetEducar(), std.Name, std.EducaDF)
		err = s.Repo.ChangeEducaDF(student.GetID(), std.EducaDF)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	}

	err = json.NewEncoder(w).Encode(classroom)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
