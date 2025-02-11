package database

import (
    "context"
    "log"
    "github.com/jackc/pgx/v5"
    "miapp/config"
)

var DB *pgx.Conn

func ConnectDB() {
    var err error
    dsn := config.GetDatabaseURL()
    
    DB, err = pgx.Connect(context.Background(), dsn)
    if err != nil {
        log.Fatalf("Error conectando a la base de datos: %v", err)
    }
    log.Println("Conexión a PostgreSQL establecida")
}

func CloseDB() {
    if DB != nil {
        DB.Close(context.Background())
    }
}