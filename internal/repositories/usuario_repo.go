package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"miapp/internal/database"
	"miapp/internal/models"
)

const TableUsuarios = "up_users"

// GetUsuarioByEmail busca un usuario en la base de datos por su email
func GetUsuarioByEmail(email string) (*models.Usuario, error) {
	query := `SELECT id, email, password FROM up_users WHERE email = $1`
	row := database.DB.QueryRow(context.Background(), query, email)

	var usuario models.Usuario
	err := row.Scan(&usuario.ID, &usuario.Email, &usuario.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}
	return &usuario, nil
}

// Actualiza la contraseña de un usuario
func UpdatePassword(userID int, newPassword string) error {
	query := `UPDATE up_users SET password = $1 WHERE id = $2`
	_, err := database.DB.Exec(context.Background(), query, newPassword, userID)
	return err
}

// Obtener usuario por ID (para comparar contraseñas)
func GetUsuarioByID(id int) (*models.Usuario, error) {
	query := fmt.Sprintf(`SELECT id, email, password FROM %s WHERE id = $1`, TableUsuarios)
	row := database.DB.QueryRow(context.Background(), query, id)

	var usuario models.Usuario
	err := row.Scan(&usuario.ID, &usuario.Email, &usuario.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}
	return &usuario, nil
}
