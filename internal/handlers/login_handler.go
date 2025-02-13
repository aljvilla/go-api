package handlers

import (
	"encoding/json"
	"miapp/internal/repositories"
	"miapp/internal/utils"
	"net/http"
)

// Estructura para recibir las credenciales de login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Estructura para la respuesta con JWT
type LoginResponse struct {
	Token string `json:"token"`
}

// LoginHandler maneja la autenticación de usuarios
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var creds LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Error en el formato del JSON", http.StatusBadRequest)
		return
	}

	// Buscar usuario en la base de datos
	usuario, err := repositories.GetUsuarioByEmail(creds.Email)
	if err != nil {
		http.Error(w, "Usuario o contraseña incorrectos este", http.StatusUnauthorized)
		return
	}

	// Validar la contraseña usando bcrypt manualmente
	if !utils.CompareBcryptHash(creds.Password, usuario.Password) {
		http.Error(w, "Usuario o contraseña incorrectos aqui ", http.StatusUnauthorized)
		return
	}

	// Generar token JWT
	token, err := utils.GenerateJWT(usuario.ID)
	if err != nil {
		http.Error(w, "Error generando el token", http.StatusInternalServerError)
		return
	}

	// Responder con el JWT
	response := LoginResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
