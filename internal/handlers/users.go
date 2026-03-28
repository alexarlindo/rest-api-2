package handlers

import (
	"encoding/json"
	"net/http"

	"rest-api-2/internal/models"
)

func (h Handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.usecases.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h Handlers) AddUser(w http.ResponseWriter, r *http.Request) {

	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	id := h.usecases.Add(user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateUserResponse{NewUserID: id})

}

func (h Handlers) registerUserEndpoints(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			h.getAllUsers(w, r)
		case http.MethodPost:
			h.AddUser(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)	
	}
}
