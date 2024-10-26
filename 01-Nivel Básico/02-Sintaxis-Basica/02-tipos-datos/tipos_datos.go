/*
Archivo: tipos_datos.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra los tipos de datos básicos en Go y cómo usarlos.
*/

// Package main es el punto de entrada del programa.
package main

import "fmt" // Importa el paquete fmt para la salida de datos.

func main() {
    var entero int = 10
    var decimal float64 = 3.14
    var booleano bool = true
    var texto string = "Hola, Go!"

    fmt.Println("Entero:", entero)
    fmt.Println("Decimal:", decimal)
    fmt.Println("Booleano:", booleano)
    fmt.Println("Texto:", texto)
}
