/*
Archivo: bucles.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo usar bucles en Go.
*/

// Package main es el punto de entrada del programa.
package main

import "fmt" // Importa el paquete fmt para la salida de datos.

func main() {
    // Bucle for básico
    for i := 0; i < 5; i++ {
        fmt.Println("Iteración:", i)
    }

    // Bucle for con condicional
    contador := 0
    for contador < 5 {
        fmt.Println("Contador:", contador)
        contador++
    }

    // Bucle for con range
    frutas := []string{"manzana", "banana", "naranja"}
    for indice, fruta := range frutas {
        fmt.Println("Índice:", indice, "Fruta:", fruta)
    }
}
