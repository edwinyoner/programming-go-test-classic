/*
Archivo: goroutines.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo utilizar goroutines en Go.
*/

package main

import (
    "fmt"
    "time"
)

func imprimirMensaje(mensaje string) {
    fmt.Println(mensaje)
}

func main() {
    go imprimirMensaje("Mensaje desde una goroutine")
    fmt.Println("Mensaje desde la función main")
    time.Sleep(1 * time.Second) // Espera para permitir la ejecución de la goroutine.
}
