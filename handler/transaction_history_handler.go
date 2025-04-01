package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTransacationHistory(s *Bank) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := ExtractURLParam(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if _, ok := s.users[user]; !ok {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		// Fetch the transaction history for the user
		transactions, err := s.GetTransactionHistory(user)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving transaction history: %s", err.Error()), http.StatusNotFound)
			return
		}

		// Return the transaction history as JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Encode the transactions slice as JSON
		err = json.NewEncoder(w).Encode(transactions)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}

}
