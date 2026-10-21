package main

import (
	"artikel/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"artikel/artikel"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	err = database.Migrate(db)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	repo := artikel.NewArtikelRepository(db)
	service := artikel.NewArtikelService(repo)
	controller := artikel.NewArtikelController(service)

	router := mux.NewRouter()
	controller.RegisterRoutes(router)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", cors(router)))
}
