/*
Archivo: condicionales.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo usar estructuras condicionales en Go.
*/

// Package main es el punto de entrada del programa.
package main

import "fmt" // Importa el paquete fmt para la salida de datos.

func main() {
    edad := 18

    // Estructura condicional if-else
    if edad >= 18 {
        fmt.Println("Eres mayor de edad.")
    } else {
        fmt.Println("Eres menor de edad.")
    }

    // Estructura condicional con else if
    nota := 85

    if nota >= 90 {
        fmt.Println("Calificación: A")
    } else if nota >= 80 {
        fmt.Println("Calificación: B")
    } else if nota >= 70 {
        fmt.Println("Calificación: C")
    } else {
        fmt.Println("Calificación: D")
    }
}
