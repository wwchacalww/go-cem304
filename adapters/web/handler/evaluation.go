package handler

import (
	"encoding/json"
	"net/http"
	"wwchacalww/go-cem304/domain/repository"
	"wwchacalww/go-cem304/domain/utils"

	"github.com/go-chi/chi/v5"
)

type EvaluationHandler struct {
	Repo repository.EvaluationRepositoryInterface
}

func MakeEvaluationHandlers(r *chi.Mux, repo repository.EvaluationRepositoryInterface) {
	handler := &EvaluationHandler{
		Repo: repo,
	}

	r.Route("/evaluations", func(r chi.Router) {
		r.Post("/", handler.Store)
		r.Post("/import", handler.ImportCards)
		r.Put("/change", handler.ChangeEvaluation)
	})
}

func (e *EvaluationHandler) Store(w http.ResponseWriter, r *http.Request) {
	var input repository.EvaluationInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	eval, err := e.Repo.Create(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(eval)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (e *EvaluationHandler) ChangeEvaluation(w http.ResponseWriter, r *http.Request) {
	var ip struct {
		ID       int    `json:"id"`
		Note     string `json:"note"`
		Absences int    `json:"absences"`
	}
	err := json.NewDecoder(r.Body).Decode(&ip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	eval, err := e.Repo.Update(ip.Note, ip.ID, ip.Absences)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(eval)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (e *EvaluationHandler) ImportCards(w http.ResponseWriter, r *http.Request) {
	f, fh, err := r.FormFile("file")
	classroom_id := r.FormValue("classroom_id")
	term := r.FormValue("term")
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

	cards, err := utils.CsvToCards(f, term, classroom_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = e.Repo.ImportCSVEvaluation(cards)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(cards)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
