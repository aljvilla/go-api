package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"
	"strings"
)

// BcryptHash genera un hash bcrypt usando SHA-256 manualmente sin librerías externas
func BcryptHash(password string) (string, error) {
	salt := os.Getenv("SALT")
	if salt == "" {
		return "", nil
	}
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(password))
	hashed := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return salt + hashed, nil
}

func CompareBcryptHash(password, hashedPassword string) bool {
	parts := strings.Split(hashedPassword, "$")
	if len(parts) < 2 {
		return false
	}

	// Volver a hashear la contraseña ingresada usando el mismo salt
	hashedInput, _ := BcryptHash(password)
	return hmac.Equal([]byte(hashedInput), []byte(hashedPassword))
}
