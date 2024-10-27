/*
Archivo: declaracion_uso.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo declarar y usar funciones en Go.
*/

// Package main es el punto de entrada del programa.
package main

import "fmt" // Importa el paquete fmt para la salida de datos.

// Función que suma dos enteros y retorna el resultado.
func sumar(a int, b int) int {
    return a + b
}

// Función principal
func main() {
    resultado := sumar(3, 5)
    fmt.Println("El resultado de la suma es:", resultado)
}
