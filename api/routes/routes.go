package routes

import (
	"miapp/internal/handlers"
	"net/http"
)

// SetupRoutes configura las rutas usando net/http
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/empresas", handlers.GetEmpresasHandler)
	mux.HandleFunc("/empresas/create", handlers.CreateEmpresaHandler)
	mux.HandleFunc("/empresas/update/", handlers.UpdateEmpresaHandler)

	return mux
}
