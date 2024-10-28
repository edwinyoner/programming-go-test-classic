/* Archivo: operaciones_test.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

*/

package operaciones

import "testing"

func TestSumar(t *testing.T) {
	// Tabla de casos de prueba
	casos := []struct {
		nombre   string
		a, b     int
		esperado int
	}{
		{"números positivos", 2, 3, 5},
		{"con cero", 0, 5, 5},
		{"números negativos", -2, -3, -5},
		{"positivo y negativo", 2, -3, -1},
	}

	// Ejecutar casos de prueba
	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			resultado := Sumar(caso.a, caso.b)
			if resultado != caso.esperado {
				t.Errorf("Sumar(%d, %d) = %d; se esperaba %d",
					caso.a, caso.b, resultado, caso.esperado)
			}
		})
	}
}
