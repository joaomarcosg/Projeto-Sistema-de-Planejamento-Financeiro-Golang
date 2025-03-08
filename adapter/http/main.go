package http

import (
	"financial-planning-system/adapter/http/actuator"
	"financial-planning-system/adapter/http/transactions"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init() {

	http.HandleFunc("/transactions", transactions.GetTransactions)
	http.HandleFunc("/transactions/create", transactions.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)

}
