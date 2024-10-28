package persona

import (
    "fmt"
    "time"
)

// Persona representa a un individuo con sus datos básicos.
type Persona struct {
    Nombre string // Nombre completo de la persona
    Edad   int    // Edad en años
}

// Nueva crea una nueva instancia de Persona con los datos proporcionados.
// Retorna error si los datos no son válidos.
func Nueva(nombre string, edad int) (*Persona, error) {
    if nombre == "" {
        return nil, fmt.Errorf("el nombre no puede estar vacío")
    }
    if edad < 0 || edad > 150 {
        return nil, fmt.Errorf("la edad debe estar entre 0 y 150 años")
    }
    return &Persona{
        Nombre: nombre,
        Edad:   edad,
    }, nil
}

// Saludar retorna un saludo personalizado con el nombre de la persona.
func (p *Persona) Saludar() string {
    return fmt.Sprintf("¡Hola! Me llamo %s y tengo %d años", p.Nombre, p.Edad)
}

// CalcularAñoNacimiento estima el año de nacimiento basado en la edad actual.
// Este es un cálculo aproximado que asume que ya se cumplieron años en el año actual.
func (p *Persona) CalcularAñoNacimiento() int {
    añoActual := time.Now().Year()
    return añoActual - p.Edad
}

// ActualizarEdad incrementa la edad de la persona en un año.
func (p *Persona) ActualizarEdad() {
    p.Edad++
}