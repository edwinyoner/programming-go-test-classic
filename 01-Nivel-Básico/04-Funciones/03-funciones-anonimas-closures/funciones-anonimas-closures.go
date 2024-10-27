/*
Archivo: funciones_anonimas_closures.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra el uso de funciones anónimas y closures en Go.
*/

// Package main es el punto de entrada del programa.
package main

import "fmt" // Importa el paquete fmt para la salida de datos.

func main() {
    // Función anónima
    suma := func(a, b int) int {
        return a + b
    }
    fmt.Println("El resultado de la suma es:", suma(10, 20))

    // Closure
    contador := 0
    incrementar := func() int {
        contador++
        return contador
    }

    fmt.Println("Contador:", incrementar())
    fmt.Println("Contador:", incrementar())
}
