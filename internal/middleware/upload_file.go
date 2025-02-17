package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"miapp/internal/parsers"
)

// Clave para almacenar los datos parseados en el contexto
type contextKey string

const UploadDataKey contextKey = "uploadData"

// UploadFileMiddleware procesa un archivo y lo guarda en el contexto
func UploadFileMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20) // 10MB máximo

		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "No se pudo obtener el archivo internal/middleware/upload_file.go", http.StatusBadRequest)
			return
		}
		defer file.Close()

		fmt.Printf("📂 Archivo recibido: %s\n", handler.Filename)

		// Detectar la extensión del archivo
		extension := strings.ToLower(handler.Filename)
		var rawData [][]string

		if strings.HasSuffix(extension, ".csv") {
			rawData, err = parsers.ParseCSV(file)
		} else {
			http.Error(w, "Formato de archivo no soportado", http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, fmt.Sprintf("Error procesando el archivo: %v", err), http.StatusInternalServerError)
			return
		}

		// Almacenar los datos en el contexto de la request
		ctx := context.WithValue(r.Context(), UploadDataKey, rawData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
