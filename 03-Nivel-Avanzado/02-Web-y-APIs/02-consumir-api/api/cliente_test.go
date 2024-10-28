/* Archivo: cliente_test.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo consumir una API externa en Go.
*/

package api

import (
	"context"
	"testing"
	"time"
)

func TestCliente(t *testing.T) {
	cliente := NuevoCliente(10 * time.Second)
	ctx := context.Background()

	t.Run("obtener usuarios", func(t *testing.T) {
		usuarios, err := cliente.ObtenerUsuarios(ctx)
		if err != nil {
			t.Fatalf("Error obteniendo usuarios: %v", err)
		}
		if len(usuarios) == 0 {
			t.Error("No se obtuvieron usuarios")
		}
	})

	t.Run("obtener usuario específico", func(t *testing.T) {
		usuario, err := cliente.ObtenerUsuario(ctx, 1)
		if err != nil {
			t.Fatalf("Error obteniendo usuario: %v", err)
		}
		if usuario.ID != 1 {
			t.Errorf("ID esperado 1, obtenido %d", usuario.ID)
		}
	})

	t.Run("obtener posts de usuario", func(t *testing.T) {
		posts, err := cliente.ObtenerPostsDeUsuario(ctx, 1)
		if err != nil {
			t.Fatalf("Error obteniendo posts: %v", err)
		}
		if len(posts) == 0 {
			t.Error("No se obtuvieron posts")
		}
	})
}
