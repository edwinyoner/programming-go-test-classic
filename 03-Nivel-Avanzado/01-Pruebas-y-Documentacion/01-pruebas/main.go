/* Archivo: main.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo permite ejecutar la función `Sumar` y mostrar su resultado.
*/

package main

import (
	"fmt"

	"github.com/edwinyoner/programming-go-test-classic/03-Nivel-Avanzado/01-Pruebas-y-Documentacion/01-pruebas/operaciones"
)

func main() {
	resultado := operaciones.Sumar(2, 3)
	fmt.Printf("Resultado de la suma: %d\n", resultado)
}
