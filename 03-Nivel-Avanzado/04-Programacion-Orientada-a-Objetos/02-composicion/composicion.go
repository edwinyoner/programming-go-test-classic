/*
   Archivo: herencia.go
   Autor: Edwin Yoner
   Fecha: 28/10/2024
   Versión: 1.0
   Este archivo demuestra el uso efectivo de composición en Go.
*/

package main

import (
	"fmt"
	"time"
)

// Interfaz para definir comportamientos comunes
type Trabajador interface {
	Trabajar() string
	ObtenerSalario() float64
	ObtenerInfo() string
}

// Información de contacto reutilizable
type InfoContacto struct {
	Email     string
	Telefono  string
	Direccion string
}

func (i InfoContacto) ObtenerContacto() string {
	return fmt.Sprintf("Email: %s, Tel: %s, Dir: %s",
		i.Email, i.Telefono, i.Direccion)
}

// Información laboral base
type InfoLaboral struct {
	FechaIngreso time.Time
	Salario      float64
	Departamento string
}

func (i InfoLaboral) CalcularAntiguedad() int {
	return int(time.Since(i.FechaIngreso).Hours() / 24 / 365)
}

// Persona como tipo base
type Persona struct {
	Nombre       string
	Apellido     string
	Edad         int
	InfoContacto // Composición de InfoContacto
}

func (p Persona) NombreCompleto() string {
	return fmt.Sprintf("%s %s", p.Nombre, p.Apellido)
}

func (p Persona) Saludar() string {
	return fmt.Sprintf("Hola, soy %s", p.NombreCompleto())
}

// Empleado compone Persona e InfoLaboral
type Empleado struct {
	Persona     // Composición de Persona
	InfoLaboral // Composición de InfoLaboral
	Puesto      string
	Supervisor  *Empleado
}

// Constructor para Empleado
func NuevoEmpleado(nombre, apellido string, edad int,
	email, telefono, direccion string,
	fechaIngreso time.Time, salario float64,
	departamento, puesto string) *Empleado {

	return &Empleado{
		Persona: Persona{
			Nombre:   nombre,
			Apellido: apellido,
			Edad:     edad,
			InfoContacto: InfoContacto{
				Email:     email,
				Telefono:  telefono,
				Direccion: direccion,
			},
		},
		InfoLaboral: InfoLaboral{
			FechaIngreso: fechaIngreso,
			Salario:      salario,
			Departamento: departamento,
		},
		Puesto: puesto,
	}
}

// Implementación de la interfaz Trabajador
func (e Empleado) Trabajar() string {
	return fmt.Sprintf("%s está trabajando como %s en el departamento de %s",
		e.NombreCompleto(), e.Puesto, e.Departamento)
}

func (e Empleado) ObtenerSalario() float64 {
	return e.Salario
}

func (e Empleado) ObtenerInfo() string {
	return fmt.Sprintf(`
Información del Empleado:
------------------------
%s
Puesto: %s
Departamento: %s
Antigüedad: %d años
Contacto: %s
Supervisor: %s`,
		e.Saludar(),
		e.Puesto,
		e.Departamento,
		e.CalcularAntiguedad(),
		e.ObtenerContacto(),
		func() string {
			if e.Supervisor != nil {
				return e.Supervisor.NombreCompleto()
			}
			return "Sin supervisor"
		}())
}

// Contratista también implementa Trabajador pero de manera diferente
type Contratista struct {
	Persona
	TarifaHora   float64
	HorasTrabajo int
	Proyecto     string
}

func (c Contratista) Trabajar() string {
	return fmt.Sprintf("%s está trabajando en el proyecto: %s",
		c.NombreCompleto(), c.Proyecto)
}

func (c Contratista) ObtenerSalario() float64 {
	return c.TarifaHora * float64(c.HorasTrabajo)
}

func (c Contratista) ObtenerInfo() string {
	return fmt.Sprintf(`
Información del Contratista:
---------------------------
%s
Proyecto: %s
Tarifa por Hora: $%.2f
Horas Trabajadas: %d
Contacto: %s`,
		c.Saludar(),
		c.Proyecto,
		c.TarifaHora,
		c.HorasTrabajo,
		c.ObtenerContacto())
}

func main() {
	// Crear un supervisor
	supervisor := NuevoEmpleado(
		"Ana", "López", 35,
		"ana@empresa.com", "555-0100", "Calle Principal 123",
		time.Date(2018, time.January, 15, 0, 0, 0, 0, time.UTC),
		75000, "Tecnología", "Gerente de Desarrollo",
	)

	// Crear un empleado
	empleado := NuevoEmpleado(
		"Edwin", "Yoner", 28,
		"edwin@empresa.com", "555-0101", "Avenida Central 456",
		time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC),
		50000, "Tecnología", "Desarrollador Senior",
	)
	empleado.Supervisor = supervisor

	// Crear un contratista
	contratista := &Contratista{
		Persona: Persona{
			Nombre:   "Carlos",
			Apellido: "Ruiz",
			Edad:     31,
			InfoContacto: InfoContacto{
				Email:     "carlos@freelance.com",
				Telefono:  "555-0102",
				Direccion: "Plaza Mayor 789",
			},
		},
		TarifaHora:   50,
		HorasTrabajo: 160,
		Proyecto:     "Actualización Sistema Legacy",
	}

	// Demostrar el uso de la interfaz Trabajador
	trabajadores := []Trabajador{empleado, contratista}

	fmt.Println("=== Demostración de Composición en Go ===")
	for _, t := range trabajadores {
		fmt.Println(t.ObtenerInfo())
		fmt.Printf("Actividad: %s\n", t.Trabajar())
		fmt.Printf("Salario: $%.2f\n", t.ObtenerSalario())
		fmt.Println("----------------------------------------")
	}
}
