package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

type RiskRequest struct {
	UserID  string `json:"user_id"`
	Account string `json:"account_id"`
}

type RiskResponse struct {
	RiskScore int `json:"risk_score"`
}

func main() {
	http.HandleFunc("/risk", func(w http.ResponseWriter, r *http.Request) {
		var req RiskRequest
		_ = json.NewDecoder(r.Body).Decode(&req)

		// Fake risk logic
		score := rand.Intn(101) // 0 to 100

		json.NewEncoder(w).Encode(RiskResponse{RiskScore: score})
	})

	log.Println("🚨 Risk engine running on :9090")
	http.ListenAndServe(":9090", nil)
}
