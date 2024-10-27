/*
Archivo: usar_paquete.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo importar y utilizar un paquete personalizado en Go.
*/

package main

import (
    "fmt"
    "github.com/edwinyoner/programming-go-test-classic/02-Nivel-Intermedio/04-Paquetes-y-Modulos/calculadora"
)

func main() {
    suma := calculadora.Sumar(5, 3)
    resta := calculadora.Restar(10, 4)
    fmt.Println("Suma:", suma)
    fmt.Println("Resta:", resta)
}
