package routes

import (
	"miapp/internal/handlers"
	"miapp/internal/middleware"
	"net/http"
)

// SetupRoutes configura las rutas usando net/http
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.Handle("/upload", middleware.MiddlewaresConcat(
		http.HandlerFunc(handlers.BulkInsertEmpresaHandler),
		middleware.AuthMiddleware,
		middleware.UploadFileMiddleware,
		middleware.ValidateEmpresaMiddleware,
	))
	mux.Handle("/user/update-password",
		middleware.MiddlewaresConcat(
			http.HandlerFunc(handlers.UpdatePasswordHandler),
			middleware.AuthMiddleware,
		),
	)
	mux.HandleFunc("/empresas", handlers.GetEmpresasHandler)
	mux.Handle("/empresas/create",
		middleware.MiddlewaresConcat(
			http.HandlerFunc(handlers.CreateEmpresaHandler),
			middleware.AuthMiddleware,
		),
	)
	mux.HandleFunc("/empresas/update/", handlers.UpdateEmpresaHandler)
	mux.HandleFunc("/empresas/delete/", handlers.DeleteEmpresaHandler)
	return mux
}
