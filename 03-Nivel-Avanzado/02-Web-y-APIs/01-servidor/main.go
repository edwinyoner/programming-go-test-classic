/* Archivo: main.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo construir un servidor HTTP básico en Go.
*/

package main

import (
	"log"

	"github.com/edwinyoner/programming-go-test-classic/03-Nivel-Avanzado/02-Web-y-APIs/01-servidor/servidor"
)

func main() {
	// Configurar logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Crear configuración
	config := servidor.NuevaConfig()

	// Crear e iniciar servidor
	srv := servidor.Nuevo(config)
	if err := srv.Iniciar(); err != nil {
		log.Fatal(err)
	}
}
