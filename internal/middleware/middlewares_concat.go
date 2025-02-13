package middleware

import "net/http"

// ChainMiddlewares permite encadenar varios middlewares de forma flexible
func MiddlewaresConcat(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- { // Aplica en orden correcto
		handler = middlewares[i](handler)
	}
	return handler
}
