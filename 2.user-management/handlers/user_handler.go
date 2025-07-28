package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/neel07sanghvi/crud-api/models"
	"github.com/neel07sanghvi/crud-api/storage"
)

type UserHandler struct {
	storage *storage.UserStorage
}

func New(storage *storage.UserStorage) *UserHandler {
	return &UserHandler{storage: storage}
}

func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		h.handleGet(w, r)
	case "POST":
		h.handlePost(w, r)
	case "PUT":
		h.handlePut(w, r)
	case "DELETE":
		h.handleDelete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/users" {
		users := h.storage.GetAllUsers()
		json.NewEncoder(w).Encode(users)
	} else {
		id, err := h.extractIDFromPath(path)

		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)

			return
		}

		json.NewEncoder(w).Encode(id)
	}
}

func (h *UserHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" {
		http.Error(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	user := h.storage.CreateUser(req.Name, req.Email)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) handlePut(w http.ResponseWriter, r *http.Request) {
	id, err := h.extractIDFromPath(r.URL.Path)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req models.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" {
		http.Error(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	user, exists := h.storage.UpdateUser(id, req.Name, req.Email)

	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := h.extractIDFromPath(r.URL.Path)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if !h.storage.DeleteUser(id) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) extractIDFromPath(path string) (int, error) {
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		return 0, http.ErrNotSupported
	}

	return strconv.Atoi(parts[2])
}
