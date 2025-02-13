package handlers

import (
	"encoding/json"
	"miapp/internal/middleware"
	"miapp/internal/repositories"
	"miapp/internal/utils"
	"net/http"
)

// Estructura para recibir los datos de actualización de contraseña
type UpdatePasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

// UpdatePasswordHandler maneja la actualización de contraseña
func UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	// Extraer el UserID desde el contexto
	userID := r.Context().Value(middleware.UserIDKey)
	if userID == nil {
		http.Error(w, "Usuario no autorizado", http.StatusUnauthorized)
		return
	}

	var input UpdatePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Error en el formato del JSON", http.StatusBadRequest)
		return
	}

	// Validar que se ingresen ambas contraseñas
	if input.CurrentPassword == "" || input.NewPassword == "" {
		http.Error(w, "Debe ingresar la contraseña actual y la nueva contraseña", http.StatusBadRequest)
		return
	}

	// Obtener usuario desde la base de datos
	usuario, err := repositories.GetUsuarioByID(userID.(int))
	if err != nil {
		http.Error(w, "Usuario no encontrado >", http.StatusNotFound)
		return
	}

	// Validar que la contraseña actual sea correcta
	if !utils.CompareBcryptHash(input.CurrentPassword, usuario.Password) {
		http.Error(w, "La contraseña actual es incorrecta", http.StatusUnauthorized)
		return
	}

	// Hashear la nueva contraseña
	hashedPassword, err := utils.BcryptHash(input.NewPassword)
	if err != nil {
		http.Error(w, "Error al procesar la nueva contraseña", http.StatusInternalServerError)
		return
	}

	// Actualizar la contraseña en la base de datos
	err = repositories.UpdatePassword(usuario.ID, hashedPassword)
	if err != nil {
		http.Error(w, "Error al actualizar la contraseña", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Contraseña actualizada exitosamente"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
