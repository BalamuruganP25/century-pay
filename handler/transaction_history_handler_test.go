package handler_test

import (
	"century-pay/handler"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestGetTransactionHistoryHandler(t *testing.T) {
	// Set up bank instance and users
	bank := handler.NewBank()
	_ = bank.AddUser("Mark", 100)
	_ = bank.AddUser("Jane", 50)

	// Perform transfers to generate transaction history
	_ = bank.TransferMoney("Mark", "Jane", 30)

	tests := []struct {
		name           string
		user           string
		expectedStatus int
	}{
		{
			name:           "Get Mark's transaction history",
			user:           "Mark",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Get Jane's transaction history",
			user:           "Jane",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "User not found",
			user:           "NonExistent",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "No user provided",
			user:           "",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := chi.NewRouter()
			r.Get("/v1/transaction/{user}/transaction_history", handler.GetTransacationHistory(bank))
			req := httptest.NewRequest("GET", fmt.Sprintf("/v1/transaction/%s/transaction_history", tt.user), nil)
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			// Check the response status code
			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, rr.Code)
			}

		})
	}
}
