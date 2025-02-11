package validation

import (
    "errors"
    "regexp"
)

var allowedSortColumns = map[string]bool{
    "id": true,
    "razon_social": true,
    "numero_identificador": true,
    "tipo_numero_identificador": true,
}

var sortRegex = regexp.MustCompile(`^(asc|desc)$`)

func ValidatePaginationParams(page int, sortBy string, sort string) (string, string, error) {
    if page < 1 {
        return "", "", errors.New("el número de página debe ser mayor o igual a 1")
    }

    if _, exists := allowedSortColumns[sortBy]; !exists {
        return "", "", errors.New("columna de ordenación no permitida")
    }

    if !sortRegex.MatchString(sort) {
        return "", "", errors.New("el parámetro de orden debe ser 'asc' o 'desc'")
    }

    return sortBy, sort, nil
}