package handler

import (
	"encoding/json"
	"net/http"
)

func TransferMoney(s *Bank) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req TransferRequest

		// Decode the request body
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Process the transfer
		err = s.TransferMoney(req.Sender, req.Receiver, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Transfer successful"))
	}

}
