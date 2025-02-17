package handlers

import (
	"encoding/json"
	"miapp/internal/middleware"
	"miapp/internal/models"
	"miapp/internal/repositories"
	"net/http"
)

// BulkInsertEmpresaHandler inserta los datos en la DB
func BulkInsertEmpresaHandler(w http.ResponseWriter, r *http.Request) {
	empresas, ok := r.Context().Value(middleware.ValidatedDataKey).([]models.Empresa)
	if !ok || len(empresas) == 0 {
		http.Error(w, "No hay datos validados en la request", http.StatusBadRequest)
		return
	}

	var insertedIDs []int
	for _, empresa := range empresas {
		id, err := repositories.InsertEmpresa(empresa)
		if err != nil {
			http.Error(w, "Error al insertar empresas", http.StatusInternalServerError)
			return
		}
		insertedIDs = append(insertedIDs, id)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Empresas insertadas correctamente",
		"ids":     insertedIDs,
	})
}
