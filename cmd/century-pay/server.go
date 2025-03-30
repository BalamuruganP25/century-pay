package main

import (
	"century-pay/handler"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func initWebServer() {
	bank := handler.NewBank()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("v1/transaction/transfer_money", handler.TransferMoney(bank))
	r.Get("v1/balance/{user}", handler.GetUserBalance())
	r.Get("v1/transaction/history/{user}", handler.GetTransacationHistory())
	http.ListenAndServe(":8080", r)
}
