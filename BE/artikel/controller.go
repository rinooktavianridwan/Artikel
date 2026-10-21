package artikel

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ArtikelController struct {
	service *ArtikelService
}

func NewArtikelController(service *ArtikelService) *ArtikelController {
	return &ArtikelController{service: service}
}

func (c *ArtikelController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/artikel", c.GetAllArtikel).Methods("GET")
	router.HandleFunc("/artikel", c.CreateArtikel).Methods("POST")
	router.HandleFunc("/artikel/{id}", c.GetArtikelByID).Methods("GET")
	router.HandleFunc("/artikel/{id}", c.UpdateArtikel).Methods("PUT")
	router.HandleFunc("/artikel/{id}", c.DeleteArtikel).Methods("DELETE")
}

func (c *ArtikelController) GetAllArtikel(w http.ResponseWriter, r *http.Request) {
	limit := 10
	offset := 0

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	status := r.URL.Query().Get("status")
	if status == "" {
		status = PUBLISH
	}

	articles, err := c.service.GetAllArtikel(limit, offset, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func (c *ArtikelController) CreateArtikel(w http.ResponseWriter, r *http.Request) {
	var artikel Artikel
	if err := json.NewDecoder(r.Body).Decode(&artikel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.CreateArtikel(&artikel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *ArtikelController) GetArtikelByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Missing artikel ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artikel ID", http.StatusBadRequest)
		return
	}

	artikel, err := c.service.GetArtikelByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artikel)
}

func (c *ArtikelController) UpdateArtikel(w http.ResponseWriter, r *http.Request) {
	var artikel Artikel
	if err := json.NewDecoder(r.Body).Decode(&artikel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.UpdateArtikel(&artikel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artikel)
}

func (c *ArtikelController) DeleteArtikel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Missing artikel ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artikel ID", http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteArtikel(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
