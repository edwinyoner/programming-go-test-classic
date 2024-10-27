/*
Archivo: break_continue.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo usar las instrucciones break y continue en bucles.
*/

// Package main es el punto de entrada del programa.
package main

import "fmt" // Importa el paquete fmt para la salida de datos.

func main() {
    // Uso de break
    for i := 0; i < 10; i++ {
        if i == 5 {
            fmt.Println("Se encontró el número 5, saliendo del bucle.")
            break // Sale del bucle cuando i es igual a 5
        }
        fmt.Println("Número:", i)
    }

    // Uso de continue
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue // Salta a la siguiente iteración si i es par
        }
        fmt.Println("Número impar:", i)
    }
}
