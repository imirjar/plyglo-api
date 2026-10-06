package controller

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

// AuthMiddleware выполняет проверку токена через OIDC-верификатор Keycloak.
func (srv *HTTP) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const (
			issuerURL = "https://auth.redbeaver.ru/realms/local-dev"
			clientID  = "local-go-api"
		)

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		rawToken := parts[1]

		provider, err := oidc.NewProvider(r.Context(), issuerURL)
		if err != nil {
			log.Printf("OIDC provider init error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		verifier := provider.Verifier(&oidc.Config{
			ClientID:          clientID,
			SkipClientIDCheck: true, // отключаем проверку aud для теста
		})

		idToken, err := verifier.Verify(r.Context(), rawToken)
		if err != nil {
			log.Printf("token verification failed: %v", err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			log.Printf("failed to parse claims: %v", err)
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
