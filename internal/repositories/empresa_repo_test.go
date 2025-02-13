package repositories

import (
	"context"
	"log"
	"miapp/internal/database"
	"miapp/internal/models"
	"testing"
)

// Test para `InsertEmpresa`
func TestInsertEmpresa(t *testing.T) {
	// Conectar a la base de datos de prueba
	database.ConnectDB()
	defer database.CloseDB()

	// Crear una empresa de prueba
	razonSocial := "Empresa Test"
	numeroIdentificador := "123456789"
	tipoNumeroIdentificador := "RUC"

	empresa := models.Empresa{
		RazonSocial:             &razonSocial,
		NumeroIdentificador:     &numeroIdentificador,
		TipoNumeroIdentificador: &tipoNumeroIdentificador,
	}

	// Ejecutar la función que estamos probando
	id, err := InsertEmpresa(empresa)

	// Validaciones sin librerías externas
	if err != nil {
		t.Fatalf("❌ InsertEmpresa falló con error: %v", err)
	}
	if id == 0 {
		t.Fatalf("❌ El ID insertado no debe ser 0")
	}

	// Verificar que la empresa se guardó en la base de datos
	var count int
	query := "SELECT COUNT(*) FROM empresas WHERE id = $1"
	err = database.DB.QueryRow(context.Background(), query, id).Scan(&count)
	if err != nil {
		t.Fatalf("❌ Error al consultar la base de datos: %v", err)
	}
	if count == 0 {
		t.Fatalf("❌ La empresa no fue encontrada en la base de datos")
	}

	log.Println("✅ TestInsertEmpresa pasó exitosamente")
}
