package handlers

import (
	"encoding/json"
	"miapp/internal/repositories"
	"net/http"
	"strconv"
	"strings"
)

func DeleteEmpresaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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

	// Llamar a la función de actualización
	err = repositories.DeleteEmpresa(id)
	if err != nil {
		http.Error(w, "Error al eliminar la empresa", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":      empresa.ID,
		"message": "Empresa eliminada exitosamente",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
