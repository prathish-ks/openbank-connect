package handler

import (
	"encoding/json"
	"net/http"

	"github.com/prathish-ks/openbank-connect/internal/client"
	"github.com/prathish-ks/openbank-connect/internal/middleware"
	"github.com/prathish-ks/openbank-connect/internal/service"
)

type Account struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

type Transaction struct {
	ID        string `json:"id"`
	Date      string `json:"date"`
	Amount    string `json:"amount"`
	Reference string `json:"reference"`
}

type TransactionWithRisk struct {
	Transaction
	RiskScore int `json:"risk_score"`
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	_, ok := service.GetToken("mock-user-id") // Simulate auth
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accounts := []Account{
		{ID: "acc1", Name: "Everyday Spending", Balance: "$1,234.56"},
		{ID: "acc2", Name: "Savings Vault", Balance: "$8,765.43"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	transactions := []Transaction{
		{ID: "txn1", Date: "2025-07-17", Amount: "-$50.00", Reference: "Uber"},
		{ID: "txn2", Date: "2025-07-16", Amount: "+$1,200.00", Reference: "Salary"},
	}

	// Call risk engine for first account (simulate)
	risk, err := client.GetRiskScore(userID, "acc1")
	if err != nil {
		http.Error(w, "Risk engine failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	enriched := []TransactionWithRisk{}
	for _, txn := range transactions {
		enriched = append(enriched, TransactionWithRisk{
			Transaction: txn,
			RiskScore:   risk,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enriched)
}
