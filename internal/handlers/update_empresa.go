package handlers

import (
	"encoding/json"
	"miapp/internal/models"
	"miapp/internal/repositories"
	"net/http"
	"strconv"
	"strings"
)

func UpdateEmpresaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Extraer el ID de la URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "ID no proporcionado", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[3])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Verificar si la empresa existe
	empresa, err := repositories.GetEmpresaById(id)
	if err != nil {
		http.Error(w, "Empresa no encontrada", http.StatusNotFound)
		return
	}

	// Decodificar el cuerpo de la solicitud
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

	// Llamar a la función de actualización
	err = repositories.UpdateEmpresa(id, input)
	if err != nil {
		http.Error(w, "Error al actualizar la empresa", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":      empresa.ID,
		"message": "Empresa actualizada exitosamente",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
