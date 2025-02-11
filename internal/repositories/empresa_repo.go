package repositories

import (
	"context"
	"errors"
	"log"
	"miapp/internal/database"
	"miapp/internal/models"
)

func GetEmpresas(limit, offset int, sortBy, sort string) ([]models.Empresa, error) {
	query := `SELECT id, razon_social, numero_identificador, tipo_numero_identificador 
              FROM empresas 
              ORDER BY ` + sortBy + ` ` + sort + ` 
              LIMIT $1 OFFSET $2`

	rows, err := database.DB.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var empresas []models.Empresa
	for rows.Next() {
		var empresa models.Empresa
		if err := rows.Scan(
			&empresa.ID,
			&empresa.RazonSocial,
			&empresa.NumeroIdentificador,
			&empresa.TipoNumeroIdentificador); err != nil {
			return nil, err
		}
		empresas = append(empresas, empresa)
	}
	return empresas, nil
}

// Verifica si una empresa ya existe
func EmpresaExists(numeroIdentificador, tipoIdentificador *string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM empresas WHERE numero_identificador = $1 AND tipo_numero_identificador = $2`
	err := database.DB.QueryRow(context.Background(), query, numeroIdentificador, tipoIdentificador).Scan(&count)
	return count > 0, err
}

// Inserta una nueva empresa
func InsertEmpresa(input models.Empresa) (int, error) {
	var id int
	query := `
						INSERT 
						INTO 
							empresas 
								(
									razon_social, 
									numero_identificador, 
									tipo_numero_identificador
								) 
						VALUES 
								($1, $2, $3)
						RETURNING id;`
	err := database.DB.QueryRow(context.Background(), query, input.RazonSocial, input.NumeroIdentificador, input.TipoNumeroIdentificador).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Actualiza una empresa
func UpdateEmpresa(id int, input models.Empresa) error {
	query := `
						UPDATE 
							empresas 
						SET 
							razon_social = $1, 
							tipo_numero_identificador = $2 
						WHERE 
							id = $3;`
	_, err := database.DB.Exec(context.Background(), query, input.RazonSocial, input.TipoNumeroIdentificador, id)
	log.Println(err)
	return err
}

// Obtiene una empresa por su ID
func GetEmpresaById(id int) (*models.Empresa, error) {
	query := `SELECT id, razon_social, numero_identificador, tipo_numero_identificador FROM empresas WHERE id = $1`
	row := database.DB.QueryRow(context.Background(), query, id)

	var empresa models.Empresa
	err := row.Scan(&empresa.ID, &empresa.RazonSocial, &empresa.NumeroIdentificador, &empresa.TipoNumeroIdentificador)
	if err != nil {
		return nil, errors.New("empresa no encontrada")
	}

	return &empresa, nil
}
