package handler
import (
    "net/http"

    "github.com/prathish-ks/openbank-connect/internal/config"
)

func StartAuth(w http.ResponseWriter, r *http.Request) {
    provider := r.URL.Query().Get("provider")
    if provider != "mockbank" {
        http.Error(w, "unsupported provider", http.StatusBadRequest)
        return
    }

    // State should be random and validated later (static for now)
    state := "abc123"
    url := config.MockBankOAuthConfig.AuthCodeURL(state)
    http.Redirect(w, r, url, http.StatusFound)
}
