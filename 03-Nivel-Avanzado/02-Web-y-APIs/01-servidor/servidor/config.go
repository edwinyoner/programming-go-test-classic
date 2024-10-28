/* Archivo: config.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo construir un servidor HTTP básico en Go.
*/

package servidor

import (
	"os"
	"strconv"
)

// Config almacena la configuración del servidor
type Config struct {
	Puerto            int
	Host              string
	EntornoDesarrollo bool
}

// NuevaConfig crea una nueva configuración con valores por defecto o desde variables de entorno
func NuevaConfig() *Config {
	puerto, err := strconv.Atoi(getEnvODefault("PUERTO", "8080"))
	if err != nil {
		puerto = 8080
	}

	return &Config{
		Puerto:            puerto,
		Host:              getEnvODefault("HOST", "localhost"),
		EntornoDesarrollo: getEnvODefault("ENTORNO", "desarrollo") == "desarrollo",
	}
}

func getEnvODefault(key, valorDefault string) string {
	if valor, existe := os.LookupEnv(key); existe {
		return valor
	}
	return valorDefault
}
