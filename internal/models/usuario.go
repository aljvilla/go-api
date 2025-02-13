package models

// Usuario representa la estructura de un usuario en la base de datos
type Usuario struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // No devolver la contraseña en la respuesta
}
