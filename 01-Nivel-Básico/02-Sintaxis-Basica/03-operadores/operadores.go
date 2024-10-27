/*
Archivo: operadores.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra los operadores básicos en Go: aritméticos, relacionales y lógicos.
*/

// Package main es el punto de entrada del programa.
package main

import "fmt" // Importa el paquete fmt para la salida de datos.

func main() {
    // Operadores aritméticos
    suma := 5 + 3
    resta := 10 - 2
    producto := 4 * 2
    cociente := 8 / 4

    fmt.Println("Suma:", suma)
    fmt.Println("Resta:", resta)
    fmt.Println("Producto:", producto)
    fmt.Println("Cociente:", cociente)

    // Operadores relacionales
    mayor := 10 > 5
    igual := 5 == 5

    fmt.Println("¿10 es mayor que 5?", mayor)
    fmt.Println("¿5 es igual a 5?", igual)

    // Operadores lógicos
    and := true && false
    or := true || false

    fmt.Println("Operador AND:", and)
    fmt.Println("Operador OR:", or)
}
