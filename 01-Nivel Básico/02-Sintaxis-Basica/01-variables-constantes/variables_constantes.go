/*
Archivo: variables_constantes.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo declarar y usar variables y constantes en Go.
*/

// Package main es el punto de entrada del programa.
package main

import "fmt" // Importa el paquete fmt para la salida de datos.

func main() {
    // Declaración de una variable
    var nombre string = "Edwin"
    edad := 25 // Declaración usando inferencia de tipo

    // Declaración de una constante
    const pi float64 = 3.1416

    fmt.Println("Nombre:", nombre)
    fmt.Println("Edad:", edad)
    fmt.Println("Valor de Pi:", pi)
}
