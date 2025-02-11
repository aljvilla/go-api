package handlers

import (
	"encoding/json"
	"fmt"
	"miapp/internal/repositories"
	"miapp/internal/validation"
	"net/http"
	"strconv"
)

func GetEmpresasHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	sortBy := r.URL.Query().Get("sortBy")
	if sortBy == "" {
		sortBy = "id"
	}

	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "asc"
	}

	validatedSortBy, validatedSort, validationErr := validation.ValidatePaginationParams(page, sortBy, sort)
	if validationErr != nil {
		http.Error(w, validationErr.Error(), http.StatusBadRequest)
		return
	}

	limit := 10
	offset := (page - 1) * limit

	empresas, err := repositories.GetEmpresas(limit, offset, validatedSortBy, validatedSort)
	if err != nil {
		fmt.Printf("Error conectando a la base de datos: %s", err.Error())
		http.Error(w, "Error al obtener datos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(empresas)
}
