/*
Archivo: errores_personalizados.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo crear y manejar errores personalizados en Go.
*/

package main

import (
    "fmt"
)

// DivideError representa un error de división personalizada.
type DivideError struct {
    Numerador   float64
    Denominador float64
}

func (e *DivideError) Error() string {
    return fmt.Sprintf("no se puede dividir %v entre %v", e.Numerador, e.Denominador)
}

func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, &DivideError{Numerador: a, Denominador: b}
    }
    return a / b, nil
}

func main() {
    resultado, err := dividir(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Resultado:", resultado)
}
