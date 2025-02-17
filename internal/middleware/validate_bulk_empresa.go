package middleware

import (
	"context"
	"miapp/internal/models"
	"net/http"
	"strconv"
)

// Context key para almacenar los datos validados
const ValidatedDataKey contextKey = "validatedData"

// ValidateEmpresaMiddleware valida los datos y los convierte en `Empresa`
func ValidateEmpresaMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawData, ok := r.Context().Value(UploadDataKey).([][]string)
		if !ok || len(rawData) == 0 {
			http.Error(w, "No se encontraron datos en el archivo", http.StatusBadRequest)
			return
		}

		var empresas []models.Empresa

		// Omitir la primera fila (asumimos que es encabezado)
		for i, row := range rawData {
			if i == 0 {
				continue
			}

			if len(row) < 4 {
				http.Error(w, "El archivo CSV tiene menos columnas de las esperadas", http.StatusBadRequest)
				return
			}

			id, err := strconv.Atoi(row[0])
			if err != nil {
				http.Error(w, "El ID debe ser un número", http.StatusBadRequest)
				return
			}

			empresa := models.Empresa{
				ID:                      id,
				RazonSocial:             &row[1],
				NumeroIdentificador:     &row[2],
				TipoNumeroIdentificador: &row[3],
			}
			empresas = append(empresas, empresa)
		}

		// Guardar los datos validados en el contexto
		ctx := context.WithValue(r.Context(), ValidatedDataKey, empresas)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
