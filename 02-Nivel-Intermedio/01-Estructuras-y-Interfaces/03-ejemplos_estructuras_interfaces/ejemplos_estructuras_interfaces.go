/*
Archivo: ejemplos_estructuras_interfaces.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo proporciona ejemplos prácticos que combinan estructuras e interfaces.
*/

package main

import "fmt"

// Definición de una interfaz y estructura para demostrar su uso en conjunto.
type Describible interface {
    Describir() string
}

type Producto struct {
    Nombre  string
    Precio  float64
}

func (p Producto) Describir() string {
    return fmt.Sprintf("Producto: %s, Precio: %.2f", p.Nombre, p.Precio)
}

func main() {
    prod := Producto{Nombre: "Laptop", Precio: 1499.99}
    var desc Describible = prod
    fmt.Println(desc.Describir())
}
