/*
Archivo: interfaces.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo definir y usar interfaces en Go.
*/

package main

import "fmt" // Importa el paquete fmt para la salida de datos.

// Impresora define un comportamiento de impresión.
type Impresora interface {
    Imprimir()
}

// Documento representa un documento a imprimir.
type Documento struct {
    Titulo string
}

func (d Documento) Imprimir() {
    fmt.Println("Imprimiendo documento:", d.Titulo)
}

func main() {
    doc := Documento{Titulo: "Reporte Anual"}
    var impresora Impresora = doc
    impresora.Imprimir()
}
