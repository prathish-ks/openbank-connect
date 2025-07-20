package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RiskRequest struct {
	UserID  string `json:"user_id"`
	Account string `json:"account_id"`
}

type RiskResponse struct {
	RiskScore int `json:"risk_score"`
}

func GetRiskScore(userID, accountID string) (int, error) {
	url := "http://localhost:9090/risk" // mock internal microservice

	reqBody, _ := json.Marshal(RiskRequest{
		UserID:  userID,
		Account: accountID,
	})

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return 0, fmt.Errorf("failed to call risk engine: %w", err)
	}
	defer resp.Body.Close()

	var result RiskResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("invalid risk engine response: %w", err)
	}

	return result.RiskScore, nil
}
