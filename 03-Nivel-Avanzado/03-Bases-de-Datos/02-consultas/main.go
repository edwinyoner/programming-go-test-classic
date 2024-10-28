/* Archivo: consultas.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo realizar operaciones CRUD en una base de datos.
*/

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/edwinyoner/programming-go-test-classic/03-Nivel-Avanzado/03-Bases-de-Datos/02-consultas/db"
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

	// Crear servicio de usuarios
	usuarioService := db.NuevoUsuarioService(database)

	// Ejemplo de crear un usuario
	ctx := context.Background()
	nuevoUsuario := &db.Usuario{
		Nombre: "Yoner",
		Email:  "yoner@ejemplo.com",
	}

	if err := usuarioService.Crear(ctx, nuevoUsuario); err != nil {
		log.Printf("Error al crear usuario: %v", err)
	} else {
		log.Printf("Usuario creado con ID: %d", nuevoUsuario.ID)
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

			// Ejemplo de listar usuarios
			usuarios, err := usuarioService.ListarTodos(ctx, 10, 0)
			if err != nil {
				log.Printf("Error al listar usuarios: %v", err)
			} else {
				log.Printf("Usuarios totales: %d", len(usuarios))
			}
		}
	}
}
