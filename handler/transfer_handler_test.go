package handler_test

import (
	"bytes"
	"century-pay/handler"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransferHandlerSuccess(t *testing.T) {
	// Create a new bank instance and add initial users
	bank := handler.NewBank()
	_ = bank.AddUser("Mark", 100)
	_ = bank.AddUser("Jane", 50)
	_ = bank.AddUser("Adam", 0)

	tests := []struct {
		name           string
		reqBody        handler.TransferRequest
		expectedStatus int
	}{
		{
			name: "Successful Transfer",
			reqBody: handler.TransferRequest{
				Sender:   "Mark",
				Receiver: "Jane",
				Amount:   30,
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Insufficient Funds",
			reqBody: handler.TransferRequest{
				Sender:   "Mark",
				Receiver: "Jane",
				Amount:   200, // More than Mark's balance
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Non-Existent Sender",
			reqBody: handler.TransferRequest{
				Sender:   "NonExistent",
				Receiver: "Jane",
				Amount:   10,
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Non-Existent Receiver",
			reqBody: handler.TransferRequest{
				Sender:   "Mark",
				Receiver: "NonExistent",
				Amount:   10,
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Transfer to Same User",
			reqBody: handler.TransferRequest{
				Sender:   "Mark",
				Receiver: "Mark", // Cannot transfer to yourself
				Amount:   10,
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the request body
			reqBody, err := json.Marshal(tt.reqBody)
			if err != nil {
				t.Fatalf("failed to marshal request body: %v", err)
			}

			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewReader(reqBody))
			w := httptest.NewRecorder()

			// Call the TransferMoney handler
			// Notice we pass the bank instance here
			handler := handler.TransferMoney(bank)
			handler(w, req)

			// Check the status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, w.Code)
			}

		})
	}
}
