/* Archivo: config.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo consumir una API externa en Go.
*/

package db

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config almacena la configuración de la base de datos
type Config struct {
	Usuario    string
	Contraseña string
	Host       string
	Puerto     int
	Nombre     string
	Parametros map[string]string
	// Configuración del pool de conexiones
	MaxConexionesAbiertas  int
	MaxConexionesInactivas int
	TiempoMaximoVida       time.Duration
	TiempoMaximoInactivo   time.Duration
}

// NuevaConfig crea una nueva configuración desde variables de entorno
func NuevaConfig() *Config {
	puerto, _ := strconv.Atoi(getEnvODefault("DB_PUERTO", "3306"))
	maxConn, _ := strconv.Atoi(getEnvODefault("DB_MAX_CONEXIONES", "25"))
	maxIdle, _ := strconv.Atoi(getEnvODefault("DB_MAX_INACTIVAS", "25"))
	vidaConn, _ := time.ParseDuration(getEnvODefault("DB_TIEMPO_VIDA", "5m"))
	tiempoIdle, _ := time.ParseDuration(getEnvODefault("DB_TIEMPO_INACTIVO", "5m"))

	return &Config{
		Usuario:    getEnvODefault("DB_USUARIO", "root"),
		Contraseña: getEnvODefault("DB_PASSWORD", "edwinyoner"),
		Host:       getEnvODefault("DB_HOST", "localhost"),
		Puerto:     puerto,
		Nombre:     getEnvODefault("DB_NOMBRE", "GO"),
		Parametros: map[string]string{
			"parseTime":    "true",
			"loc":          "Local",
			"charset":      "utf8mb4",
			"collation":    "utf8mb4_unicode_ci",
			"timeout":      "5s",
			"writeTimeout": "5s",
			"readTimeout":  "5s",
		},
		MaxConexionesAbiertas:  maxConn,
		MaxConexionesInactivas: maxIdle,
		TiempoMaximoVida:       vidaConn,
		TiempoMaximoInactivo:   tiempoIdle,
	}
}

func getEnvODefault(key, valorDefault string) string {
	if valor, existe := os.LookupEnv(key); existe {
		return valor
	}
	return valorDefault
}

// DSN genera la cadena de conexión para MySQL
func (c *Config) DSN() string {
	parametros := ""
	for key, valor := range c.Parametros {
		if parametros != "" {
			parametros += "&"
		}
		parametros += fmt.Sprintf("%s=%s", key, valor)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		c.Usuario, c.Contraseña, c.Host, c.Puerto, c.Nombre, parametros)
}
