package config

import (
    "golang.org/x/oauth2"
)

var MockBankOAuthConfig = &oauth2.Config{
    ClientID:     "mockbank-client-id",
    ClientSecret: "mockbank-client-secret",
    RedirectURL:  "http://localhost:8080/auth/callback",
    Scopes:       []string{"accounts", "transactions"},
    Endpoint: oauth2.Endpoint{
        AuthURL:  "http://localhost:8081/oauth/authorize", // This will be mocked later
        TokenURL: "http://localhost:8081/oauth/token",
    },
}
