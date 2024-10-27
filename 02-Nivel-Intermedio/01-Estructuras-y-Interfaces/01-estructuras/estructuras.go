/*
Archivo: estructuras.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo definir y usar estructuras en Go.
*/

package main

import "fmt" // Importa el paquete fmt para la salida de datos.

// Persona representa una persona con nombre y edad.
type Persona struct {
    Nombre string
    Edad   int
}

func main() {
    persona := Persona{Nombre: "Juan", Edad: 30}
    fmt.Println("Nombre:", persona.Nombre)
    fmt.Println("Edad:", persona.Edad)
}
