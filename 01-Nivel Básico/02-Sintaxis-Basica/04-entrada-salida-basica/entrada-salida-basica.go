/*
Archivo: entrada_salida_basica.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo leer y escribir datos en la consola en Go.
*/

// Package main es el punto de entrada del programa.
package main

import "fmt" // Importa el paquete fmt para la entrada y salida de datos.

func main() {
    var nombre string
    var edad int

    // Solicitar entrada de datos al usuario
    fmt.Print("Ingrese su nombre: ")
    fmt.Scanln(&nombre)
    fmt.Print("Ingrese su edad: ")
    fmt.Scanln(&edad)

    // Mostrar los datos ingresados
    fmt.Println("Nombre ingresado:", nombre)
    fmt.Println("Edad ingresada:", edad)
}
