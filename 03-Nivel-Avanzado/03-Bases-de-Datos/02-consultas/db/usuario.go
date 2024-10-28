/* Archivo: consultas.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo realizar operaciones CRUD en una base de datos.
*/

package db

import (
    "context"
    "database/sql"
    "errors"
    "fmt"
    "time"
)

// Usuario representa la estructura de un usuario en la base de datos
type Usuario struct {
    ID        int64     `json:"id"`
    Nombre    string    `json:"nombre"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// UsuarioService maneja todas las operaciones relacionadas con usuarios
type UsuarioService struct {
    db *DB
}

// NuevoUsuarioService crea una nueva instancia del servicio de usuarios
func NuevoUsuarioService(db *DB) *UsuarioService {
    return &UsuarioService{db: db}
}

// Crear inserta un nuevo usuario en la base de datos
func (s *UsuarioService) Crear(ctx context.Context, usuario *Usuario) error {
    query := `
        INSERT INTO usuarios (nombre, email, created_at, updated_at)
        VALUES (?, ?, ?, ?)
    `
    
    ahora := time.Now()
    usuario.CreatedAt = ahora
    usuario.UpdatedAt = ahora

    resultado, err := s.db.ExecContext(ctx, query,
        usuario.Nombre,
        usuario.Email,
        usuario.CreatedAt,
        usuario.UpdatedAt,
    )
    if err != nil {
        return fmt.Errorf("error al crear usuario: %w", err)
    }

    id, err := resultado.LastInsertId()
    if err != nil {
        return fmt.Errorf("error al obtener ID insertado: %w", err)
    }

    usuario.ID = id
    return nil
}

// ObtenerPorID busca un usuario por su ID
func (s *UsuarioService) ObtenerPorID(ctx context.Context, id int64) (*Usuario, error) {
    query := `
        SELECT id, nombre, email, created_at, updated_at
        FROM usuarios
        WHERE id = ?
    `

    usuario := &Usuario{}
    err := s.db.QueryRowContext(ctx, query, id).Scan(
        &usuario.ID,
        &usuario.Nombre,
        &usuario.Email,
        &usuario.CreatedAt,
        &usuario.UpdatedAt,
    )

    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("usuario no encontrado con ID %d", id)
    }
    if err != nil {
        return nil, fmt.Errorf("error al buscar usuario: %w", err)
    }

    return usuario, nil
}

// Actualizar modifica los datos de un usuario existente
func (s *UsuarioService) Actualizar(ctx context.Context, usuario *Usuario) error {
    query := `
        UPDATE usuarios
        SET nombre = ?, email = ?, updated_at = ?
        WHERE id = ?
    `

    usuario.UpdatedAt = time.Now()
    resultado, err := s.db.ExecContext(ctx, query,
        usuario.Nombre,
        usuario.Email,
        usuario.UpdatedAt,
        usuario.ID,
    )
    if err != nil {
        return fmt.Errorf("error al actualizar usuario: %w", err)
    }

    filas, err := resultado.RowsAffected()
    if err != nil {
        return fmt.Errorf("error al verificar actualización: %w", err)
    }
    if filas == 0 {
        return errors.New("no se encontró el usuario para actualizar")
    }

    return nil
}

// Eliminar borra un usuario de la base de datos
func (s *UsuarioService) Eliminar(ctx context.Context, id int64) error {
    query := `DELETE FROM usuarios WHERE id = ?`

    resultado, err := s.db.ExecContext(ctx, query, id)
    if err != nil {
        return fmt.Errorf("error al eliminar usuario: %w", err)
    }

    filas, err := resultado.RowsAffected()
    if err != nil {
        return fmt.Errorf("error al verificar eliminación: %w", err)
    }
    if filas == 0 {
        return errors.New("no se encontró el usuario para eliminar")
    }

    return nil
}

// ListarTodos obtiene todos los usuarios con paginación
func (s *UsuarioService) ListarTodos(ctx context.Context, limite, offset int) ([]*Usuario, error) {
    query := `
        SELECT id, nombre, email, created_at, updated_at
        FROM usuarios
        ORDER BY id DESC
        LIMIT ? OFFSET ?
    `

    filas, err := s.db.QueryContext(ctx, query, limite, offset)
    if err != nil {
        return nil, fmt.Errorf("error al listar usuarios: %w", err)
    }
    defer filas.Close()

    var usuarios []*Usuario
    for filas.Next() {
        usuario := &Usuario{}
        err := filas.Scan(
            &usuario.ID,
            &usuario.Nombre,
            &usuario.Email,
            &usuario.CreatedAt,
            &usuario.UpdatedAt,
        )
        if err != nil {
            return nil, fmt.Errorf("error al escanear usuario: %w", err)
        }
        usuarios = append(usuarios, usuario)
    }

    if err = filas.Err(); err != nil {
        return nil, fmt.Errorf("error al iterar usuarios: %w", err)
    }

    return usuarios, nil
}