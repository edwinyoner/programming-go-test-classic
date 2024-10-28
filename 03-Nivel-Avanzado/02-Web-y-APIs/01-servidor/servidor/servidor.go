/* Archivo: servidor.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo construir un servidor HTTP básico en Go.
*/

package servidor

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Servidor encapsula la configuración y funcionalidad del servidor HTTP
type Servidor struct {
	servidor *http.Server
	config   *Config
}

// Nuevo crea una nueva instancia del servidor
func Nuevo(config *Config) *Servidor {
	s := &Servidor{
		config: config,
	}

	mux := http.NewServeMux()

	// Registrar rutas
	mux.HandleFunc("/", s.RecuperarMiddleware(s.LoggerMiddleware(s.HomeHandler)))
	mux.HandleFunc("/info", s.RecuperarMiddleware(s.LoggerMiddleware(s.InfoHandler)))

	s.servidor = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Puerto),
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return s
}

// Iniciar inicia el servidor HTTP
func (s *Servidor) Iniciar() error {
	// Canal para señales de sistema
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.servidor.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error al iniciar el servidor: %v", err)
		}
	}()

	log.Printf("Servidor iniciado en http://%s:%d", s.config.Host, s.config.Puerto)

	// Esperar señal de terminación
	<-done
	log.Print("Servidor deteniendo...")

	// Contexto para shutdown graceful
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.servidor.Shutdown(ctx); err != nil {
		log.Printf("Error al detener el servidor: %v", err)
		return err
	}

	log.Print("Servidor detenido correctamente")
	return nil
}