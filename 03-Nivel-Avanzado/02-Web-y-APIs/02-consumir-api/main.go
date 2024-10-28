/* Archivo: main.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo consumir una API externa en Go.
*/

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/edwinyoner/programming-go-test-classic/03-Nivel-Avanzado/02-Web-y-APIs/02-consumir-api/api"
)

func main() {
	// Configurar logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Crear cliente con timeout de 10 segundos
	cliente := api.NuevoCliente(10 * time.Second)

	// Crear contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Obtener todos los usuarios
	usuarios, err := cliente.ObtenerUsuarios(ctx)
	if err != nil {
		log.Fatalf("Error obteniendo usuarios: %v", err)
	}

	// Mostrar información de usuarios
	for _, usuario := range usuarios {
		fmt.Printf("\nUsuario ID: %d\n", usuario.ID)
		fmt.Printf("Nombre: %s\n", usuario.Nombre)
		fmt.Printf("Email: %s\n", usuario.Email)
		fmt.Printf("Ciudad: %s\n", usuario.Direccion.Ciudad)
		fmt.Printf("Empresa: %s\n", usuario.Empresa.Nombre)

		// Obtener posts del usuario
		posts, err := cliente.ObtenerPostsDeUsuario(ctx, usuario.ID)
		if err != nil {
			log.Printf("Error obteniendo posts del usuario %d: %v", usuario.ID, err)
			continue
		}

		fmt.Printf("Número de posts: %d\n", len(posts))
		if len(posts) > 0 {
			fmt.Printf("Último post: %s\n", posts[0].Titulo)
		}
	}
}
