/*
Archivo: defer.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra el uso de la declaración defer en funciones.
*/

// Package main es el punto de entrada del programa.
package main

import "fmt" // Importa el paquete fmt para la salida de datos.

func main() {
    defer fmt.Println("Esta línea se ejecutará al final de la función main.")
    fmt.Println("Ejecutando la función main...")
}
