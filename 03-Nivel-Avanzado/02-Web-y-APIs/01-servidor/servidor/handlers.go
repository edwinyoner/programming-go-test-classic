/* Archivo: handlers.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo construir un servidor HTTP básico en Go.
*/

package servidor

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Respuesta define la estructura de respuesta JSON
type Respuesta struct {
	Mensaje string    `json:"mensaje"`
	Hora    time.Time `json:"hora"`
	Status  string    `json:"status"`
}

// HomeHandler maneja la ruta principal
func (s *Servidor) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	respuesta := Respuesta{
		Mensaje: "¡Bienvenido a nuestro servidor web en Go!",
		Hora:    time.Now(),
		Status:  "success",
	}

	s.ResponderJSON(w, respuesta)
}

// InfoHandler proporciona información sobre el servidor
func (s *Servidor) InfoHandler(w http.ResponseWriter, r *http.Request) {
	info := struct {
		Version string    `json:"version"`
		Hora    time.Time `json:"hora"`
		Host    string    `json:"host"`
		Puerto  int       `json:"puerto"`
	}{
		Version: "1.0",
		Hora:    time.Now(),
		Host:    s.config.Host,
		Puerto:  s.config.Puerto,
	}

	s.ResponderJSON(w, info)
}

// ErrorHandler maneja los errores HTTP de forma consistente
func (s *Servidor) ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	respuesta := Respuesta{
		Mensaje: http.StatusText(status),
		Hora:    time.Now(),
		Status:  "error",
	}

	w.WriteHeader(status)
	s.ResponderJSON(w, respuesta)
}

// ResponderJSON envía una respuesta JSON al cliente
func (s *Servidor) ResponderJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error al codificar respuesta JSON: %v", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
	}
}
