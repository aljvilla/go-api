package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"
)

// Clave secreta para firmar el token
var secretKey = []byte("vVYuP/atkiF6zP7atREaXQ==")

// Payload del JWT
type JWTClaims struct {
	UserID int   `json:"user_id"`
	Exp    int64 `json:"exp"`
}

// Genera un JWT con expiración de 1 hora
func GenerateJWT(userID int) (string, error) {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))

	payloadData := JWTClaims{
		UserID: userID,
		Exp:    time.Now().Add(time.Hour * 1).Unix(),
	}
	payloadBytes, _ := json.Marshal(payloadData)
	payload := base64.RawURLEncoding.EncodeToString(payloadBytes)

	signature := SignMessage(header + "." + payload)
	return header + "." + payload + "." + signature, nil
}

// Firma el token con HMAC-SHA256
func SignMessage(message string) string {
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(message))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
