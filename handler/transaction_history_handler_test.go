package handler_test

import (
	"century-pay/handler"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTransactionHistoryHandler(t *testing.T) {
	// Set up bank instance and users
	bank := handler.NewBank()
	_ = bank.AddUser("Mark", 100) // Add Mark with an initial balance of $100
	_ = bank.AddUser("Jane", 50)  // Add Jane with an initial balance of $50

	// Perform transfers to generate transaction history
	_ = bank.TransferMoney("Mark", "Jane", 30) // Mark transfers $30 to Jane

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
			// Corrected URL format with user path parameter
			url := "/v1/transaction/" + tt.user + "/transaction_history"
			fmt.Println("url ==>", url)
			req := httptest.NewRequest(http.MethodGet, "/v1/transaction/"+tt.user+"/transaction_history", nil)
			w := httptest.NewRecorder()

			// Call the handler
			handler := handler.GetTransacationHistory(bank)
			handler(w, req)

			// Check the status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, w.Code)
			}

		})
	}
}
