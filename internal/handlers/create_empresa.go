package handlers

import (
	"encoding/json"
	"log"
	"miapp/internal/middleware"
	"miapp/internal/models"
	"miapp/internal/repositories"
	"net/http"
)

func CreateEmpresaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey)
	if userID == nil {
		http.Error(w, "Usuario no autorizado", http.StatusUnauthorized)
		return
	}
	log.Println("Usuario autorizado:", userID)

	var input models.Empresa
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Error en el formato del JSON", http.StatusBadRequest)
		return
	}

	if input.RazonSocial == nil || input.NumeroIdentificador == nil || input.TipoNumeroIdentificador == nil {
		http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
		return
	}

	if *input.RazonSocial == "" || *input.NumeroIdentificador == "" || *input.TipoNumeroIdentificador == "" {
		http.Error(w, "Todos los campos deben contener un valor", http.StatusBadRequest)
		return
	}

	exists, err := repositories.EmpresaExists(input.NumeroIdentificador, input.TipoNumeroIdentificador)
	if err != nil {
		http.Error(w, "Error al verificar la empresa", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "La empresa ya existe", http.StatusConflict)
		return
	}

	id, err := repositories.InsertEmpresa(input)
	if err != nil {
		http.Error(w, "Error al crear la empresa", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":      id,
		"message": "Empresa creada correctamente",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
