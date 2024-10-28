/* Archivo: main.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo consumir una API externa en Go.
*/

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/edwinyoner/programming-go-test-classic/03-Nivel-Avanzado/03-Bases-de-Datos/01-conexion-db/db"
)

func main() {
	// Configurar logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Crear configuración
	config := db.NuevaConfig()

	// Conectar a la base de datos
	database, err := db.Nuevo(config)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}
	defer database.CerrarConexion()

	// Verificar conexión
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := database.Verificar(ctx); err != nil {
		log.Fatalf("Error verificando conexión: %v", err)
	}

	// Canal para señales de sistema
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Mostrar estadísticas periódicamente
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			log.Println("Cerrando conexión a la base de datos...")
			return
		case <-ticker.C:
			stats := database.EstadisticasConexion()
			log.Printf("Estadísticas DB - Conexiones abiertas: %d, En uso: %d, Inactivas: %d",
				stats.OpenConnections,
				stats.InUse,
				stats.Idle)
		}
	}
}
