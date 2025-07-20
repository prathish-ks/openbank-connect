package middleware

import (
    "context"
    "net/http"

    "github.com/prathish-ks/openbank-connect/internal/service"
)

type contextKey string

const userIDKey contextKey = "userID"

func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Simulated auth: look for token from in-memory store
        userID := "mock-user-id"

        token, ok := service.GetToken(userID)
        if !ok || token.AccessToken == "" {
            http.Error(w, "Unauthorized from middleware", http.StatusUnauthorized)
            return
        }

        // Add userID to context
        ctx := context.WithValue(r.Context(), userIDKey, userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// GetUserID retrieves user ID from request context
func GetUserID(r *http.Request) string {
    if val := r.Context().Value(userIDKey); val != nil {
        return val.(string)
    }
    return ""
}
