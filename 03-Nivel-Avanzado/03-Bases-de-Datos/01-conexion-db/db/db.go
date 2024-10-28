/* Archivo: db.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo consumir una API externa en Go.
*/

package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DB encapsula la conexión a la base de datos y su configuración
type DB struct {
	*sql.DB
	config *Config
}

// Nuevo crea una nueva conexión a la base de datos
func Nuevo(config *Config) (*DB, error) {
	db, err := sql.Open("mysql", config.DSN())
	if err != nil {
		return nil, fmt.Errorf("error abriendo conexión: %w", err)
	}

	// Configurar pool de conexiones
	db.SetMaxOpenConns(config.MaxConexionesAbiertas)
	db.SetMaxIdleConns(config.MaxConexionesInactivas)
	db.SetConnMaxLifetime(config.TiempoMaximoVida)
	db.SetConnMaxIdleTime(config.TiempoMaximoInactivo)

	// Verificar conexión
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("error verificando conexión: %w", err)
	}

	return &DB{
		DB:     db,
		config: config,
	}, nil
}

// Verificar realiza una verificación completa de la conexión
func (db *DB) Verificar(ctx context.Context) error {
	var version string
	err := db.QueryRowContext(ctx, "SELECT VERSION()").Scan(&version)
	if err != nil {
		return fmt.Errorf("error verificando versión MySQL: %w", err)
	}
	log.Printf("Conectado a MySQL versión: %s", version)
	return nil
}

// EstadisticasConexion retorna estadísticas del pool de conexiones
func (db *DB) EstadisticasConexion() sql.DBStats {
	return db.Stats()
}

// CerrarConexion cierra la conexión a la base de datos
func (db *DB) CerrarConexion() error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("error cerrando conexión: %w", err)
	}
	return nil
}
