/*
Archivo: canales.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo demuestra el uso de canales para comunicación entre goroutines en Go.
*/

package main

import "fmt"

func enviarMensaje(canal chan string) {
    canal <- "Mensaje enviado a través del canal"
}

func main() {
    canal := make(chan string)
    go enviarMensaje(canal)
    mensaje := <-canal
    fmt.Println("Recibido:", mensaje)
}
