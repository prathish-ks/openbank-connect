package handler

import (
    "net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("openbank-connect is healthy"))
}
