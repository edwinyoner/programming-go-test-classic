/* Archivo: middleware.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo construir un servidor HTTP básico en Go.
*/

package servidor

import (
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware registra información sobre cada petición
func (s *Servidor) LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()

		next(w, r)

		log.Printf(
			"%s %s %s %v",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(inicio),
		)
	}
}

// RecuperarMiddleware maneja pánicos en los handlers
func (s *Servidor) RecuperarMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				s.ErrorHandler(w, r, http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}
