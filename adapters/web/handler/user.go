package handler

import (
	"encoding/json"
	"net/http"
	"wwchacalww/go-cem304/application"
	"wwchacalww/go-cem304/application/dto"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserHandler struct {
	UserService application.UserServiceInterface
}

func MakeUserHandlers(r *chi.Mux, service application.UserServiceInterface) {
	handler := &UserHandler{
		UserService: service,
	}

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.Store)
		r.Get("/", handler.GetList)
	})

}

func (u *UserHandler) GetList(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserService.List()
	if err != nil {
		// return c.JSON(http.StatusUnprocessableEntity, err.Error())
		render.Status(r, http.StatusUnprocessableEntity)
		w.Write([]byte("Error"))
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

func (u *UserHandler) Store(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.UserInput
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	user, err := u.UserService.Create(userDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
