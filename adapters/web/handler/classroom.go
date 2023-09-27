package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
	"wwchacalww/go-cem304/domain/utils"
	"wwchacalww/go-cem304/usecase/check"
	reportpdf "wwchacalww/go-cem304/usecase/report-pdf"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type ClassroomHandler struct {
	Repo repository.ClassroomRepositoryInterface
}

func MakeClassroomHandlers(r *chi.Mux, repo repository.ClassroomRepositoryInterface) {
	handler := &ClassroomHandler{
		Repo: repo,
	}
	jwtoken := jwtauth.New("HS256", []byte("secret_jwt"), nil)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtoken))
		r.Use(jwtauth.Authenticator)
		r.Route("/classrooms", func(r chi.Router) {
			r.Post("/", handler.Store)
			r.Post("/import", handler.Import)
			r.Get("/{id}", handler.GetClassroom)
			r.Get("/search", handler.FindByName)
			r.Get("/list", handler.List)
			r.Patch("/enable/{id}", handler.Enable)
			r.Patch("/disable/{id}", handler.Disable)
			r.Patch("/anne", handler.ANNE)
			r.Put("/checkclassrooms", handler.CheckStudentsInClassrooms)
			r.Get("/report/{id}", handler.StudentsInClassPDF)
			r.Get("/report/all/", handler.AllClassroomsPDF)
			r.Get("/report/diary/{id}", handler.DiaryClassPDF)
			r.Get("/report/diary/all/", handler.DiaryAllClassroomsPDF)
			r.Get("/report/cover/{id}", handler.FolderCoverPDF)
		})
	})
}

func (c *ClassroomHandler) Store(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	var input repository.ClassroomInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	class, err := c.Repo.Create(input)
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

func (c *ClassroomHandler) GetClassroom(w http.ResponseWriter, r *http.Request) {
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
	class, err := c.Repo.FindById(id)
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

func (c *ClassroomHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin", "director", "secretary", "coordinator", "teather"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	name := r.URL.Query().Get("name")
	classrooms, err := c.Repo.FindByName(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(classrooms)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *ClassroomHandler) Enable(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	id := chi.URLParam(r, "id")
	class, err := c.Repo.Enable(id)
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

func (c *ClassroomHandler) Disable(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	id := chi.URLParam(r, "id")
	class, err := c.Repo.Disable(id)
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

func (c *ClassroomHandler) ANNE(w http.ResponseWriter, r *http.Request) {
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
	class, err := c.Repo.ANNE(input.ID, input.ANNE)
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

func (c *ClassroomHandler) Import(w http.ResponseWriter, r *http.Request) {
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

	list, err := utils.CsvToClassrooms(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	result, err := c.Repo.AddMass(list)
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

func (c *ClassroomHandler) StudentsInClassPDF(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	id := chi.URLParam(r, "id")
	class, err := c.Repo.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = reportpdf.StudentsInClass(class)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	fileD, err := os.Open("pdf/hello.pdf")
	if err != nil {
		log.Panic(err)
	}
	file_bytes, err := ioutil.ReadAll(fileD)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Add("content-type", "application/pdf")
	w.Write(file_bytes)
	os.Remove("pdf/hello.pdf")
}

func (c *ClassroomHandler) List(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin", "director", "secretary", "coordinator", "teather"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	year := r.URL.Query().Get("year")
	classrooms, err := c.Repo.List(year)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	var class []model.ClassroomInterface
	for _, cl := range classrooms {
		classroom, err := c.Repo.FindById(cl.GetID())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		class = append(class, classroom)
	}

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
	w.WriteHeader(http.StatusOK)
}

func (c *ClassroomHandler) AllClassroomsPDF(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin", "director", "secretary", "coordinator"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	classrooms, err := c.Repo.List("2023")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	var class []model.ClassroomInterface
	for _, cl := range classrooms {
		classroom, err := c.Repo.FindById(cl.GetID())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		class = append(class, classroom)
	}
	err = reportpdf.ReportAllClass(class)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	fileD, err := os.Open("pdf/all-classrooms.pdf")
	if err != nil {
		log.Panic(err)
	}
	file_bytes, err := ioutil.ReadAll(fileD)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Add("content-type", "application/pdf")
	w.Write(file_bytes)
	os.Remove("pdf/all-classrooms.pdf")
}

func (c *ClassroomHandler) DiaryClassPDF(w http.ResponseWriter, r *http.Request) {
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
	class, err := c.Repo.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = reportpdf.DiaryClass(class)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	fileD, err := os.Open("pdf/diary_class.pdf")
	if err != nil {
		log.Panic(err)
	}
	file_bytes, err := ioutil.ReadAll(fileD)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Add("content-type", "application/pdf")
	w.Write(file_bytes)
	os.Remove("pdf/diary_class.pdf")
}

func (c *ClassroomHandler) DiaryAllClassroomsPDF(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	roles := []string{"admin", "director", "secretary", "coordinator"}
	err := utils.CheckRoles(roles, fmt.Sprintf("%v", role))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError(err.Error()))
		return
	}
	classrooms, err := c.Repo.List("2023")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	var class []model.ClassroomInterface
	for _, cl := range classrooms {
		classroom, err := c.Repo.FindById(cl.GetID())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		class = append(class, classroom)
	}
	err = reportpdf.DiaryAllClass(class)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	fileD, err := os.Open("pdf/diary_all_classrooms.pdf")
	if err != nil {
		log.Panic(err)
	}
	file_bytes, err := ioutil.ReadAll(fileD)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Add("content-type", "application/pdf")
	w.Write(file_bytes)
	os.Remove("pdf/diary_all_classrooms.pdf")
}

func (c *ClassroomHandler) FolderCoverPDF(w http.ResponseWriter, r *http.Request) {
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
	class, err := c.Repo.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = reportpdf.FolderCoverClass(class)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	fileD, err := os.Open("pdf/folderCover.pdf")
	if err != nil {
		log.Panic(err)
	}
	file_bytes, err := ioutil.ReadAll(fileD)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Add("content-type", "application/pdf")
	w.Write(file_bytes)
	os.Remove("pdf/folderCover.pdf")
}

func (c *ClassroomHandler) CheckStudentsInClassrooms(w http.ResponseWriter, r *http.Request) {
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
	list, err := utils.StudentsInClassrooms(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	classrooms, err := c.Repo.List("2023")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	var class []model.ClassroomInterface
	for _, cl := range classrooms {
		classroom, err := c.Repo.FindById(cl.GetID())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		class = append(class, classroom)
	}
	output, err := check.CheckAllStudentsAndClass(list, class)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(201)
}
