/*
Archivo: parametros_retorno.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo manejar parámetros y valores de retorno en funciones.
*/

// Package main es el punto de entrada del programa.
package main

import "fmt" // Importa el paquete fmt para la salida de datos.

// Función que calcula el área de un rectángulo y retorna el área.
func areaRectangulo(base, altura float64) float64 {
    return base * altura
}

// Función principal
func main() {
    area := areaRectangulo(5.0, 3.0)
    fmt.Println("El área del rectángulo es:", area)
}
