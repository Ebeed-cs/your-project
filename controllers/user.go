package controllers

import (
	"encoding/json"
	"net/http"
	"sec_2/models"
	"strconv"
)

// userController represents the HTTP controller for user-related actions
type userController struct{}

// Handle POST request – create a new user
func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	u, err = models.AddUser(u)
	if err != nil {
		http.Error(w, "Error adding user", http.StatusInternalServerError)
		return
	}

	encodeResponseAsJSON(u, w)
}

// Handle GET request – fetch user by ID
func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	encodeResponseAsJSON(u, w)
}

// Helper to parse JSON request body into a User struct
func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

// Main router for /users endpoint
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			ids := r.URL.Query()["id"]
			if len(ids) == 0 {
				http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
				return
			}

			userID, err := strconv.Atoi(ids[0])
			if err != nil {
				http.Error(w, "Invalid 'id' value", http.StatusBadRequest)
				return
			}

			uc.get(userID, w)

		case http.MethodPost:
			uc.post(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	}
}

// Constructor for userController
func newUserController() *userController {
	return &userController{}
}
