/*
Archivo: manejo_errores.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo manejar errores en Go de manera básica.
*/

package main

import (
	"errors"
	"fmt"
)

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir entre cero")
	}
	return a / b, nil
}

func main() {
	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		//return
	} else {
		fmt.Println("Resultado:", resultado)
	}

	resultado, err = dividir(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Resultado:", resultado)
	}

}
