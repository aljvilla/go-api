package main

import (
	"log"
	"miapp/api/routes"
	"miapp/config"
	"miapp/internal/database"
	"net/http"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()
	defer database.CloseDB()

	router := routes.SetupRoutes()

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
