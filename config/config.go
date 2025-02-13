package config

import (
	"fmt"
	"os"
)

// LoadEnv verifica si las variables de entorno están definidas
func LoadEnv() {
	requiredVars := []string{"DATABASE_HOST", "DATABASE_NAME", "DATABASE_PASSWORD", "DATABASE_PORT", "DATABASE_USERNAME", "SALT"}

	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			fmt.Printf("⚠️ Advertencia: La variable de entorno %s no está definida\n", v)
		}
	}
}

// GetEnv obtiene una variable de entorno con un valor por defecto si no está definida
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetDatabaseURL() string {
	dsn := "postgres://" + os.Getenv("DATABASE_USERNAME") + ":" +
		os.Getenv("DATABASE_PASSWORD") + "@" +
		os.Getenv("DATABASE_HOST") + ":" +
		os.Getenv("DATABASE_PORT") + "/" +
		os.Getenv("DATABASE_NAME") + "?sslmode=disable"
	fmt.Printf("DB: %s\n", dsn)
	fmt.Printf("PASS: %s\n", os.Getenv("DATABASE_PASSWORD"))
	return dsn
}
