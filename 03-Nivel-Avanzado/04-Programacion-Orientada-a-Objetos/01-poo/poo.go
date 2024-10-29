/*
   Archivo: poo.go
   Autor: Edwin Yoner
   Fecha: 28/10/2024
   Versión: 1.0
   Este archivo demuestra conceptos avanzados de POO en Go.
*/

package main

import (
	"fmt"
	"time"
)

// Interface que define el comportamiento de un ser vivo
type SerVivo interface {
	Saludar() string
	ObtenerEdad() int
	Descripcion() string
}

// Struct base con campos comunes
type DatosPersonales struct {
	Nombre   string
	Apellido string
	FechaNac time.Time
}

// Persona ahora implementa la interface SerVivo
type Persona struct {
	DatosPersonales
	Profesion string
	Email     string
}

// Constructor para crear una nueva Persona
func NuevaPersona(nombre, apellido string, fechaNac time.Time, profesion, email string) *Persona {
	return &Persona{
		DatosPersonales: DatosPersonales{
			Nombre:   nombre,
			Apellido: apellido,
			FechaNac: fechaNac,
		},
		Profesion: profesion,
		Email:     email,
	}
}

// Método para obtener el nombre completo
func (p Persona) NombreCompleto() string {
	return fmt.Sprintf("%s %s", p.Nombre, p.Apellido)
}

// Implementación de la interface SerVivo
func (p Persona) Saludar() string {
	return fmt.Sprintf("¡Hola! Soy %s y trabajo como %s", p.NombreCompleto(), p.Profesion)
}

func (p Persona) ObtenerEdad() int {
	return int(time.Since(p.FechaNac).Hours() / 24 / 365)
}

func (p Persona) Descripcion() string {
	return fmt.Sprintf("Persona: %s\nEdad: %d años\nProfesión: %s\nEmail: %s",
		p.NombreCompleto(),
		p.ObtenerEdad(),
		p.Profesion,
		p.Email)
}

// Método para actualizar el email
func (p *Persona) ActualizarEmail(nuevoEmail string) {
	p.Email = nuevoEmail
}

// Estudiante hereda de Persona
type Estudiante struct {
	Persona
	Universidad string
	Carrera     string
	Semestre    int
}

// Constructor para crear un nuevo Estudiante
func NuevoEstudiante(persona Persona, universidad, carrera string, semestre int) *Estudiante {
	return &Estudiante{
		Persona:     persona,
		Universidad: universidad,
		Carrera:     carrera,
		Semestre:    semestre,
	}
}

// Sobreescribe el método Descripcion para Estudiante
func (e Estudiante) Descripcion() string {
	return fmt.Sprintf("%s\nUniversidad: %s\nCarrera: %s\nSemestre: %d",
		e.Persona.Descripcion(),
		e.Universidad,
		e.Carrera,
		e.Semestre)
}

func main() {
	// Crear una nueva persona
	fechaNac := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	persona := NuevaPersona("Edwin", "Yoner", fechaNac, "Desarrollador", "edwin@email.com")

	// Crear un estudiante
	estudiante := NuevoEstudiante(
		*persona,
		"Universidad Tecnológica",
		"Ingeniería en Sistemas",
		6,
	)

	// Demostrar el uso de la interface
	seres := []SerVivo{persona, estudiante}

	fmt.Println("=== Demostración de POO en Go ===")
	for _, ser := range seres {
		fmt.Printf("\n%s\n", ser.Saludar())
		fmt.Printf("\n%s\n", ser.Descripcion())
		fmt.Println("\n-------------------")
	}
}
