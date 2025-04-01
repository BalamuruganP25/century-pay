package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func GetUserBalance(s *Bank) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Extract the user from the URL path
		user := chi.URLParam(r, "user")

		// Debug: Print the extracted user to confirm it's being extracted correctly
		fmt.Println("Extracted user:", user)

		if user == "" {
			http.Error(w, "user parameter is required", http.StatusBadRequest)
			return
		}

		// Fetch the user's balance
		balance, err := s.GetBalance(user)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving balance: %s", err.Error()), http.StatusNotFound)
			return
		}

		// Return the balance as JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Encode the balance as JSON
		err = json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}

	}

}
