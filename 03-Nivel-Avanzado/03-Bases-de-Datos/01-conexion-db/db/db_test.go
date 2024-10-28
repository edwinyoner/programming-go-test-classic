/* Archivo: db_test.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo consumir una API externa en Go.
*/

package db

import (
	"context"
	"testing"
	"time"
)

func TestConexionDB(t *testing.T) {
	config := NuevaConfig()

	db, err := Nuevo(config)
	if err != nil {
		t.Fatalf("Error conectando a la base de datos: %v", err)
	}
	defer db.CerrarConexion()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.Verificar(ctx); err != nil {
		t.Errorf("Error verificando conexión: %v", err)
	}

	stats := db.EstadisticasConexion()
	if stats.OpenConnections > config.MaxConexionesAbiertas {
		t.Errorf("Demasiadas conexiones abiertas: %d", stats.OpenConnections)
	}
}
