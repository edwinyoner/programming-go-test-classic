/* Archivo: cliente.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo consumir una API externa en Go.
*/

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Cliente maneja las peticiones a la API
type Cliente struct {
	httpClient *http.Client
	baseURL    string
}

// NuevoCliente crea una nueva instancia del cliente API
func NuevoCliente(timeout time.Duration) *Cliente {
	return &Cliente{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL: "https://jsonplaceholder.typicode.com",
	}
}

// hacer realiza la petición HTTP y decodifica la respuesta
func (c *Cliente) hacer(ctx context.Context, method, path string, result interface{}) error {
	url := fmt.Sprintf("%s%s", c.baseURL, path)

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return fmt.Errorf("error creando request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error haciendo request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		cuerpo, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error del servidor: código %d, respuesta: %s",
			resp.StatusCode, string(cuerpo))
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("error decodificando respuesta: %w", err)
	}

	return nil
}

// ObtenerUsuarios obtiene la lista de usuarios
func (c *Cliente) ObtenerUsuarios(ctx context.Context) ([]Usuario, error) {
	var usuarios []Usuario
	err := c.hacer(ctx, http.MethodGet, "/users", &usuarios)
	return usuarios, err
}

// ObtenerUsuario obtiene un usuario por su ID
func (c *Cliente) ObtenerUsuario(ctx context.Context, id int) (*Usuario, error) {
	var usuario Usuario
	err := c.hacer(ctx, http.MethodGet, fmt.Sprintf("/users/%d", id), &usuario)
	return &usuario, err
}

// ObtenerPostsDeUsuario obtiene los posts de un usuario específico
func (c *Cliente) ObtenerPostsDeUsuario(ctx context.Context, userID int) ([]Post, error) {
	var posts []Post
	err := c.hacer(ctx, http.MethodGet, fmt.Sprintf("/users/%d/posts", userID), &posts)
	return posts, err
}
