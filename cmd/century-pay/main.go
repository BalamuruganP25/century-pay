package main

import (
	"century-pay/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	initWebServer()
}

func initWebServer() {
	fmt.Println("server is started")
	bank := handler.NewBank()
	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Post("/v1/transaction/transfer_money", handler.TransferMoney(bank))
	r.Get("/v1/transaction/{user}/balance", handler.GetUserBalance(bank))
	r.Get("/v1/transaction/{user}/transaction_history", handler.GetTransacationHistory(bank))
	http.ListenAndServe(":8089", r)
}
