package handler_test

import (
	"century-pay/handler"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestBalanceHandler(t *testing.T) {
	// Set up bank instance and users
	bank := handler.NewBank()
	_ = bank.AddUser("Mark", 100)
	_ = bank.AddUser("Jane", 50)

	tests := []struct {
		name           string
		user           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Get Mark's balance",
			user:           "Mark",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Get Jane's balance",
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
			r.Get("/v1/transaction/{user}/balance", handler.GetUserBalance(bank))
			req := httptest.NewRequest("GET", fmt.Sprintf("/v1/transaction/%s/balance", tt.user), nil)
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			// Check the response status code
			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, rr.Code)
			}

		})
	}
}
