package handler
import (
	//"context"
    "net/http"
	"fmt"

    "github.com/prathish-ks/openbank-connect/internal/config"
	"github.com/prathish-ks/openbank-connect/internal/service"
     "golang.org/x/oauth2"
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

func AuthCallback(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    if code == "" {
        http.Error(w, "Missing code", http.StatusBadRequest)
        return
    }

    /* // Exchange the code for a token
    token, err := config.MockBankOAuthConfig.Exchange(context.Background(), code)
    if err != nil {
        http.Error(w, "Token exchange failed: "+err.Error(), http.StatusInternalServerError)
        return
    }
    // Store the token in-memory (for now)
    service.SaveToken("mock-user-id", token)

    fmt.Fprintf(w, "Token received and stored: %s", token.AccessToken) */

        // 🔁 Instead of real token exchange, simulate it
    fakeToken := &oauth2.Token{
        AccessToken:  "mock-access-token-" + code,
        RefreshToken: "mock-refresh-token",
        TokenType:    "Bearer",
    }

    // Store using mock user ID
    service.SaveToken("mock-user-id", fakeToken)

    fmt.Fprintf(w, "✅ Token received and stored!\nAccessToken: %s\n", fakeToken.AccessToken)

}