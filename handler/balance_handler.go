package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUserBalance(s *Bank) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Extract the user from the URL path
		user, err := ExtractURLParam(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}

	}

}
