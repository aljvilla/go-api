package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

type ContextKey string

const UserIDKey ContextKey = "userID"

// Middleware para validar JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Token requerido", http.StatusForbidden)
			return
		}

		tokenParts := strings.Split(strings.Replace(token, "Bearer ", "", 1), ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}
		claims, jwtError := validateJWT(tokenParts)
		if jwtError != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// Valida el JWT
type JWTClaims struct {
	UserID int   `json:"user_id"`
	Exp    int64 `json:"exp"`
}

func validateJWT(tokenParts []string) (*JWTClaims, error) {
	payloadBytes, err := base64.RawURLEncoding.DecodeString(tokenParts[1])
	if err != nil {
		return nil, errors.New("error decodificando el payload")
	}

	var claims JWTClaims
	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return nil, errors.New("error parseando el payload")
	}

	// Verificar expiración
	if claims.Exp < time.Now().Unix() {
		return nil, errors.New("token expirado")
	}

	return &claims, nil
}
