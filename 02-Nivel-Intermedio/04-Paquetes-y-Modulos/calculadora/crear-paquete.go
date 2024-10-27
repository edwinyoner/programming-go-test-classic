/*
Archivo: crear_paquete.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo crear un paquete personalizado en Go.
*/

// Package calculadora proporciona funciones para realizar operaciones aritméticas básicas.
package calculadora

// Sumar devuelve la suma de dos números enteros.
func Sumar(a, b int) int {
    return a + b
}

// Restar devuelve la resta de dos números enteros.
func Restar(a, b int) int {
    return a - b
}
