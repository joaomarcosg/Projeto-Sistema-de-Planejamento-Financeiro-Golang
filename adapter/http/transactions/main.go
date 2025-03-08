package transactions

import (
	"encoding/json"
	"financial-planning-system/model/transactions"
	"financial-planning-system/util"
	"fmt"
	"io"
	"net/http"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = transactions.Transactions{
		transactions.Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2025-03-06T15:33:30"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)

}

func CreateATransaction(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var response = transactions.Transactions{}

	var body, _ = io.ReadAll(r.Body)

	_ = json.Unmarshal(body, &response)

	fmt.Print(response)
}
