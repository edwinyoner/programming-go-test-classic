/* Archivo: main.go
   Autor: Edwin Yoner
   Fecha: 26/10/2024
   Versión: 1.0

   Este archivo permite ejecutar la función `Sumar` y mostrar su resultado.
*/

package main

import (
	"fmt"
	"log"

	"github.com/edwinyoner/programming-go-test-classic/03-Nivel-Avanzado/01-Pruebas-y-Documentacion/02-documentacion/persona"
)

func main() {
	// Crear una nueva persona
	p, err := persona.Nueva("Carlos", 28)
	if err != nil {
		log.Fatalf("Error al crear persona: %v", err)
	}

	// Mostrar saludo
	fmt.Println(p.Saludar())

	// Calcular y mostrar año de nacimiento aproximado
	añoNacimiento := p.CalcularAñoNacimiento()
	fmt.Printf("Año de nacimiento aproximado: %d\n", añoNacimiento)

	// Actualizar edad y mostrar nuevo saludo
	p.ActualizarEdad()
	fmt.Println("Después de cumplir años:")
	fmt.Println(p.Saludar())

	// Intentar crear una persona con datos inválidos
	_, err = persona.Nueva("", -5)
	if err != nil {
		fmt.Printf("Error esperado: %v\n", err)
	}
}
